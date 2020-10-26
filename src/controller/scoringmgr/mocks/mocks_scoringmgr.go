// Code generated by MockGen. DO NOT EDIT.
// Source: controller/scoringmgr (interfaces: Scoring)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockScoring is a mock of Scoring interface.
type MockScoring struct {
	ctrl     *gomock.Controller
	recorder *MockScoringMockRecorder
}

// MockScoringMockRecorder is the mock recorder for MockScoring.
type MockScoringMockRecorder struct {
	mock *MockScoring
}

// NewMockScoring creates a new mock instance.
func NewMockScoring(ctrl *gomock.Controller) *MockScoring {
	mock := &MockScoring{ctrl: ctrl}
	mock.recorder = &MockScoringMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScoring) EXPECT() *MockScoringMockRecorder {
	return m.recorder
}

// GetResource mocks base method.
func (m *MockScoring) GetResource(arg0 string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockScoringMockRecorder) GetResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockScoring)(nil).GetResource), arg0)
}

// GetScore mocks base method.
func (m *MockScoring) GetScore(arg0 string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScore", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScore indicates an expected call of GetScore.
func (mr *MockScoringMockRecorder) GetScore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScore", reflect.TypeOf((*MockScoring)(nil).GetScore), arg0)
}

// GetScoreWithResource mocks base method.
func (m *MockScoring) GetScoreWithResource(arg0 map[string]interface{}) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScoreWithResource", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScoreWithResource indicates an expected call of GetScoreWithResource.
func (mr *MockScoringMockRecorder) GetScoreWithResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScoreWithResource", reflect.TypeOf((*MockScoring)(nil).GetScoreWithResource), arg0)
}
