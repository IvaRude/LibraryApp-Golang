// Code generated by MockGen. DO NOT EDIT.
// Source: ./receiver_handler.go

// Package mock_receiver_handler is a generated GoMock package.
package mock_receiver_handler

import (
	reflect "reflect"

	sarama "github.com/IBM/sarama"
	gomock "github.com/golang/mock/gomock"
)

// MockKafkaReceiverHandler is a mock of KafkaReceiverHandler interface.
type MockKafkaReceiverHandler struct {
	ctrl     *gomock.Controller
	recorder *MockKafkaReceiverHandlerMockRecorder
}

// MockKafkaReceiverHandlerMockRecorder is the mock recorder for MockKafkaReceiverHandler.
type MockKafkaReceiverHandlerMockRecorder struct {
	mock *MockKafkaReceiverHandler
}

// NewMockKafkaReceiverHandler creates a new mock instance.
func NewMockKafkaReceiverHandler(ctrl *gomock.Controller) *MockKafkaReceiverHandler {
	mock := &MockKafkaReceiverHandler{ctrl: ctrl}
	mock.recorder = &MockKafkaReceiverHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKafkaReceiverHandler) EXPECT() *MockKafkaReceiverHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockKafkaReceiverHandler) Handle(message *sarama.ConsumerMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", message)
}

// Handle indicates an expected call of Handle.
func (mr *MockKafkaReceiverHandlerMockRecorder) Handle(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockKafkaReceiverHandler)(nil).Handle), message)
}
