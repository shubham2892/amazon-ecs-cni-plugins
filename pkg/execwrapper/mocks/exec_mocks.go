// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-cni-plugins/pkg/execwrapper (interfaces: Cmd,Exec)

package mock_execwrapper

import (
	io "io"

	execwrapper "github.com/aws/amazon-ecs-cni-plugins/pkg/execwrapper"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Cmd interface
type MockCmd struct {
	ctrl     *gomock.Controller
	recorder *_MockCmdRecorder
}

// Recorder for MockCmd (not exported)
type _MockCmdRecorder struct {
	mock *MockCmd
}

func NewMockCmd(ctrl *gomock.Controller) *MockCmd {
	mock := &MockCmd{ctrl: ctrl}
	mock.recorder = &_MockCmdRecorder{mock}
	return mock
}

func (_m *MockCmd) EXPECT() *_MockCmdRecorder {
	return _m.recorder
}

func (_m *MockCmd) CombinedOutput() ([]byte, error) {
	ret := _m.ctrl.Call(_m, "CombinedOutput")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCmdRecorder) CombinedOutput() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CombinedOutput")
}

func (_m *MockCmd) Output() ([]byte, error) {
	ret := _m.ctrl.Call(_m, "Output")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCmdRecorder) Output() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Output")
}

func (_m *MockCmd) Run() error {
	ret := _m.ctrl.Call(_m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCmdRecorder) Run() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run")
}

func (_m *MockCmd) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCmdRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockCmd) StderrPipe() (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "StderrPipe")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCmdRecorder) StderrPipe() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StderrPipe")
}

func (_m *MockCmd) StdinPipe() (io.WriteCloser, error) {
	ret := _m.ctrl.Call(_m, "StdinPipe")
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCmdRecorder) StdinPipe() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StdinPipe")
}

func (_m *MockCmd) StdoutPipe() (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "StdoutPipe")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCmdRecorder) StdoutPipe() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StdoutPipe")
}

func (_m *MockCmd) Wait() error {
	ret := _m.ctrl.Call(_m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCmdRecorder) Wait() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Wait")
}

// Mock of Exec interface
type MockExec struct {
	ctrl     *gomock.Controller
	recorder *_MockExecRecorder
}

// Recorder for MockExec (not exported)
type _MockExecRecorder struct {
	mock *MockExec
}

func NewMockExec(ctrl *gomock.Controller) *MockExec {
	mock := &MockExec{ctrl: ctrl}
	mock.recorder = &_MockExecRecorder{mock}
	return mock
}

func (_m *MockExec) EXPECT() *_MockExecRecorder {
	return _m.recorder
}

func (_m *MockExec) Command(_param0 string, _param1 ...string) execwrapper.Cmd {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Command", _s...)
	ret0, _ := ret[0].(execwrapper.Cmd)
	return ret0
}

func (_mr *_MockExecRecorder) Command(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Command", _s...)
}

func (_m *MockExec) LookPath(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "LookPath", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockExecRecorder) LookPath(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LookPath", arg0)
}
