// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockstore is a mock of store interface
type Mockstore struct {
	ctrl     *gomock.Controller
	recorder *MockstoreMockRecorder
}

// MockstoreMockRecorder is the mock recorder for Mockstore
type MockstoreMockRecorder struct {
	mock *Mockstore
}

// NewMockstore creates a new mock instance
func NewMockstore(ctrl *gomock.Controller) *Mockstore {
	mock := &Mockstore{ctrl: ctrl}
	mock.recorder = &MockstoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockstore) EXPECT() *MockstoreMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *Mockstore) Set(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set
func (mr *MockstoreMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*Mockstore)(nil).Set), arg0, arg1)
}

// Remove mocks base method
func (m *Mockstore) Remove(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove
func (mr *MockstoreMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*Mockstore)(nil).Remove), arg0)
}
