// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nokamoto/demo20-apis/cloud/rdb/v1alpha (interfaces: RdbClient)

// Package mockrdb is a generated GoMock package.
package mockrdb

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha "github.com/nokamoto/demo20-apis/cloud/rdb/v1alpha"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockRdbClient is a mock of RdbClient interface
type MockRdbClient struct {
	ctrl     *gomock.Controller
	recorder *MockRdbClientMockRecorder
}

// MockRdbClientMockRecorder is the mock recorder for MockRdbClient
type MockRdbClientMockRecorder struct {
	mock *MockRdbClient
}

// NewMockRdbClient creates a new mock instance
func NewMockRdbClient(ctrl *gomock.Controller) *MockRdbClient {
	mock := &MockRdbClient{ctrl: ctrl}
	mock.recorder = &MockRdbClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRdbClient) EXPECT() *MockRdbClientMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method
func (m *MockRdbClient) CreateCluster(arg0 context.Context, arg1 *v1alpha.CreateClusterRequest, arg2 ...grpc.CallOption) (*v1alpha.Cluster, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCluster", varargs...)
	ret0, _ := ret[0].(*v1alpha.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockRdbClientMockRecorder) CreateCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockRdbClient)(nil).CreateCluster), varargs...)
}

// DeleteCluster mocks base method
func (m *MockRdbClient) DeleteCluster(arg0 context.Context, arg1 *v1alpha.DeleteClusterRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCluster", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockRdbClientMockRecorder) DeleteCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockRdbClient)(nil).DeleteCluster), varargs...)
}

// GetCluster mocks base method
func (m *MockRdbClient) GetCluster(arg0 context.Context, arg1 *v1alpha.GetClusterRequest, arg2 ...grpc.CallOption) (*v1alpha.Cluster, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCluster", varargs...)
	ret0, _ := ret[0].(*v1alpha.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockRdbClientMockRecorder) GetCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockRdbClient)(nil).GetCluster), varargs...)
}

// ListCluster mocks base method
func (m *MockRdbClient) ListCluster(arg0 context.Context, arg1 *v1alpha.ListClusterRequest, arg2 ...grpc.CallOption) (*v1alpha.ListClusterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCluster", varargs...)
	ret0, _ := ret[0].(*v1alpha.ListClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCluster indicates an expected call of ListCluster
func (mr *MockRdbClientMockRecorder) ListCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCluster", reflect.TypeOf((*MockRdbClient)(nil).ListCluster), varargs...)
}

// UpdateCluster mocks base method
func (m *MockRdbClient) UpdateCluster(arg0 context.Context, arg1 *v1alpha.UpdateClusterRequest, arg2 ...grpc.CallOption) (*v1alpha.Cluster, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCluster", varargs...)
	ret0, _ := ret[0].(*v1alpha.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCluster indicates an expected call of UpdateCluster
func (mr *MockRdbClientMockRecorder) UpdateCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockRdbClient)(nil).UpdateCluster), varargs...)
}
