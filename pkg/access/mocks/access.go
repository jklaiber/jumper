// Code generated by MockGen. DO NOT EDIT.
// Source: access.go
//
// Generated by this command:
//
//	mockgen -source=access.go -destination=mocks/access.go -package=mocks Configurator
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	ssh "golang.org/x/crypto/ssh"
)

// MockConfigurator is a mock of Configurator interface.
type MockConfigurator struct {
	ctrl     *gomock.Controller
	recorder *MockConfiguratorMockRecorder
}

// MockConfiguratorMockRecorder is the mock recorder for MockConfigurator.
type MockConfiguratorMockRecorder struct {
	mock *MockConfigurator
}

// NewMockConfigurator creates a new mock instance.
func NewMockConfigurator(ctrl *gomock.Controller) *MockConfigurator {
	mock := &MockConfigurator{ctrl: ctrl}
	mock.recorder = &MockConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurator) EXPECT() *MockConfiguratorMockRecorder {
	return m.recorder
}

// BuildClientConfig mocks base method.
func (m *MockConfigurator) BuildClientConfig() (*ssh.ClientConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildClientConfig")
	ret0, _ := ret[0].(*ssh.ClientConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildClientConfig indicates an expected call of BuildClientConfig.
func (mr *MockConfiguratorMockRecorder) BuildClientConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildClientConfig", reflect.TypeOf((*MockConfigurator)(nil).BuildClientConfig))
}

// GetAddress mocks base method.
func (m *MockConfigurator) GetAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockConfiguratorMockRecorder) GetAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockConfigurator)(nil).GetAddress))
}

// GetPassword mocks base method.
func (m *MockConfigurator) GetPassword() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockConfiguratorMockRecorder) GetPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockConfigurator)(nil).GetPassword))
}

// GetPort mocks base method.
func (m *MockConfigurator) GetPort() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPort indicates an expected call of GetPort.
func (mr *MockConfiguratorMockRecorder) GetPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockConfigurator)(nil).GetPort))
}

// GetSshAgent mocks base method.
func (m *MockConfigurator) GetSshAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSshAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetSshAgent indicates an expected call of GetSshAgent.
func (mr *MockConfiguratorMockRecorder) GetSshAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSshAgent", reflect.TypeOf((*MockConfigurator)(nil).GetSshAgent))
}

// GetSshAgentForwarding mocks base method.
func (m *MockConfigurator) GetSshAgentForwarding() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSshAgentForwarding")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetSshAgentForwarding indicates an expected call of GetSshAgentForwarding.
func (mr *MockConfiguratorMockRecorder) GetSshAgentForwarding() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSshAgentForwarding", reflect.TypeOf((*MockConfigurator)(nil).GetSshAgentForwarding))
}

// GetSshKey mocks base method.
func (m *MockConfigurator) GetSshKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSshKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSshKey indicates an expected call of GetSshKey.
func (mr *MockConfiguratorMockRecorder) GetSshKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSshKey", reflect.TypeOf((*MockConfigurator)(nil).GetSshKey))
}

// GetUsername mocks base method.
func (m *MockConfigurator) GetUsername() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsername")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUsername indicates an expected call of GetUsername.
func (mr *MockConfiguratorMockRecorder) GetUsername() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsername", reflect.TypeOf((*MockConfigurator)(nil).GetUsername))
}

// SetPort mocks base method.
func (m *MockConfigurator) SetPort(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPort", arg0)
}

// SetPort indicates an expected call of SetPort.
func (mr *MockConfiguratorMockRecorder) SetPort(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPort", reflect.TypeOf((*MockConfigurator)(nil).SetPort), arg0)
}

// getPublicKeyFile mocks base method.
func (m *MockConfigurator) getPublicKeyFile(file string) (ssh.AuthMethod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPublicKeyFile", file)
	ret0, _ := ret[0].(ssh.AuthMethod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getPublicKeyFile indicates an expected call of getPublicKeyFile.
func (mr *MockConfiguratorMockRecorder) getPublicKeyFile(file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPublicKeyFile", reflect.TypeOf((*MockConfigurator)(nil).getPublicKeyFile), file)
}

// getSshAgent mocks base method.
func (m *MockConfigurator) getSshAgent() (ssh.AuthMethod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getSshAgent")
	ret0, _ := ret[0].(ssh.AuthMethod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getSshAgent indicates an expected call of getSshAgent.
func (mr *MockConfiguratorMockRecorder) getSshAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getSshAgent", reflect.TypeOf((*MockConfigurator)(nil).getSshAgent))
}

// parsePrivateKeyWithPassphrase mocks base method.
func (m *MockConfigurator) parsePrivateKeyWithPassphrase(file string, buffer []byte) (ssh.AuthMethod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "parsePrivateKeyWithPassphrase", file, buffer)
	ret0, _ := ret[0].(ssh.AuthMethod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// parsePrivateKeyWithPassphrase indicates an expected call of parsePrivateKeyWithPassphrase.
func (mr *MockConfiguratorMockRecorder) parsePrivateKeyWithPassphrase(file, buffer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "parsePrivateKeyWithPassphrase", reflect.TypeOf((*MockConfigurator)(nil).parsePrivateKeyWithPassphrase), file, buffer)
}

// readSSHKeyPassphrase mocks base method.
func (m *MockConfigurator) readSSHKeyPassphrase(file string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "readSSHKeyPassphrase", file)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// readSSHKeyPassphrase indicates an expected call of readSSHKeyPassphrase.
func (mr *MockConfiguratorMockRecorder) readSSHKeyPassphrase(file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "readSSHKeyPassphrase", reflect.TypeOf((*MockConfigurator)(nil).readSSHKeyPassphrase), file)
}