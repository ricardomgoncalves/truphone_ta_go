// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/service.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateFamily mocks base method
func (m *MockService) CreateFamily(ctx context.Context, req *CreateFamilyRequest) (*CreateFamilyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFamily", ctx, req)
	ret0, _ := ret[0].(*CreateFamilyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFamily indicates an expected call of CreateFamily
func (mr *MockServiceMockRecorder) CreateFamily(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFamily", reflect.TypeOf((*MockService)(nil).CreateFamily), ctx, req)
}

// DeleteFamily mocks base method
func (m *MockService) DeleteFamily(ctx context.Context, req *DeleteFamilyRequest) (*DeleteFamilyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFamily", ctx, req)
	ret0, _ := ret[0].(*DeleteFamilyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFamily indicates an expected call of DeleteFamily
func (mr *MockServiceMockRecorder) DeleteFamily(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFamily", reflect.TypeOf((*MockService)(nil).DeleteFamily), ctx, req)
}