// Code generated by MockGen. DO NOT EDIT.
// Source: discovery.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockdiscovery is a mock of discovery interface
type Mockdiscovery struct {
	ctrl     *gomock.Controller
	recorder *MockdiscoveryMockRecorder
}

// MockdiscoveryMockRecorder is the mock recorder for Mockdiscovery
type MockdiscoveryMockRecorder struct {
	mock *Mockdiscovery
}

// NewMockdiscovery creates a new mock instance
func NewMockdiscovery(ctrl *gomock.Controller) *Mockdiscovery {
	mock := &Mockdiscovery{ctrl: ctrl}
	mock.recorder = &MockdiscoveryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockdiscovery) EXPECT() *MockdiscoveryMockRecorder {
	return m.recorder
}

// Endpoints mocks base method
func (m *Mockdiscovery) Endpoints() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoints")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Endpoints indicates an expected call of Endpoints
func (mr *MockdiscoveryMockRecorder) Endpoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoints", reflect.TypeOf((*Mockdiscovery)(nil).Endpoints))
}
