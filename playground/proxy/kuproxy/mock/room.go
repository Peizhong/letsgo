// Code generated by MockGen. DO NOT EDIT.
// Source: room.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockroom is a mock of room interface
type Mockroom struct {
	ctrl     *gomock.Controller
	recorder *MockroomMockRecorder
}

// MockroomMockRecorder is the mock recorder for Mockroom
type MockroomMockRecorder struct {
	mock *Mockroom
}

// NewMockroom creates a new mock instance
func NewMockroom(ctrl *gomock.Controller) *Mockroom {
	mock := &Mockroom{ctrl: ctrl}
	mock.recorder = &MockroomMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockroom) EXPECT() *MockroomMockRecorder {
	return m.recorder
}

// Join mocks base method
func (m *Mockroom) Join(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Join", arg0)
}

// Join indicates an expected call of Join
func (mr *MockroomMockRecorder) Join(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*Mockroom)(nil).Join), arg0)
}

// Leave mocks base method
func (m *Mockroom) Leave(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Leave", arg0)
}

// Leave indicates an expected call of Leave
func (mr *MockroomMockRecorder) Leave(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*Mockroom)(nil).Leave), arg0)
}
