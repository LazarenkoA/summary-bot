// Code generated by MockGen. DO NOT EDIT.
// Source: bot.go

// Package mock_bot is a generated GoMock package.
package mock_bot

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIAdapter is a mock of IAdapter interface.
type MockIAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockIAdapterMockRecorder
}

// MockIAdapterMockRecorder is the mock recorder for MockIAdapter.
type MockIAdapterMockRecorder struct {
	mock *MockIAdapter
}

// NewMockIAdapter creates a new mock instance.
func NewMockIAdapter(ctrl *gomock.Controller) *MockIAdapter {
	mock := &MockIAdapter{ctrl: ctrl}
	mock.recorder = &MockIAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAdapter) EXPECT() *MockIAdapterMockRecorder {
	return m.recorder
}

// AppendTimeData mocks base method.
func (m *MockIAdapter) AppendTimeData(key string, t time.Time, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendTimeData", key, t, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendTimeData indicates an expected call of AppendTimeData.
func (mr *MockIAdapterMockRecorder) AppendTimeData(key, t, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendTimeData", reflect.TypeOf((*MockIAdapter)(nil).AppendTimeData), key, t, data)
}

// DeleteMessageData mocks base method.
func (m *MockIAdapter) DeleteMessageData(key string, t time.Time, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessageData", key, t, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessageData indicates an expected call of DeleteMessageData.
func (mr *MockIAdapterMockRecorder) DeleteMessageData(key, t, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessageData", reflect.TypeOf((*MockIAdapter)(nil).DeleteMessageData), key, t, data)
}

// DeleteMessageDataByTime mocks base method.
func (m *MockIAdapter) DeleteMessageDataByTime(key string, t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessageDataByTime", key, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessageDataByTime indicates an expected call of DeleteMessageDataByTime.
func (mr *MockIAdapterMockRecorder) DeleteMessageDataByTime(key, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessageDataByTime", reflect.TypeOf((*MockIAdapter)(nil).DeleteMessageDataByTime), key, t)
}

// GetMessageData mocks base method.
func (m *MockIAdapter) GetMessageData(key string, tstart, tfinish time.Time) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageData", key, tstart, tfinish)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageData indicates an expected call of GetMessageData.
func (mr *MockIAdapterMockRecorder) GetMessageData(key, tstart, tfinish interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageData", reflect.TypeOf((*MockIAdapter)(nil).GetMessageData), key, tstart, tfinish)
}

// GetMessageDataForClear mocks base method.
func (m *MockIAdapter) GetMessageDataForClear(key string, tfinish time.Time) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageDataForClear", key, tfinish)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageDataForClear indicates an expected call of GetMessageDataForClear.
func (mr *MockIAdapterMockRecorder) GetMessageDataForClear(key, tfinish interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageDataForClear", reflect.TypeOf((*MockIAdapter)(nil).GetMessageDataForClear), key, tfinish)
}