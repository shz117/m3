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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db/storage/bootstrap/types.go

package bootstrap

import (
	"reflect"

	"github.com/m3db/m3db/storage/bootstrap/result"
	"github.com/m3db/m3db/storage/namespace"

	"github.com/golang/mock/gomock"
)

// MockProcess is a mock of Process interface
type MockProcess struct {
	ctrl     *gomock.Controller
	recorder *MockProcessMockRecorder
}

// MockProcessMockRecorder is the mock recorder for MockProcess
type MockProcessMockRecorder struct {
	mock *MockProcess
}

// NewMockProcess creates a new mock instance
func NewMockProcess(ctrl *gomock.Controller) *MockProcess {
	mock := &MockProcess{ctrl: ctrl}
	mock.recorder = &MockProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockProcess) EXPECT() *MockProcessMockRecorder {
	return _m.recorder
}

// SetBootstrapper mocks base method
func (_m *MockProcess) SetBootstrapper(bootstrapper Bootstrapper) {
	_m.ctrl.Call(_m, "SetBootstrapper", bootstrapper)
}

// SetBootstrapper indicates an expected call of SetBootstrapper
func (_mr *MockProcessMockRecorder) SetBootstrapper(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBootstrapper", reflect.TypeOf((*MockProcess)(nil).SetBootstrapper), arg0)
}

// Bootstrapper mocks base method
func (_m *MockProcess) Bootstrapper() Bootstrapper {
	ret := _m.ctrl.Call(_m, "Bootstrapper")
	ret0, _ := ret[0].(Bootstrapper)
	return ret0
}

// Bootstrapper indicates an expected call of Bootstrapper
func (_mr *MockProcessMockRecorder) Bootstrapper() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Bootstrapper", reflect.TypeOf((*MockProcess)(nil).Bootstrapper))
}

// Run mocks base method
func (_m *MockProcess) Run(ns namespace.Metadata, shards []uint32, targetRanges []TargetRange) (result.BootstrapResult, error) {
	ret := _m.ctrl.Call(_m, "Run", ns, shards, targetRanges)
	ret0, _ := ret[0].(result.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (_mr *MockProcessMockRecorder) Run(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Run", reflect.TypeOf((*MockProcess)(nil).Run), arg0, arg1, arg2)
}

// MockRunOptions is a mock of RunOptions interface
type MockRunOptions struct {
	ctrl     *gomock.Controller
	recorder *MockRunOptionsMockRecorder
}

// MockRunOptionsMockRecorder is the mock recorder for MockRunOptions
type MockRunOptionsMockRecorder struct {
	mock *MockRunOptions
}

// NewMockRunOptions creates a new mock instance
func NewMockRunOptions(ctrl *gomock.Controller) *MockRunOptions {
	mock := &MockRunOptions{ctrl: ctrl}
	mock.recorder = &MockRunOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockRunOptions) EXPECT() *MockRunOptionsMockRecorder {
	return _m.recorder
}

// SetIncremental mocks base method
func (_m *MockRunOptions) SetIncremental(value bool) RunOptions {
	ret := _m.ctrl.Call(_m, "SetIncremental", value)
	ret0, _ := ret[0].(RunOptions)
	return ret0
}

// SetIncremental indicates an expected call of SetIncremental
func (_mr *MockRunOptionsMockRecorder) SetIncremental(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetIncremental", reflect.TypeOf((*MockRunOptions)(nil).SetIncremental), arg0)
}

// Incremental mocks base method
func (_m *MockRunOptions) Incremental() bool {
	ret := _m.ctrl.Call(_m, "Incremental")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Incremental indicates an expected call of Incremental
func (_mr *MockRunOptionsMockRecorder) Incremental() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Incremental", reflect.TypeOf((*MockRunOptions)(nil).Incremental))
}

// MockBootstrapper is a mock of Bootstrapper interface
type MockBootstrapper struct {
	ctrl     *gomock.Controller
	recorder *MockBootstrapperMockRecorder
}

// MockBootstrapperMockRecorder is the mock recorder for MockBootstrapper
type MockBootstrapperMockRecorder struct {
	mock *MockBootstrapper
}

// NewMockBootstrapper creates a new mock instance
func NewMockBootstrapper(ctrl *gomock.Controller) *MockBootstrapper {
	mock := &MockBootstrapper{ctrl: ctrl}
	mock.recorder = &MockBootstrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockBootstrapper) EXPECT() *MockBootstrapperMockRecorder {
	return _m.recorder
}

// String mocks base method
func (_m *MockBootstrapper) String() string {
	ret := _m.ctrl.Call(_m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (_mr *MockBootstrapperMockRecorder) String() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "String", reflect.TypeOf((*MockBootstrapper)(nil).String))
}

// Can mocks base method
func (_m *MockBootstrapper) Can(strategy Strategy) bool {
	ret := _m.ctrl.Call(_m, "Can", strategy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Can indicates an expected call of Can
func (_mr *MockBootstrapperMockRecorder) Can(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Can", reflect.TypeOf((*MockBootstrapper)(nil).Can), arg0)
}

// Bootstrap mocks base method
func (_m *MockBootstrapper) Bootstrap(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, opts RunOptions) (result.BootstrapResult, error) {
	ret := _m.ctrl.Call(_m, "Bootstrap", ns, shardsTimeRanges, opts)
	ret0, _ := ret[0].(result.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bootstrap indicates an expected call of Bootstrap
func (_mr *MockBootstrapperMockRecorder) Bootstrap(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Bootstrap", reflect.TypeOf((*MockBootstrapper)(nil).Bootstrap), arg0, arg1, arg2)
}

// MockSource is a mock of Source interface
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMockRecorder
}

// MockSourceMockRecorder is the mock recorder for MockSource
type MockSourceMockRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockSource) EXPECT() *MockSourceMockRecorder {
	return _m.recorder
}

// Can mocks base method
func (_m *MockSource) Can(strategy Strategy) bool {
	ret := _m.ctrl.Call(_m, "Can", strategy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Can indicates an expected call of Can
func (_mr *MockSourceMockRecorder) Can(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Can", reflect.TypeOf((*MockSource)(nil).Can), arg0)
}

// Available mocks base method
func (_m *MockSource) Available(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges) result.ShardTimeRanges {
	ret := _m.ctrl.Call(_m, "Available", ns, shardsTimeRanges)
	ret0, _ := ret[0].(result.ShardTimeRanges)
	return ret0
}

// Available indicates an expected call of Available
func (_mr *MockSourceMockRecorder) Available(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Available", reflect.TypeOf((*MockSource)(nil).Available), arg0, arg1)
}

// Read mocks base method
func (_m *MockSource) Read(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, opts RunOptions) (result.BootstrapResult, error) {
	ret := _m.ctrl.Call(_m, "Read", ns, shardsTimeRanges, opts)
	ret0, _ := ret[0].(result.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (_mr *MockSourceMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Read", reflect.TypeOf((*MockSource)(nil).Read), arg0, arg1, arg2)
}