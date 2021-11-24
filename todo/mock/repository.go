// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	todo "github.com/detohm/todo-app-study/todo"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CompleteToDo mocks base method.
func (m *MockRepository) CompleteToDo(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteToDo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteToDo indicates an expected call of CompleteToDo.
func (mr *MockRepositoryMockRecorder) CompleteToDo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteToDo", reflect.TypeOf((*MockRepository)(nil).CompleteToDo), arg0)
}

// CreateToDo mocks base method.
func (m *MockRepository) CreateToDo(arg0 todo.ToDo) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToDo", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToDo indicates an expected call of CreateToDo.
func (mr *MockRepositoryMockRecorder) CreateToDo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToDo", reflect.TypeOf((*MockRepository)(nil).CreateToDo), arg0)
}

// DeleteToDo mocks base method.
func (m *MockRepository) DeleteToDo(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteToDo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteToDo indicates an expected call of DeleteToDo.
func (mr *MockRepositoryMockRecorder) DeleteToDo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteToDo", reflect.TypeOf((*MockRepository)(nil).DeleteToDo), arg0)
}

// GetToDoList mocks base method.
func (m *MockRepository) GetToDoList() ([]todo.ToDo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToDoList")
	ret0, _ := ret[0].([]todo.ToDo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToDoList indicates an expected call of GetToDoList.
func (mr *MockRepositoryMockRecorder) GetToDoList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToDoList", reflect.TypeOf((*MockRepository)(nil).GetToDoList))
}
