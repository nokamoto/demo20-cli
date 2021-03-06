// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nokamoto/demo20-apis/cloud/compute/v1alpha (interfaces: ComputeClient)

// Package mockcompute is a generated GoMock package.
package mockcompute

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha "github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockComputeClient is a mock of ComputeClient interface
type MockComputeClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeClientMockRecorder
}

// MockComputeClientMockRecorder is the mock recorder for MockComputeClient
type MockComputeClientMockRecorder struct {
	mock *MockComputeClient
}

// NewMockComputeClient creates a new mock instance
func NewMockComputeClient(ctrl *gomock.Controller) *MockComputeClient {
	mock := &MockComputeClient{ctrl: ctrl}
	mock.recorder = &MockComputeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockComputeClient) EXPECT() *MockComputeClientMockRecorder {
	return m.recorder
}

// CreateInstance mocks base method
func (m *MockComputeClient) CreateInstance(arg0 context.Context, arg1 *v1alpha.CreateInstanceRequest, arg2 ...grpc.CallOption) (*v1alpha.Instance, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateInstance", varargs...)
	ret0, _ := ret[0].(*v1alpha.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInstance indicates an expected call of CreateInstance
func (mr *MockComputeClientMockRecorder) CreateInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockComputeClient)(nil).CreateInstance), varargs...)
}

// DeleteInstance mocks base method
func (m *MockComputeClient) DeleteInstance(arg0 context.Context, arg1 *v1alpha.DeleteInstanceRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteInstance", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteInstance indicates an expected call of DeleteInstance
func (mr *MockComputeClientMockRecorder) DeleteInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstance", reflect.TypeOf((*MockComputeClient)(nil).DeleteInstance), varargs...)
}

// GetInstance mocks base method
func (m *MockComputeClient) GetInstance(arg0 context.Context, arg1 *v1alpha.GetInstanceRequest, arg2 ...grpc.CallOption) (*v1alpha.Instance, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstance", varargs...)
	ret0, _ := ret[0].(*v1alpha.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance
func (mr *MockComputeClientMockRecorder) GetInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockComputeClient)(nil).GetInstance), varargs...)
}

// ListInstance mocks base method
func (m *MockComputeClient) ListInstance(arg0 context.Context, arg1 *v1alpha.ListInstanceRequest, arg2 ...grpc.CallOption) (*v1alpha.ListInstanceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInstance", varargs...)
	ret0, _ := ret[0].(*v1alpha.ListInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstance indicates an expected call of ListInstance
func (mr *MockComputeClientMockRecorder) ListInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstance", reflect.TypeOf((*MockComputeClient)(nil).ListInstance), varargs...)
}
