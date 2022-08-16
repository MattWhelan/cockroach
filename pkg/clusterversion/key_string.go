// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[invalidVersionKey - -1]
	_ = x[V21_2-0]
	_ = x[Start22_1-1]
	_ = x[ProbeRequest-2]
	_ = x[EnableSpanConfigStore-3]
	_ = x[EnableNewStoreRebalancer-4]
	_ = x[V22_1-5]
	_ = x[Start22_2-6]
	_ = x[LocalTimestamps-7]
	_ = x[PebbleFormatSplitUserKeysMarkedCompacted-8]
	_ = x[EnsurePebbleFormatVersionRangeKeys-9]
	_ = x[EnablePebbleFormatVersionRangeKeys-10]
	_ = x[TrigramInvertedIndexes-11]
	_ = x[RemoveGrantPrivilege-12]
	_ = x[MVCCRangeTombstones-13]
	_ = x[UpgradeSequenceToBeReferencedByID-14]
	_ = x[SampledStmtDiagReqs-15]
	_ = x[AddSSTableTombstones-16]
	_ = x[SystemPrivilegesTable-17]
	_ = x[EnablePredicateProjectionChangefeed-18]
	_ = x[AlterSystemSQLInstancesAddLocality-19]
	_ = x[SystemExternalConnectionsTable-20]
	_ = x[AlterSystemStatementStatisticsAddIndexRecommendations-21]
	_ = x[RoleIDSequence-22]
	_ = x[AddSystemUserIDColumn-23]
	_ = x[SystemUsersIDColumnIsBackfilled-24]
	_ = x[SetSystemUsersUserIDColumnNotNull-25]
	_ = x[SQLSchemaTelemetryScheduledJobs-26]
	_ = x[SchemaChangeSupportsCreateFunction-27]
	_ = x[DeleteRequestReturnKey-28]
	_ = x[PebbleFormatPrePebblev1Marked-29]
	_ = x[RoleOptionsTableHasIDColumn-30]
	_ = x[RoleOptionsIDColumnIsBackfilled-31]
	_ = x[SetRoleOptionsUserIDColumnNotNull-32]
	_ = x[UseDelRangeInGCJob-33]
	_ = x[WaitedForDelRangeInGCJob-34]
	_ = x[RangefeedUseOneStreamPerNode-35]
	_ = x[NoNonMVCCAddSSTable-36]
	_ = x[GCHintInReplicaState-37]
	_ = x[UpdateInvalidColumnIDsInSequenceBackReferences-38]
}

const _Key_name = "invalidVersionKeyV21_2Start22_1ProbeRequestEnableSpanConfigStoreEnableNewStoreRebalancerV22_1Start22_2LocalTimestampsPebbleFormatSplitUserKeysMarkedCompactedEnsurePebbleFormatVersionRangeKeysEnablePebbleFormatVersionRangeKeysTrigramInvertedIndexesRemoveGrantPrivilegeMVCCRangeTombstonesUpgradeSequenceToBeReferencedByIDSampledStmtDiagReqsAddSSTableTombstonesSystemPrivilegesTableEnablePredicateProjectionChangefeedAlterSystemSQLInstancesAddLocalitySystemExternalConnectionsTableAlterSystemStatementStatisticsAddIndexRecommendationsRoleIDSequenceAddSystemUserIDColumnSystemUsersIDColumnIsBackfilledSetSystemUsersUserIDColumnNotNullSQLSchemaTelemetryScheduledJobsSchemaChangeSupportsCreateFunctionDeleteRequestReturnKeyPebbleFormatPrePebblev1MarkedRoleOptionsTableHasIDColumnRoleOptionsIDColumnIsBackfilledSetRoleOptionsUserIDColumnNotNullUseDelRangeInGCJobWaitedForDelRangeInGCJobRangefeedUseOneStreamPerNodeNoNonMVCCAddSSTableGCHintInReplicaStateUpdateInvalidColumnIDsInSequenceBackReferences"

var _Key_index = [...]uint16{0, 17, 22, 31, 43, 64, 88, 93, 102, 117, 157, 191, 225, 247, 267, 286, 319, 338, 358, 379, 414, 448, 478, 531, 545, 566, 597, 630, 661, 695, 717, 746, 773, 804, 837, 855, 879, 907, 926, 946, 992}

func (i Key) String() string {
	i -= -1
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
