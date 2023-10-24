// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -source interfaces.go -destination=../mocks/mock_tfc_trigger.go -package=mocks github.com/zapier/tfbuddy/pkg/tfc_trigger
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfc_trigger "github.com/zapier/tfbuddy/pkg/tfc_trigger"
	gomock "go.uber.org/mock/gomock"
)

// MockTrigger is a mock of Trigger interface.
type MockTrigger struct {
	ctrl     *gomock.Controller
	recorder *MockTriggerMockRecorder
}

// MockTriggerMockRecorder is the mock recorder for MockTrigger.
type MockTriggerMockRecorder struct {
	mock *MockTrigger
}

// NewMockTrigger creates a new mock instance.
func NewMockTrigger(ctrl *gomock.Controller) *MockTrigger {
	mock := &MockTrigger{ctrl: ctrl}
	mock.recorder = &MockTriggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrigger) EXPECT() *MockTriggerMockRecorder {
	return m.recorder
}

// GetAction mocks base method.
func (m *MockTrigger) GetAction() tfc_trigger.TriggerAction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAction")
	ret0, _ := ret[0].(tfc_trigger.TriggerAction)
	return ret0
}

// GetAction indicates an expected call of GetAction.
func (mr *MockTriggerMockRecorder) GetAction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAction", reflect.TypeOf((*MockTrigger)(nil).GetAction))
}

// GetBranch mocks base method.
func (m *MockTrigger) GetBranch() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranch")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockTriggerMockRecorder) GetBranch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockTrigger)(nil).GetBranch))
}

// GetCommitSHA mocks base method.
func (m *MockTrigger) GetCommitSHA() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitSHA")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCommitSHA indicates an expected call of GetCommitSHA.
func (mr *MockTriggerMockRecorder) GetCommitSHA() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitSHA", reflect.TypeOf((*MockTrigger)(nil).GetCommitSHA))
}

// GetMergeRequestDiscussionID mocks base method.
func (m *MockTrigger) GetMergeRequestDiscussionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMergeRequestDiscussionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMergeRequestDiscussionID indicates an expected call of GetMergeRequestDiscussionID.
func (mr *MockTriggerMockRecorder) GetMergeRequestDiscussionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestDiscussionID", reflect.TypeOf((*MockTrigger)(nil).GetMergeRequestDiscussionID))
}

// GetMergeRequestIID mocks base method.
func (m *MockTrigger) GetMergeRequestIID() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMergeRequestIID")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMergeRequestIID indicates an expected call of GetMergeRequestIID.
func (mr *MockTriggerMockRecorder) GetMergeRequestIID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestIID", reflect.TypeOf((*MockTrigger)(nil).GetMergeRequestIID))
}

// GetMergeRequestRootNoteID mocks base method.
func (m *MockTrigger) GetMergeRequestRootNoteID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMergeRequestRootNoteID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetMergeRequestRootNoteID indicates an expected call of GetMergeRequestRootNoteID.
func (mr *MockTriggerMockRecorder) GetMergeRequestRootNoteID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestRootNoteID", reflect.TypeOf((*MockTrigger)(nil).GetMergeRequestRootNoteID))
}

// GetProjectNameWithNamespace mocks base method.
func (m *MockTrigger) GetProjectNameWithNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectNameWithNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetProjectNameWithNamespace indicates an expected call of GetProjectNameWithNamespace.
func (mr *MockTriggerMockRecorder) GetProjectNameWithNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectNameWithNamespace", reflect.TypeOf((*MockTrigger)(nil).GetProjectNameWithNamespace))
}

// GetTriggerSource mocks base method.
func (m *MockTrigger) GetTriggerSource() tfc_trigger.TriggerSource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTriggerSource")
	ret0, _ := ret[0].(tfc_trigger.TriggerSource)
	return ret0
}

// GetTriggerSource indicates an expected call of GetTriggerSource.
func (mr *MockTriggerMockRecorder) GetTriggerSource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTriggerSource", reflect.TypeOf((*MockTrigger)(nil).GetTriggerSource))
}

// GetVcsProvider mocks base method.
func (m *MockTrigger) GetVcsProvider() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVcsProvider")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVcsProvider indicates an expected call of GetVcsProvider.
func (mr *MockTriggerMockRecorder) GetVcsProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVcsProvider", reflect.TypeOf((*MockTrigger)(nil).GetVcsProvider))
}

// GetWorkspace mocks base method.
func (m *MockTrigger) GetWorkspace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockTriggerMockRecorder) GetWorkspace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockTrigger)(nil).GetWorkspace))
}

// SetMergeRequestDiscussionID mocks base method.
func (m *MockTrigger) SetMergeRequestDiscussionID(mrdisID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMergeRequestDiscussionID", mrdisID)
}

// SetMergeRequestDiscussionID indicates an expected call of SetMergeRequestDiscussionID.
func (mr *MockTriggerMockRecorder) SetMergeRequestDiscussionID(mrdisID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMergeRequestDiscussionID", reflect.TypeOf((*MockTrigger)(nil).SetMergeRequestDiscussionID), mrdisID)
}

// SetMergeRequestRootNoteID mocks base method.
func (m *MockTrigger) SetMergeRequestRootNoteID(id int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMergeRequestRootNoteID", id)
}

// SetMergeRequestRootNoteID indicates an expected call of SetMergeRequestRootNoteID.
func (mr *MockTriggerMockRecorder) SetMergeRequestRootNoteID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMergeRequestRootNoteID", reflect.TypeOf((*MockTrigger)(nil).SetMergeRequestRootNoteID), id)
}

// TriggerCleanupEvent mocks base method.
func (m *MockTrigger) TriggerCleanupEvent(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerCleanupEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TriggerCleanupEvent indicates an expected call of TriggerCleanupEvent.
func (mr *MockTriggerMockRecorder) TriggerCleanupEvent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerCleanupEvent", reflect.TypeOf((*MockTrigger)(nil).TriggerCleanupEvent), arg0)
}

// TriggerTFCEvents mocks base method.
func (m *MockTrigger) TriggerTFCEvents(arg0 context.Context) (*tfc_trigger.TriggeredTFCWorkspaces, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerTFCEvents", arg0)
	ret0, _ := ret[0].(*tfc_trigger.TriggeredTFCWorkspaces)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerTFCEvents indicates an expected call of TriggerTFCEvents.
func (mr *MockTriggerMockRecorder) TriggerTFCEvents(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerTFCEvents", reflect.TypeOf((*MockTrigger)(nil).TriggerTFCEvents), arg0)
}
