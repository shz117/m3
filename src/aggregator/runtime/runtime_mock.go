// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3aggregator/runtime (interfaces: OptionsWatcher)

package runtime

import (
	"github.com/golang/mock/gomock"
)

// Mock of OptionsWatcher interface
type MockOptionsWatcher struct {
	ctrl     *gomock.Controller
	recorder *_MockOptionsWatcherRecorder
}

// Recorder for MockOptionsWatcher (not exported)
type _MockOptionsWatcherRecorder struct {
	mock *MockOptionsWatcher
}

func NewMockOptionsWatcher(ctrl *gomock.Controller) *MockOptionsWatcher {
	mock := &MockOptionsWatcher{ctrl: ctrl}
	mock.recorder = &_MockOptionsWatcherRecorder{mock}
	return mock
}

func (_m *MockOptionsWatcher) EXPECT() *_MockOptionsWatcherRecorder {
	return _m.recorder
}

func (_m *MockOptionsWatcher) SetRuntimeOptions(_param0 Options) {
	_m.ctrl.Call(_m, "SetRuntimeOptions", _param0)
}

func (_mr *_MockOptionsWatcherRecorder) SetRuntimeOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRuntimeOptions", arg0)
}