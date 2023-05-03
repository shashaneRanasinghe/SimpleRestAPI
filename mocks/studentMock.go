// Code generated by MockGen. DO NOT EDIT.
// Source: ../src/github.com/shashaneRanasinghe/simpleAPI/internal/usecases/student/studentInterface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/shashaneRanasinghe/simpleAPI/internal/models"
)

// MockStudentUsecase is a mock of StudentUsecase interface.
type MockStudentUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockStudentUsecaseMockRecorder
}

// MockStudentUsecaseMockRecorder is the mock recorder for MockStudentUsecase.
type MockStudentUsecaseMockRecorder struct {
	mock *MockStudentUsecase
}

// NewMockStudentUsecase creates a new mock instance.
func NewMockStudentUsecase(ctrl *gomock.Controller) *MockStudentUsecase {
	mock := &MockStudentUsecase{ctrl: ctrl}
	mock.recorder = &MockStudentUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentUsecase) EXPECT() *MockStudentUsecaseMockRecorder {
	return m.recorder
}

// CreateStudent mocks base method.
func (m *MockStudentUsecase) CreateStudent(student *models.Student) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudent", student)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudent indicates an expected call of CreateStudent.
func (mr *MockStudentUsecaseMockRecorder) CreateStudent(student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudent", reflect.TypeOf((*MockStudentUsecase)(nil).CreateStudent), student)
}

// DeleteStudent mocks base method.
func (m *MockStudentUsecase) DeleteStudent(id int) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudent", id)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockStudentUsecaseMockRecorder) DeleteStudent(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockStudentUsecase)(nil).DeleteStudent), id)
}

// GetAllStudents mocks base method.
func (m *MockStudentUsecase) GetAllStudents() ([]models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllStudents")
	ret0, _ := ret[0].([]models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllStudents indicates an expected call of GetAllStudents.
func (mr *MockStudentUsecaseMockRecorder) GetAllStudents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllStudents", reflect.TypeOf((*MockStudentUsecase)(nil).GetAllStudents))
}

// GetStudent mocks base method.
func (m *MockStudentUsecase) GetStudent(id int) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudent", id)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudent indicates an expected call of GetStudent.
func (mr *MockStudentUsecaseMockRecorder) GetStudent(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudent", reflect.TypeOf((*MockStudentUsecase)(nil).GetStudent), id)
}

// SearchStudent mocks base method.
func (m *MockStudentUsecase) SearchStudent(searchString string, pagination models.Pagination, sortBy models.SortBy) (*models.StudentSearchData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchStudent", searchString, pagination, sortBy)
	ret0, _ := ret[0].(*models.StudentSearchData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchStudent indicates an expected call of SearchStudent.
func (mr *MockStudentUsecaseMockRecorder) SearchStudent(searchString, pagination, sortBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchStudent", reflect.TypeOf((*MockStudentUsecase)(nil).SearchStudent), searchString, pagination, sortBy)
}

// UpdateStudent mocks base method.
func (m *MockStudentUsecase) UpdateStudent(student *models.Student) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudent", student)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStudent indicates an expected call of UpdateStudent.
func (mr *MockStudentUsecaseMockRecorder) UpdateStudent(student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudent", reflect.TypeOf((*MockStudentUsecase)(nil).UpdateStudent), student)
}