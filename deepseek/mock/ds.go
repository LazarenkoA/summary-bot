// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-deepseek/deepseek (interfaces: Client)

// Package mock_deepseek is a generated GoMock package.
package mock_deepseek

import (
	context "context"
	reflect "reflect"

	request "github.com/go-deepseek/deepseek/request"
	response "github.com/go-deepseek/deepseek/response"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CallChatCompletionsChat mocks base method.
func (m *MockClient) CallChatCompletionsChat(arg0 context.Context, arg1 *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallChatCompletionsChat", arg0, arg1)
	ret0, _ := ret[0].(*response.ChatCompletionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallChatCompletionsChat indicates an expected call of CallChatCompletionsChat.
func (mr *MockClientMockRecorder) CallChatCompletionsChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallChatCompletionsChat", reflect.TypeOf((*MockClient)(nil).CallChatCompletionsChat), arg0, arg1)
}

// CallChatCompletionsReasoner mocks base method.
func (m *MockClient) CallChatCompletionsReasoner(arg0 context.Context, arg1 *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallChatCompletionsReasoner", arg0, arg1)
	ret0, _ := ret[0].(*response.ChatCompletionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallChatCompletionsReasoner indicates an expected call of CallChatCompletionsReasoner.
func (mr *MockClientMockRecorder) CallChatCompletionsReasoner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallChatCompletionsReasoner", reflect.TypeOf((*MockClient)(nil).CallChatCompletionsReasoner), arg0, arg1)
}

// PingChatCompletions mocks base method.
func (m *MockClient) PingChatCompletions(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingChatCompletions", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PingChatCompletions indicates an expected call of PingChatCompletions.
func (mr *MockClientMockRecorder) PingChatCompletions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingChatCompletions", reflect.TypeOf((*MockClient)(nil).PingChatCompletions), arg0, arg1)
}

// StreamChatCompletionsChat mocks base method.
func (m *MockClient) StreamChatCompletionsChat(arg0 context.Context, arg1 *request.ChatCompletionsRequest) (response.StreamReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamChatCompletionsChat", arg0, arg1)
	ret0, _ := ret[0].(response.StreamReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamChatCompletionsChat indicates an expected call of StreamChatCompletionsChat.
func (mr *MockClientMockRecorder) StreamChatCompletionsChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamChatCompletionsChat", reflect.TypeOf((*MockClient)(nil).StreamChatCompletionsChat), arg0, arg1)
}

// StreamChatCompletionsReasoner mocks base method.
func (m *MockClient) StreamChatCompletionsReasoner(arg0 context.Context, arg1 *request.ChatCompletionsRequest) (response.StreamReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamChatCompletionsReasoner", arg0, arg1)
	ret0, _ := ret[0].(response.StreamReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamChatCompletionsReasoner indicates an expected call of StreamChatCompletionsReasoner.
func (mr *MockClientMockRecorder) StreamChatCompletionsReasoner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamChatCompletionsReasoner", reflect.TypeOf((*MockClient)(nil).StreamChatCompletionsReasoner), arg0, arg1)
}
