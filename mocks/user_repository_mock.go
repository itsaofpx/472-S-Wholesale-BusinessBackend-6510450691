package mocks

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of repositories.UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *entities.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id int) (*entities.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) GetAllUsers() (*[]entities.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entities.User), args.Error(1)
}

func (m *MockUserRepository) FindUserByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserByID(id int, user *request.UpdateUserByIDRequest) (*entities.User, error) {
	args := m.Called(id, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserTierByID(req *request.UpdateTierByUserIDRequest, user *entities.User) (*entities.User, error) {
	args := m.Called(req, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) ChangePassword(req *request.ChangePasswordRequest) error {
	args := m.Called(req)
	return args.Error(0)
}
