// Code generated by MockGen. DO NOT EDIT.
// Source: reader.go
//
// Generated by this command:
//
//	mockgen -source=reader.go -destination=mocks/reader_mock.go -package=mocks InventoryReader
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInventoryReader is a mock of InventoryReader interface.
type MockInventoryReader struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryReaderMockRecorder
}

// MockInventoryReaderMockRecorder is the mock recorder for MockInventoryReader.
type MockInventoryReaderMockRecorder struct {
	mock *MockInventoryReader
}

// NewMockInventoryReader creates a new mock instance.
func NewMockInventoryReader(ctrl *gomock.Controller) *MockInventoryReader {
	mock := &MockInventoryReader{ctrl: ctrl}
	mock.recorder = &MockInventoryReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryReader) EXPECT() *MockInventoryReaderMockRecorder {
	return m.recorder
}

// ReadInventory mocks base method.
func (m *MockInventoryReader) ReadInventory(filePath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadInventory", filePath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadInventory indicates an expected call of ReadInventory.
func (mr *MockInventoryReaderMockRecorder) ReadInventory(filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInventory", reflect.TypeOf((*MockInventoryReader)(nil).ReadInventory), filePath)
}