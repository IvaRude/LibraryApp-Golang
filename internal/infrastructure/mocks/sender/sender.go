// Code generated by MockGen. DO NOT EDIT.
// Source: ./sender.go

// Package mock_sender is a generated GoMock package.
package mock_sender

import (
	models "homework-3/internal/pkg/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSender is a mock of Sender interface.
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender.
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance.
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// SendAsyncMessage mocks base method.
func (m *MockSender) SendAsyncMessage(message *models.HandlerMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAsyncMessage", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAsyncMessage indicates an expected call of SendAsyncMessage.
func (mr *MockSenderMockRecorder) SendAsyncMessage(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAsyncMessage", reflect.TypeOf((*MockSender)(nil).SendAsyncMessage), message)
}

// SendMessage mocks base method.
func (m *MockSender) SendMessage(message *models.HandlerMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockSenderMockRecorder) SendMessage(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockSender)(nil).SendMessage), message)
}

// SendMessages mocks base method.
func (m *MockSender) SendMessages(messages []*models.HandlerMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessages", messages)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessages indicates an expected call of SendMessages.
func (mr *MockSenderMockRecorder) SendMessages(messages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessages", reflect.TypeOf((*MockSender)(nil).SendMessages), messages)
}
