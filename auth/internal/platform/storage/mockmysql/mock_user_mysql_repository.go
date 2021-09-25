// Code generated by MockGen. DO NOT EDIT.
// Source: internal/user_repository.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	reflect "reflect"

	internal "github.com/apascualco/apascualco-auth/internal"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockUserRepository) Save(ctx context.Context, user internal.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockUserRepositoryMockRecorder) Save(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRepository)(nil).Save), ctx, user)
}

// SearchUserByEmail mocks base method.
func (m *MockUserRepository) SearchUserByEmail(ctx context.Context, email string) (internal.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUserByEmail", ctx, email)
	ret0, _ := ret[0].(internal.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUserByEmail indicates an expected call of SearchUserByEmail.
func (mr *MockUserRepositoryMockRecorder) SearchUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).SearchUserByEmail), ctx, email)
}