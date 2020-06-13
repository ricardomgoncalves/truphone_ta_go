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

// GetFamily mocks base method
func (m *MockService) GetFamily(ctx context.Context, req *GetFamilyRequest) (*GetFamilyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFamily", ctx, req)
	ret0, _ := ret[0].(*GetFamilyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFamily indicates an expected call of GetFamily
func (mr *MockServiceMockRecorder) GetFamily(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFamily", reflect.TypeOf((*MockService)(nil).GetFamily), ctx, req)
}

// ListFamilies mocks base method
func (m *MockService) ListFamilies(ctx context.Context, req *ListFamiliesRequest) (*ListFamiliesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFamilies", ctx, req)
	ret0, _ := ret[0].(*ListFamiliesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFamilies indicates an expected call of ListFamilies
func (mr *MockServiceMockRecorder) ListFamilies(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFamilies", reflect.TypeOf((*MockService)(nil).ListFamilies), ctx, req)
}

// UpdateFamily mocks base method
func (m *MockService) UpdateFamily(ctx context.Context, req *UpdateFamilyRequest) (*UpdateFamilyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFamily", ctx, req)
	ret0, _ := ret[0].(*UpdateFamilyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFamily indicates an expected call of UpdateFamily
func (mr *MockServiceMockRecorder) UpdateFamily(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFamily", reflect.TypeOf((*MockService)(nil).UpdateFamily), ctx, req)
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

// CreateMember mocks base method
func (m *MockService) CreateMember(ctx context.Context, req *CreateMemberRequest) (*CreateMemberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMember", ctx, req)
	ret0, _ := ret[0].(*CreateMemberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMember indicates an expected call of CreateMember
func (mr *MockServiceMockRecorder) CreateMember(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMember", reflect.TypeOf((*MockService)(nil).CreateMember), ctx, req)
}

// GetMember mocks base method
func (m *MockService) GetMember(ctx context.Context, req *GetMemberRequest) (*GetMemberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMember", ctx, req)
	ret0, _ := ret[0].(*GetMemberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMember indicates an expected call of GetMember
func (mr *MockServiceMockRecorder) GetMember(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMember", reflect.TypeOf((*MockService)(nil).GetMember), ctx, req)
}

// ListMembers mocks base method
func (m *MockService) ListMembers(ctx context.Context, req *ListMembersRequest) (*ListMembersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMembers", ctx, req)
	ret0, _ := ret[0].(*ListMembersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembers indicates an expected call of ListMembers
func (mr *MockServiceMockRecorder) ListMembers(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockService)(nil).ListMembers), ctx, req)
}

// UpdateMember mocks base method
func (m *MockService) UpdateMember(ctx context.Context, req *UpdateMemberRequest) (*UpdateMemberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMember", ctx, req)
	ret0, _ := ret[0].(*UpdateMemberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMember indicates an expected call of UpdateMember
func (mr *MockServiceMockRecorder) UpdateMember(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMember", reflect.TypeOf((*MockService)(nil).UpdateMember), ctx, req)
}

// DeleteMember mocks base method
func (m *MockService) DeleteMember(ctx context.Context, req *DeleteMemberRequest) (*DeleteMemberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMember", ctx, req)
	ret0, _ := ret[0].(*DeleteMemberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMember indicates an expected call of DeleteMember
func (mr *MockServiceMockRecorder) DeleteMember(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMember", reflect.TypeOf((*MockService)(nil).DeleteMember), ctx, req)
}
