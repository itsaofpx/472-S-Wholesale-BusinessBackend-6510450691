package tests

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

// Mock Repository Definitions
type MockChatRepository struct {
	mock.Mock
}

func (m *MockChatRepository) CreateChat(c entities.Chat) (entities.Chat, error) {
	args := m.Called(c)
	return args.Get(0).(entities.Chat), args.Error(1)
}

func (m *MockChatRepository) GetAllChats() ([]entities.Chat, error) {
	args := m.Called()
	return args.Get(0).([]entities.Chat), args.Error(1)
}

func (m *MockChatRepository) GetChatByUserID(id string) (entities.Chat, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Chat), args.Error(1)
}

// Chat Usecase Tests
func TestChatService_CreateChat(t *testing.T) {
	mockRepo := new(MockChatRepository)
	service := usecases.InitiateChatService(mockRepo)

	tests := []struct {
		name        string
		input       entities.Chat
		mockReturn  entities.Chat
		expectError bool
	}{
		{
			name: "Success",
			input: entities.Chat{
				UserID: 1,
			},
			mockReturn: entities.Chat{
				Model:  gorm.Model{ID: 1},
				UserID: 1,
			},
			expectError: false,
		},
		{
			name: "Repository Error",
			input: entities.Chat{
				UserID: 2,
			},
			mockReturn:  entities.Chat{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectError {
				mockRepo.On("CreateChat", tt.input).Return(entities.Chat{}, errors.New("db error")).Once()
			} else {
				mockRepo.On("CreateChat", tt.input).Return(tt.mockReturn, nil).Once()
			}

			result, err := service.CreateChat(tt.input)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockReturn, result)
			}
		})
	}
}

func TestChatService_GetAllChats(t *testing.T) {
	mockRepo := new(MockChatRepository)
	service := usecases.InitiateChatService(mockRepo)

	expectedChats := []entities.Chat{
		{
			Model:  gorm.Model{ID: 1},
			UserID: 1,
		},
		{
			Model:  gorm.Model{ID: 2},
			UserID: 2,
		},
	}

	mockRepo.On("GetAllChats").Return(expectedChats, nil)

	result, err := service.GetAllChats()

	assert.NoError(t, err)
	assert.Equal(t, expectedChats, result)
}

func TestChatService_GetChatByUserID(t *testing.T) {
	mockRepo := new(MockChatRepository)
	service := usecases.InitiateChatService(mockRepo)

	tests := []struct {
		name        string
		userID      string
		mockReturn  entities.Chat
		expectError bool
	}{
		{
			name:   "Success",
			userID: "1",
			mockReturn: entities.Chat{
				Model:  gorm.Model{ID: 1},
				UserID: 1,
			},
			expectError: false,
		},
		{
			name:        "Not Found",
			userID:      "999",
			mockReturn:  entities.Chat{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectError {
				mockRepo.On("GetChatByUserID", tt.userID).Return(entities.Chat{}, errors.New("not found")).Once()
			} else {
				mockRepo.On("GetChatByUserID", tt.userID).Return(tt.mockReturn, nil).Once()
			}

			result, err := service.GetChatByUserID(tt.userID)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockReturn, result)
			}
		})
	}

}
