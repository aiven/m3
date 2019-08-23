// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/query/block (interfaces: Block,StepIter,SeriesIter,Builder,Step,UnconsolidatedBlock,UnconsolidatedStepIter,UnconsolidatedSeriesIter,UnconsolidatedStep)

// Copyright (c) 2019 Uber Technologies, Inc.
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

// Package block is a generated GoMock package.
package block

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/query/ts"

	"github.com/golang/mock/gomock"
)

// MockBlock is a mock of Block interface
type MockBlock struct {
	ctrl     *gomock.Controller
	recorder *MockBlockMockRecorder
}

// MockBlockMockRecorder is the mock recorder for MockBlock
type MockBlockMockRecorder struct {
	mock *MockBlock
}

// NewMockBlock creates a new mock instance
func NewMockBlock(ctrl *gomock.Controller) *MockBlock {
	mock := &MockBlock{ctrl: ctrl}
	mock.recorder = &MockBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlock) EXPECT() *MockBlockMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockBlock) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockBlockMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBlock)(nil).Close))
}

// Info mocks base method
func (m *MockBlock) Info() BlockInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(BlockInfo)
	return ret0
}

// Info indicates an expected call of Info
func (mr *MockBlockMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBlock)(nil).Info))
}

// Meta mocks base method
func (m *MockBlock) Meta() Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(Metadata)
	return ret0
}

// Meta indicates an expected call of Meta
func (mr *MockBlockMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockBlock)(nil).Meta))
}

// SeriesIter mocks base method
func (m *MockBlock) SeriesIter() (SeriesIter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesIter")
	ret0, _ := ret[0].(SeriesIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeriesIter indicates an expected call of SeriesIter
func (mr *MockBlockMockRecorder) SeriesIter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesIter", reflect.TypeOf((*MockBlock)(nil).SeriesIter))
}

// StepIter mocks base method
func (m *MockBlock) StepIter() (StepIter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StepIter")
	ret0, _ := ret[0].(StepIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StepIter indicates an expected call of StepIter
func (mr *MockBlockMockRecorder) StepIter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StepIter", reflect.TypeOf((*MockBlock)(nil).StepIter))
}

// Unconsolidated mocks base method
func (m *MockBlock) Unconsolidated() (UnconsolidatedBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unconsolidated")
	ret0, _ := ret[0].(UnconsolidatedBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unconsolidated indicates an expected call of Unconsolidated
func (mr *MockBlockMockRecorder) Unconsolidated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unconsolidated", reflect.TypeOf((*MockBlock)(nil).Unconsolidated))
}

// WithMetadata mocks base method
func (m *MockBlock) WithMetadata(arg0 Metadata, arg1 []SeriesMeta) (Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithMetadata", arg0, arg1)
	ret0, _ := ret[0].(Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithMetadata indicates an expected call of WithMetadata
func (mr *MockBlockMockRecorder) WithMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithMetadata", reflect.TypeOf((*MockBlock)(nil).WithMetadata), arg0, arg1)
}

// MockStepIter is a mock of StepIter interface
type MockStepIter struct {
	ctrl     *gomock.Controller
	recorder *MockStepIterMockRecorder
}

// MockStepIterMockRecorder is the mock recorder for MockStepIter
type MockStepIterMockRecorder struct {
	mock *MockStepIter
}

// NewMockStepIter creates a new mock instance
func NewMockStepIter(ctrl *gomock.Controller) *MockStepIter {
	mock := &MockStepIter{ctrl: ctrl}
	mock.recorder = &MockStepIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStepIter) EXPECT() *MockStepIterMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStepIter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockStepIterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStepIter)(nil).Close))
}

// Current mocks base method
func (m *MockStepIter) Current() Step {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(Step)
	return ret0
}

// Current indicates an expected call of Current
func (mr *MockStepIterMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockStepIter)(nil).Current))
}

// Err mocks base method
func (m *MockStepIter) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockStepIterMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockStepIter)(nil).Err))
}

// Next mocks base method
func (m *MockStepIter) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockStepIterMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockStepIter)(nil).Next))
}

// SeriesMeta mocks base method
func (m *MockStepIter) SeriesMeta() []SeriesMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesMeta")
	ret0, _ := ret[0].([]SeriesMeta)
	return ret0
}

// SeriesMeta indicates an expected call of SeriesMeta
func (mr *MockStepIterMockRecorder) SeriesMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesMeta", reflect.TypeOf((*MockStepIter)(nil).SeriesMeta))
}

// StepCount mocks base method
func (m *MockStepIter) StepCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StepCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// StepCount indicates an expected call of StepCount
func (mr *MockStepIterMockRecorder) StepCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StepCount", reflect.TypeOf((*MockStepIter)(nil).StepCount))
}

// MockSeriesIter is a mock of SeriesIter interface
type MockSeriesIter struct {
	ctrl     *gomock.Controller
	recorder *MockSeriesIterMockRecorder
}

// MockSeriesIterMockRecorder is the mock recorder for MockSeriesIter
type MockSeriesIterMockRecorder struct {
	mock *MockSeriesIter
}

// NewMockSeriesIter creates a new mock instance
func NewMockSeriesIter(ctrl *gomock.Controller) *MockSeriesIter {
	mock := &MockSeriesIter{ctrl: ctrl}
	mock.recorder = &MockSeriesIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSeriesIter) EXPECT() *MockSeriesIterMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSeriesIter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSeriesIterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSeriesIter)(nil).Close))
}

// Current mocks base method
func (m *MockSeriesIter) Current() Series {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(Series)
	return ret0
}

// Current indicates an expected call of Current
func (mr *MockSeriesIterMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockSeriesIter)(nil).Current))
}

// Err mocks base method
func (m *MockSeriesIter) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockSeriesIterMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSeriesIter)(nil).Err))
}

// Next mocks base method
func (m *MockSeriesIter) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockSeriesIterMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockSeriesIter)(nil).Next))
}

// SeriesCount mocks base method
func (m *MockSeriesIter) SeriesCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// SeriesCount indicates an expected call of SeriesCount
func (mr *MockSeriesIterMockRecorder) SeriesCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesCount", reflect.TypeOf((*MockSeriesIter)(nil).SeriesCount))
}

// SeriesMeta mocks base method
func (m *MockSeriesIter) SeriesMeta() []SeriesMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesMeta")
	ret0, _ := ret[0].([]SeriesMeta)
	return ret0
}

// SeriesMeta indicates an expected call of SeriesMeta
func (mr *MockSeriesIterMockRecorder) SeriesMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesMeta", reflect.TypeOf((*MockSeriesIter)(nil).SeriesMeta))
}

// MockBuilder is a mock of Builder interface
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// AddCols mocks base method
func (m *MockBuilder) AddCols(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCols", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCols indicates an expected call of AddCols
func (mr *MockBuilderMockRecorder) AddCols(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCols", reflect.TypeOf((*MockBuilder)(nil).AddCols), arg0)
}

// AppendValue mocks base method
func (m *MockBuilder) AppendValue(arg0 int, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendValue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendValue indicates an expected call of AppendValue
func (mr *MockBuilderMockRecorder) AppendValue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendValue", reflect.TypeOf((*MockBuilder)(nil).AppendValue), arg0, arg1)
}

// AppendValues mocks base method
func (m *MockBuilder) AppendValues(arg0 int, arg1 []float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendValues", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendValues indicates an expected call of AppendValues
func (mr *MockBuilderMockRecorder) AppendValues(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendValues", reflect.TypeOf((*MockBuilder)(nil).AppendValues), arg0, arg1)
}

// Build mocks base method
func (m *MockBuilder) Build() Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(Block)
	return ret0
}

// Build indicates an expected call of Build
func (mr *MockBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuilder)(nil).Build))
}

// BuildAsType mocks base method
func (m *MockBuilder) BuildAsType(arg0 BlockType) Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildAsType", arg0)
	ret0, _ := ret[0].(Block)
	return ret0
}

// BuildAsType indicates an expected call of BuildAsType
func (mr *MockBuilderMockRecorder) BuildAsType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAsType", reflect.TypeOf((*MockBuilder)(nil).BuildAsType), arg0)
}

// MockStep is a mock of Step interface
type MockStep struct {
	ctrl     *gomock.Controller
	recorder *MockStepMockRecorder
}

// MockStepMockRecorder is the mock recorder for MockStep
type MockStepMockRecorder struct {
	mock *MockStep
}

// NewMockStep creates a new mock instance
func NewMockStep(ctrl *gomock.Controller) *MockStep {
	mock := &MockStep{ctrl: ctrl}
	mock.recorder = &MockStepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStep) EXPECT() *MockStepMockRecorder {
	return m.recorder
}

// Time mocks base method
func (m *MockStep) Time() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Time indicates an expected call of Time
func (mr *MockStepMockRecorder) Time() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockStep)(nil).Time))
}

// Values mocks base method
func (m *MockStep) Values() []float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].([]float64)
	return ret0
}

// Values indicates an expected call of Values
func (mr *MockStepMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockStep)(nil).Values))
}

// MockUnconsolidatedBlock is a mock of UnconsolidatedBlock interface
type MockUnconsolidatedBlock struct {
	ctrl     *gomock.Controller
	recorder *MockUnconsolidatedBlockMockRecorder
}

// MockUnconsolidatedBlockMockRecorder is the mock recorder for MockUnconsolidatedBlock
type MockUnconsolidatedBlockMockRecorder struct {
	mock *MockUnconsolidatedBlock
}

// NewMockUnconsolidatedBlock creates a new mock instance
func NewMockUnconsolidatedBlock(ctrl *gomock.Controller) *MockUnconsolidatedBlock {
	mock := &MockUnconsolidatedBlock{ctrl: ctrl}
	mock.recorder = &MockUnconsolidatedBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnconsolidatedBlock) EXPECT() *MockUnconsolidatedBlockMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockUnconsolidatedBlock) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockUnconsolidatedBlockMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUnconsolidatedBlock)(nil).Close))
}

// Consolidate mocks base method
func (m *MockUnconsolidatedBlock) Consolidate() (Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consolidate")
	ret0, _ := ret[0].(Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consolidate indicates an expected call of Consolidate
func (mr *MockUnconsolidatedBlockMockRecorder) Consolidate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consolidate", reflect.TypeOf((*MockUnconsolidatedBlock)(nil).Consolidate))
}

// Meta mocks base method
func (m *MockUnconsolidatedBlock) Meta() Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(Metadata)
	return ret0
}

// Meta indicates an expected call of Meta
func (mr *MockUnconsolidatedBlockMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockUnconsolidatedBlock)(nil).Meta))
}

// SeriesIter mocks base method
func (m *MockUnconsolidatedBlock) SeriesIter() (UnconsolidatedSeriesIter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesIter")
	ret0, _ := ret[0].(UnconsolidatedSeriesIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeriesIter indicates an expected call of SeriesIter
func (mr *MockUnconsolidatedBlockMockRecorder) SeriesIter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesIter", reflect.TypeOf((*MockUnconsolidatedBlock)(nil).SeriesIter))
}

// StepIter mocks base method
func (m *MockUnconsolidatedBlock) StepIter() (UnconsolidatedStepIter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StepIter")
	ret0, _ := ret[0].(UnconsolidatedStepIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StepIter indicates an expected call of StepIter
func (mr *MockUnconsolidatedBlockMockRecorder) StepIter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StepIter", reflect.TypeOf((*MockUnconsolidatedBlock)(nil).StepIter))
}

// WithMetadata mocks base method
func (m *MockUnconsolidatedBlock) WithMetadata(arg0 Metadata, arg1 []SeriesMeta) (UnconsolidatedBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithMetadata", arg0, arg1)
	ret0, _ := ret[0].(UnconsolidatedBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithMetadata indicates an expected call of WithMetadata
func (mr *MockUnconsolidatedBlockMockRecorder) WithMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithMetadata", reflect.TypeOf((*MockUnconsolidatedBlock)(nil).WithMetadata), arg0, arg1)
}

// MockUnconsolidatedStepIter is a mock of UnconsolidatedStepIter interface
type MockUnconsolidatedStepIter struct {
	ctrl     *gomock.Controller
	recorder *MockUnconsolidatedStepIterMockRecorder
}

// MockUnconsolidatedStepIterMockRecorder is the mock recorder for MockUnconsolidatedStepIter
type MockUnconsolidatedStepIterMockRecorder struct {
	mock *MockUnconsolidatedStepIter
}

// NewMockUnconsolidatedStepIter creates a new mock instance
func NewMockUnconsolidatedStepIter(ctrl *gomock.Controller) *MockUnconsolidatedStepIter {
	mock := &MockUnconsolidatedStepIter{ctrl: ctrl}
	mock.recorder = &MockUnconsolidatedStepIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnconsolidatedStepIter) EXPECT() *MockUnconsolidatedStepIterMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockUnconsolidatedStepIter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockUnconsolidatedStepIterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUnconsolidatedStepIter)(nil).Close))
}

// Current mocks base method
func (m *MockUnconsolidatedStepIter) Current() UnconsolidatedStep {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(UnconsolidatedStep)
	return ret0
}

// Current indicates an expected call of Current
func (mr *MockUnconsolidatedStepIterMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockUnconsolidatedStepIter)(nil).Current))
}

// Err mocks base method
func (m *MockUnconsolidatedStepIter) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockUnconsolidatedStepIterMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockUnconsolidatedStepIter)(nil).Err))
}

// Next mocks base method
func (m *MockUnconsolidatedStepIter) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockUnconsolidatedStepIterMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockUnconsolidatedStepIter)(nil).Next))
}

// SeriesMeta mocks base method
func (m *MockUnconsolidatedStepIter) SeriesMeta() []SeriesMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesMeta")
	ret0, _ := ret[0].([]SeriesMeta)
	return ret0
}

// SeriesMeta indicates an expected call of SeriesMeta
func (mr *MockUnconsolidatedStepIterMockRecorder) SeriesMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesMeta", reflect.TypeOf((*MockUnconsolidatedStepIter)(nil).SeriesMeta))
}

// StepCount mocks base method
func (m *MockUnconsolidatedStepIter) StepCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StepCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// StepCount indicates an expected call of StepCount
func (mr *MockUnconsolidatedStepIterMockRecorder) StepCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StepCount", reflect.TypeOf((*MockUnconsolidatedStepIter)(nil).StepCount))
}

// MockUnconsolidatedSeriesIter is a mock of UnconsolidatedSeriesIter interface
type MockUnconsolidatedSeriesIter struct {
	ctrl     *gomock.Controller
	recorder *MockUnconsolidatedSeriesIterMockRecorder
}

// MockUnconsolidatedSeriesIterMockRecorder is the mock recorder for MockUnconsolidatedSeriesIter
type MockUnconsolidatedSeriesIterMockRecorder struct {
	mock *MockUnconsolidatedSeriesIter
}

// NewMockUnconsolidatedSeriesIter creates a new mock instance
func NewMockUnconsolidatedSeriesIter(ctrl *gomock.Controller) *MockUnconsolidatedSeriesIter {
	mock := &MockUnconsolidatedSeriesIter{ctrl: ctrl}
	mock.recorder = &MockUnconsolidatedSeriesIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnconsolidatedSeriesIter) EXPECT() *MockUnconsolidatedSeriesIterMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockUnconsolidatedSeriesIter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockUnconsolidatedSeriesIterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUnconsolidatedSeriesIter)(nil).Close))
}

// Current mocks base method
func (m *MockUnconsolidatedSeriesIter) Current() UnconsolidatedSeries {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(UnconsolidatedSeries)
	return ret0
}

// Current indicates an expected call of Current
func (mr *MockUnconsolidatedSeriesIterMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockUnconsolidatedSeriesIter)(nil).Current))
}

// Err mocks base method
func (m *MockUnconsolidatedSeriesIter) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockUnconsolidatedSeriesIterMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockUnconsolidatedSeriesIter)(nil).Err))
}

// Next mocks base method
func (m *MockUnconsolidatedSeriesIter) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockUnconsolidatedSeriesIterMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockUnconsolidatedSeriesIter)(nil).Next))
}

// SeriesCount mocks base method
func (m *MockUnconsolidatedSeriesIter) SeriesCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// SeriesCount indicates an expected call of SeriesCount
func (mr *MockUnconsolidatedSeriesIterMockRecorder) SeriesCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesCount", reflect.TypeOf((*MockUnconsolidatedSeriesIter)(nil).SeriesCount))
}

// SeriesMeta mocks base method
func (m *MockUnconsolidatedSeriesIter) SeriesMeta() []SeriesMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesMeta")
	ret0, _ := ret[0].([]SeriesMeta)
	return ret0
}

// SeriesMeta indicates an expected call of SeriesMeta
func (mr *MockUnconsolidatedSeriesIterMockRecorder) SeriesMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesMeta", reflect.TypeOf((*MockUnconsolidatedSeriesIter)(nil).SeriesMeta))
}

// MockUnconsolidatedStep is a mock of UnconsolidatedStep interface
type MockUnconsolidatedStep struct {
	ctrl     *gomock.Controller
	recorder *MockUnconsolidatedStepMockRecorder
}

// MockUnconsolidatedStepMockRecorder is the mock recorder for MockUnconsolidatedStep
type MockUnconsolidatedStepMockRecorder struct {
	mock *MockUnconsolidatedStep
}

// NewMockUnconsolidatedStep creates a new mock instance
func NewMockUnconsolidatedStep(ctrl *gomock.Controller) *MockUnconsolidatedStep {
	mock := &MockUnconsolidatedStep{ctrl: ctrl}
	mock.recorder = &MockUnconsolidatedStepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnconsolidatedStep) EXPECT() *MockUnconsolidatedStepMockRecorder {
	return m.recorder
}

// Time mocks base method
func (m *MockUnconsolidatedStep) Time() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Time indicates an expected call of Time
func (mr *MockUnconsolidatedStepMockRecorder) Time() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockUnconsolidatedStep)(nil).Time))
}

// Values mocks base method
func (m *MockUnconsolidatedStep) Values() []ts.Datapoints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].([]ts.Datapoints)
	return ret0
}

// Values indicates an expected call of Values
func (mr *MockUnconsolidatedStepMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockUnconsolidatedStep)(nil).Values))
}
