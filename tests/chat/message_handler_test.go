package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/adapters/api"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMessageUseCase struct {
	mock.Mock
}

func (m *MockMessageUseCase) CreateMessage(msg entities.Message) (entities.Message, error) {
	args := m.Called(msg)
	return args.Get(0).(entities.Message), args.Error(1)
}

func (m *MockMessageUseCase) CreateMessageByChatID(msg entities.Message) (entities.Message, error) {
	args := m.Called(msg)
	return args.Get(0).(entities.Message), args.Error(1)
}

func TestMessageHandler_CreateMessage(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockMessageUseCase)
	handler := api.InitiateMessageHandler(mockUseCase)

	app.Post("/chat/:id/message", handler.CreateMessage)

	tests := []struct {
		name         string
		chatID       string
		message      entities.Message
		expectedCode int
		expectError  bool
	}{
		{
			name:   "Valid Message Creation",
			chatID: "1",
			message: entities.Message{
				UserID:    1,
				ChatID:    1,
				Body:      "Test message",
				Timestamp: time.Now(),
			},
			expectedCode: fiber.StatusCreated,
			expectError:  false,
		},
		{
			name:   "Invalid Message - No UserID",
			chatID: "1",
			message: entities.Message{
				Body: "Test message",
			},
			expectedCode: fiber.StatusBadRequest,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectError {
				mockUseCase.On("CreateMessage", mock.AnythingOfType("entities.Message")).Return(tt.message, nil).Once()
			}

			jsonData, _ := json.Marshal(tt.message)
			req := httptest.NewRequest(http.MethodPost, "/chat/"+tt.chatID+"/message", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			if !tt.expectError {
				var responseMessage entities.Message
				json.NewDecoder(resp.Body).Decode(&responseMessage)
				assert.Equal(t, tt.message.Body, responseMessage.Body)
				assert.Equal(t, tt.message.UserID, responseMessage.UserID)
				assert.Equal(t, tt.message.ChatID, responseMessage.ChatID)
			}
		})
	}
}

func TestMessageHandler_CreateMessageByChatID(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockMessageUseCase)
	handler := api.InitiateMessageHandler(mockUseCase)

	app.Post("/chat/:id/message/create", handler.CreateMessageByChatID)

	tests := []struct {
		name         string
		chatID       string
		message      entities.Message
		expectedCode int
		expectError  bool
	}{
		{
			name:   "Valid Message Creation",
			chatID: "1",
			message: entities.Message{
				UserID:    1,
				ChatID:    1,
				Body:      "Test message",
				Timestamp: time.Now(),
			},
			expectedCode: fiber.StatusCreated,
			expectError:  false,
		},
		{
			name:   "Invalid Message",
			chatID: "1",
			message: entities.Message{
				Body: "Test message",
			},
			expectedCode: fiber.StatusBadRequest,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectError {
				mockUseCase.On("CreateMessageByChatID", mock.AnythingOfType("entities.Message")).
					Return(tt.message, nil).Once()
			}

			jsonData, _ := json.Marshal(tt.message)
			req := httptest.NewRequest(http.MethodPost, "/chat/"+tt.chatID+"/message/create", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			if !tt.expectError {
				var responseMessage entities.Message
				json.NewDecoder(resp.Body).Decode(&responseMessage)
				assert.Equal(t, tt.message.Body, responseMessage.Body)
				assert.Equal(t, tt.message.UserID, responseMessage.UserID)
				assert.Equal(t, tt.message.ChatID, responseMessage.ChatID)
			}
		})
	}
}
