// Copyright 2024 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package rac2

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/kvflowcontrol"
	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/kvflowcontrol/kvflowcontrolpb"
	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/kvflowcontrol/kvflowinspectpb"
	"github.com/cockroachdb/cockroach/pkg/raft/raftpb"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)

// entryID is the ID of a raft log entry. Should not be confused with LogMark.
// TODO(pav-kv): export raft.entryID and use it here.
type entryID struct {
	index uint64
	term  uint64
}

// Tracker tracks flow token deductions for a replicaSendStream. Tokens are
// deducted for an in-flight log entry identified by its raft entryID, with
// a given priority.
type Tracker struct {
	// tracked contains the per-priority tracked log entries ordered by log index.
	tracked [raftpb.NumPriorities][]tracked

	stream kvflowcontrol.Stream // used for logging only
}

// tracked represents a flow token deduction, identified by the raft entry ID.
type tracked struct {
	id     entryID
	tokens kvflowcontrol.Tokens
}

func (t *Tracker) Init(stream kvflowcontrol.Stream) {
	*t = Tracker{
		tracked: [raftpb.NumPriorities][]tracked{},
		stream:  stream,
	}
}

func (t *Tracker) Empty() bool {
	// TODO(pav-kv): can optimize this loop out if needed. We can maintain the
	// total number of tokens held, and return whether it's zero. It's also
	// possible to make it atomic and avoid locking the mutex in replicaSendStream
	// when calling this.
	for pri := range t.tracked {
		if len(t.tracked[pri]) != 0 {
			return false
		}
	}
	return true
}

// Track token deductions of the given priority with the given raft log index and term.
func (t *Tracker) Track(
	ctx context.Context, term uint64, index uint64, pri raftpb.Priority, tokens kvflowcontrol.Tokens,
) bool {
	if len(t.tracked[pri]) >= 1 {
		last := t.tracked[pri][len(t.tracked[pri])-1].id
		// Tracker exists in the context of a single replicaSendStream, which cannot
		// span the leader losing leadership and regaining it. So the indices must
		// advance.
		if last.index >= index {
			log.Fatalf(ctx, "expected in order tracked log indexes (%d < %d)",
				last.index, index)
			return false
		}
		if last.term > term {
			log.Fatalf(ctx, "expected in order tracked leader terms (%d < %d)",
				last.term, term)
			return false
		}
	}

	t.tracked[pri] = append(t.tracked[pri], tracked{
		id:     entryID{index: index, term: term},
		tokens: tokens,
	})

	if log.V(1) {
		log.Infof(ctx, "tracking %v flow control tokens for pri=%s stream=%s log-position=%d/%d",
			tokens, pri, t.stream, term, index)
	}
	return true
}

// Untrack all token deductions of the given priority that have indexes less
// than or equal to the one provided, per priority, and terms less than or
// equal to the leader term. evalTokensGEIndex is used to separately count the
// untracked (eval) tokens that are for indices >= evalTokensGEIndex.
func (t *Tracker) Untrack(
	av AdmittedVector, evalTokensGEIndex uint64,
) (returnedSend, returnedEval [raftpb.NumPriorities]kvflowcontrol.Tokens) {
	for pri, uptoIndex := range av.Admitted {
		var untracked int
		for n := len(t.tracked[pri]); untracked < n; untracked++ {
			deduction := t.tracked[pri][untracked]
			if deduction.id.term > av.Term ||
				deduction.id.term == av.Term && deduction.id.index > uptoIndex {
				break
			}
			returnedSend[pri] += deduction.tokens
			if deduction.id.index >= evalTokensGEIndex {
				returnedEval[pri] += deduction.tokens
			}
		}
		t.tracked[pri] = t.tracked[pri][untracked:]
	}

	return returnedSend, returnedEval
}

// UntrackAll iterates through all tracked token deductions, untracking all of them
// and returning the sum of tokens for each priority.
func (t *Tracker) UntrackAll() (returned [raftpb.NumPriorities]kvflowcontrol.Tokens) {
	for pri, deductions := range t.tracked {
		for _, deduction := range deductions {
			returned[pri] += deduction.tokens
		}
	}
	t.tracked = [raftpb.NumPriorities][]tracked{}

	return returned
}

// Inspect returns a snapshot of all tracked token deductions. It's used to
// power /inspectz-style debugging pages.
func (t *Tracker) Inspect() []kvflowinspectpb.TrackedDeduction {
	var res []kvflowinspectpb.TrackedDeduction
	for pri, deductions := range t.tracked {
		for _, deduction := range deductions {
			res = append(res, kvflowinspectpb.TrackedDeduction{
				Tokens: int64(deduction.tokens),
				RaftLogPosition: kvflowcontrolpb.RaftLogPosition{
					Index: deduction.id.index,
					Term:  deduction.id.term,
				},
				Priority: int32(RaftToAdmissionPriority(raftpb.Priority(pri))),
			})
		}
	}
	return res
}
