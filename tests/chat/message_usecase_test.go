package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) CreateMessage(msg entities.Message) (entities.Message, error) {
	args := m.Called(msg)
	return args.Get(0).(entities.Message), args.Error(1)
}

func (m *MockMessageRepository) CreateMessageByChatID(msg entities.Message) (entities.Message, error) {
	args := m.Called(msg)
	return args.Get(0).(entities.Message), args.Error(1)
}

func (m *MockMessageRepository) ValidateReferences(userID, chatID int) error {
	args := m.Called(userID, chatID)
	return args.Error(0)
}

func TestMessageService_CreateMessage(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	service := usecases.InitiateMessageService(mockRepo)

	message := entities.Message{
		UserID:    1,
		ChatID:    1,
		Body:      "Test message",
		Timestamp: time.Now(),
	}

	tests := []struct {
		name        string
		input       entities.Message
		mockReturn  entities.Message
		expectError bool
	}{
		{
			name:        "Success",
			input:       message,
			mockReturn:  message,
			expectError: false,
		},
		{
			name:        "Validation Error",
			input:       message,
			mockReturn:  entities.Message{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectError {
				mockRepo.On("ValidateReferences", tt.input.UserID, tt.input.ChatID).
					Return(errors.New("validation error")).Once()
			} else {
				mockRepo.On("ValidateReferences", tt.input.UserID, tt.input.ChatID).
					Return(nil).Once()
				mockRepo.On("CreateMessage", tt.input).Return(tt.mockReturn, nil).Once()
			}

			result, err := service.CreateMessage(tt.input)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockReturn, result)
			}
		})
	}
}

func TestMessageService_CreateMessageByChatID(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	service := usecases.InitiateMessageService(mockRepo)

	message := entities.Message{
		UserID:    1,
		ChatID:    1,
		Body:      "Test message",
		Timestamp: time.Now(),
	}

	tests := []struct {
		name        string
		input       entities.Message
		mockReturn  entities.Message
		expectError bool
	}{
		{
			name:        "Success",
			input:       message,
			mockReturn:  message,
			expectError: false,
		},
		{
			name:        "Validation Error",
			input:       message,
			mockReturn:  entities.Message{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectError {
				// Mock ValidateReferences for success case
				mockRepo.On("ValidateReferences", tt.input.UserID, tt.input.ChatID).
					Return(nil).Once()

				// Mock CreateMessage for success case
				mockRepo.On("CreateMessage", mock.AnythingOfType("entities.Message")).
					Return(tt.mockReturn, nil).Once()
			} else {
				// Mock ValidateReferences for error case
				mockRepo.On("ValidateReferences", tt.input.UserID, tt.input.ChatID).
					Return(errors.New("validation error")).Once()
			}

			result, err := service.CreateMessageByChatID(tt.input)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockReturn, result)
			}

			// Verify all mocked calls were made
			mockRepo.AssertExpectations(t)
		})
	}
}
