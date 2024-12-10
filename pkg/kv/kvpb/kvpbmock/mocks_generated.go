// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cockroachdb/cockroach/pkg/kv/kvpb (interfaces: InternalClient,Internal_MuxRangeFeedClient)

// Package kvpbmock is a generated GoMock package.
package kvpbmock

import (
	context "context"
	reflect "reflect"

	kvpb "github.com/cockroachdb/cockroach/pkg/kv/kvpb"
	roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockInternalClient is a mock of InternalClient interface.
type MockInternalClient struct {
	ctrl     *gomock.Controller
	recorder *MockInternalClientMockRecorder
}

// MockInternalClientMockRecorder is the mock recorder for MockInternalClient.
type MockInternalClientMockRecorder struct {
	mock *MockInternalClient
}

// NewMockInternalClient creates a new mock instance.
func NewMockInternalClient(ctrl *gomock.Controller) *MockInternalClient {
	mock := &MockInternalClient{ctrl: ctrl}
	mock.recorder = &MockInternalClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInternalClient) EXPECT() *MockInternalClientMockRecorder {
	return m.recorder
}

// Batch mocks base method.
func (m *MockInternalClient) Batch(arg0 context.Context, arg1 *kvpb.BatchRequest, arg2 ...grpc.CallOption) (*kvpb.BatchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Batch", varargs...)
	ret0, _ := ret[0].(*kvpb.BatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Batch indicates an expected call of Batch.
func (mr *MockInternalClientMockRecorder) Batch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batch", reflect.TypeOf((*MockInternalClient)(nil).Batch), varargs...)
}

// BatchStream mocks base method.
func (m *MockInternalClient) BatchStream(arg0 context.Context, arg1 ...grpc.CallOption) (kvpb.Internal_BatchStreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchStream", varargs...)
	ret0, _ := ret[0].(kvpb.Internal_BatchStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchStream indicates an expected call of BatchStream.
func (mr *MockInternalClientMockRecorder) BatchStream(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchStream", reflect.TypeOf((*MockInternalClient)(nil).BatchStream), varargs...)
}

// GetAllSystemSpanConfigsThatApply mocks base method.
func (m *MockInternalClient) GetAllSystemSpanConfigsThatApply(arg0 context.Context, arg1 *roachpb.GetAllSystemSpanConfigsThatApplyRequest, arg2 ...grpc.CallOption) (*roachpb.GetAllSystemSpanConfigsThatApplyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllSystemSpanConfigsThatApply", varargs...)
	ret0, _ := ret[0].(*roachpb.GetAllSystemSpanConfigsThatApplyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSystemSpanConfigsThatApply indicates an expected call of GetAllSystemSpanConfigsThatApply.
func (mr *MockInternalClientMockRecorder) GetAllSystemSpanConfigsThatApply(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSystemSpanConfigsThatApply", reflect.TypeOf((*MockInternalClient)(nil).GetAllSystemSpanConfigsThatApply), varargs...)
}

// GetRangeDescriptors mocks base method.
func (m *MockInternalClient) GetRangeDescriptors(arg0 context.Context, arg1 *kvpb.GetRangeDescriptorsRequest, arg2 ...grpc.CallOption) (kvpb.Internal_GetRangeDescriptorsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRangeDescriptors", varargs...)
	ret0, _ := ret[0].(kvpb.Internal_GetRangeDescriptorsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRangeDescriptors indicates an expected call of GetRangeDescriptors.
func (mr *MockInternalClientMockRecorder) GetRangeDescriptors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRangeDescriptors", reflect.TypeOf((*MockInternalClient)(nil).GetRangeDescriptors), varargs...)
}

// GetSpanConfigs mocks base method.
func (m *MockInternalClient) GetSpanConfigs(arg0 context.Context, arg1 *roachpb.GetSpanConfigsRequest, arg2 ...grpc.CallOption) (*roachpb.GetSpanConfigsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSpanConfigs", varargs...)
	ret0, _ := ret[0].(*roachpb.GetSpanConfigsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpanConfigs indicates an expected call of GetSpanConfigs.
func (mr *MockInternalClientMockRecorder) GetSpanConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpanConfigs", reflect.TypeOf((*MockInternalClient)(nil).GetSpanConfigs), varargs...)
}

// GossipSubscription mocks base method.
func (m *MockInternalClient) GossipSubscription(arg0 context.Context, arg1 *kvpb.GossipSubscriptionRequest, arg2 ...grpc.CallOption) (kvpb.Internal_GossipSubscriptionClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GossipSubscription", varargs...)
	ret0, _ := ret[0].(kvpb.Internal_GossipSubscriptionClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GossipSubscription indicates an expected call of GossipSubscription.
func (mr *MockInternalClientMockRecorder) GossipSubscription(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GossipSubscription", reflect.TypeOf((*MockInternalClient)(nil).GossipSubscription), varargs...)
}

// Join mocks base method.
func (m *MockInternalClient) Join(arg0 context.Context, arg1 *kvpb.JoinNodeRequest, arg2 ...grpc.CallOption) (*kvpb.JoinNodeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Join", varargs...)
	ret0, _ := ret[0].(*kvpb.JoinNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Join indicates an expected call of Join.
func (mr *MockInternalClientMockRecorder) Join(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockInternalClient)(nil).Join), varargs...)
}

// MuxRangeFeed mocks base method.
func (m *MockInternalClient) MuxRangeFeed(arg0 context.Context, arg1 ...grpc.CallOption) (kvpb.Internal_MuxRangeFeedClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MuxRangeFeed", varargs...)
	ret0, _ := ret[0].(kvpb.Internal_MuxRangeFeedClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MuxRangeFeed indicates an expected call of MuxRangeFeed.
func (mr *MockInternalClientMockRecorder) MuxRangeFeed(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MuxRangeFeed", reflect.TypeOf((*MockInternalClient)(nil).MuxRangeFeed), varargs...)
}

// RangeLookup mocks base method.
func (m *MockInternalClient) RangeLookup(arg0 context.Context, arg1 *kvpb.RangeLookupRequest, arg2 ...grpc.CallOption) (*kvpb.RangeLookupResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RangeLookup", varargs...)
	ret0, _ := ret[0].(*kvpb.RangeLookupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RangeLookup indicates an expected call of RangeLookup.
func (mr *MockInternalClientMockRecorder) RangeLookup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeLookup", reflect.TypeOf((*MockInternalClient)(nil).RangeLookup), varargs...)
}

// ResetQuorum mocks base method.
func (m *MockInternalClient) ResetQuorum(arg0 context.Context, arg1 *kvpb.ResetQuorumRequest, arg2 ...grpc.CallOption) (*kvpb.ResetQuorumResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetQuorum", varargs...)
	ret0, _ := ret[0].(*kvpb.ResetQuorumResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetQuorum indicates an expected call of ResetQuorum.
func (mr *MockInternalClientMockRecorder) ResetQuorum(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetQuorum", reflect.TypeOf((*MockInternalClient)(nil).ResetQuorum), varargs...)
}

// SpanConfigConformance mocks base method.
func (m *MockInternalClient) SpanConfigConformance(arg0 context.Context, arg1 *roachpb.SpanConfigConformanceRequest, arg2 ...grpc.CallOption) (*roachpb.SpanConfigConformanceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SpanConfigConformance", varargs...)
	ret0, _ := ret[0].(*roachpb.SpanConfigConformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpanConfigConformance indicates an expected call of SpanConfigConformance.
func (mr *MockInternalClientMockRecorder) SpanConfigConformance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpanConfigConformance", reflect.TypeOf((*MockInternalClient)(nil).SpanConfigConformance), varargs...)
}

// TenantSettings mocks base method.
func (m *MockInternalClient) TenantSettings(arg0 context.Context, arg1 *kvpb.TenantSettingsRequest, arg2 ...grpc.CallOption) (kvpb.Internal_TenantSettingsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TenantSettings", varargs...)
	ret0, _ := ret[0].(kvpb.Internal_TenantSettingsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TenantSettings indicates an expected call of TenantSettings.
func (mr *MockInternalClientMockRecorder) TenantSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantSettings", reflect.TypeOf((*MockInternalClient)(nil).TenantSettings), varargs...)
}

// TokenBucket mocks base method.
func (m *MockInternalClient) TokenBucket(arg0 context.Context, arg1 *kvpb.TokenBucketRequest, arg2 ...grpc.CallOption) (*kvpb.TokenBucketResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TokenBucket", varargs...)
	ret0, _ := ret[0].(*kvpb.TokenBucketResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenBucket indicates an expected call of TokenBucket.
func (mr *MockInternalClientMockRecorder) TokenBucket(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenBucket", reflect.TypeOf((*MockInternalClient)(nil).TokenBucket), varargs...)
}

// UpdateSpanConfigs mocks base method.
func (m *MockInternalClient) UpdateSpanConfigs(arg0 context.Context, arg1 *roachpb.UpdateSpanConfigsRequest, arg2 ...grpc.CallOption) (*roachpb.UpdateSpanConfigsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSpanConfigs", varargs...)
	ret0, _ := ret[0].(*roachpb.UpdateSpanConfigsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSpanConfigs indicates an expected call of UpdateSpanConfigs.
func (mr *MockInternalClientMockRecorder) UpdateSpanConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSpanConfigs", reflect.TypeOf((*MockInternalClient)(nil).UpdateSpanConfigs), varargs...)
}

// MockInternal_MuxRangeFeedClient is a mock of Internal_MuxRangeFeedClient interface.
type MockInternal_MuxRangeFeedClient struct {
	ctrl     *gomock.Controller
	recorder *MockInternal_MuxRangeFeedClientMockRecorder
}

// MockInternal_MuxRangeFeedClientMockRecorder is the mock recorder for MockInternal_MuxRangeFeedClient.
type MockInternal_MuxRangeFeedClientMockRecorder struct {
	mock *MockInternal_MuxRangeFeedClient
}

// NewMockInternal_MuxRangeFeedClient creates a new mock instance.
func NewMockInternal_MuxRangeFeedClient(ctrl *gomock.Controller) *MockInternal_MuxRangeFeedClient {
	mock := &MockInternal_MuxRangeFeedClient{ctrl: ctrl}
	mock.recorder = &MockInternal_MuxRangeFeedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInternal_MuxRangeFeedClient) EXPECT() *MockInternal_MuxRangeFeedClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockInternal_MuxRangeFeedClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockInternal_MuxRangeFeedClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockInternal_MuxRangeFeedClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockInternal_MuxRangeFeedClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockInternal_MuxRangeFeedClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockInternal_MuxRangeFeedClient)(nil).Context))
}

// Header mocks base method.
func (m *MockInternal_MuxRangeFeedClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockInternal_MuxRangeFeedClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockInternal_MuxRangeFeedClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockInternal_MuxRangeFeedClient) Recv() (*kvpb.MuxRangeFeedEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*kvpb.MuxRangeFeedEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockInternal_MuxRangeFeedClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockInternal_MuxRangeFeedClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockInternal_MuxRangeFeedClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockInternal_MuxRangeFeedClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockInternal_MuxRangeFeedClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockInternal_MuxRangeFeedClient) Send(arg0 *kvpb.RangeFeedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockInternal_MuxRangeFeedClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockInternal_MuxRangeFeedClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockInternal_MuxRangeFeedClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockInternal_MuxRangeFeedClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockInternal_MuxRangeFeedClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockInternal_MuxRangeFeedClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockInternal_MuxRangeFeedClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockInternal_MuxRangeFeedClient)(nil).Trailer))
}
