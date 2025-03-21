package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/ppwlsw/sa-project-backend/adapters/api"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
)

// Mock Definitions
type MockChatUseCase struct {
	mock.Mock
}

func (m *MockChatUseCase) CreateChat(c entities.Chat) (entities.Chat, error) {
	args := m.Called(c)
	return args.Get(0).(entities.Chat), args.Error(1)
}

func (m *MockChatUseCase) GetAllChats() ([]entities.Chat, error) {
	args := m.Called()
	return args.Get(0).([]entities.Chat), args.Error(1)
}

func (m *MockChatUseCase) GetChatByUserID(id string) (entities.Chat, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Chat), args.Error(1)
}

func TestChatHandler_CreateChat(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockChatUseCase)
	handler := api.InitiateChatHandler(mockUseCase)

	app.Post("/chat", handler.CreateChat)

	tests := []struct {
		name         string
		input        entities.Chat
		expectedCode int
		expectError  bool
	}{
		{
			name: "Valid Chat Creation",
			input: entities.Chat{
				UserID: 1,
			},
			expectedCode: fiber.StatusCreated,
			expectError:  false,
		},
		{
			name:         "Invalid Chat - No UserID",
			input:        entities.Chat{},
			expectedCode: fiber.StatusBadRequest,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectError {
				mockUseCase.On("CreateChat", tt.input).Return(tt.input, nil).Once()
			}

			jsonData, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			if !tt.expectError {
				var responseChat entities.Chat
				json.NewDecoder(resp.Body).Decode(&responseChat)
				assert.Equal(t, tt.input.UserID, responseChat.UserID)
			}
		})
	}
}

func TestChatHandler_GetAllChats(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockChatUseCase)
	handler := api.InitiateChatHandler(mockUseCase)

	app.Get("/chats", handler.GetAllChats)

	chats := []entities.Chat{
		{
			Model:  gorm.Model{ID: 1},
			UserID: 1,
		},
		{
			Model:  gorm.Model{ID: 2},
			UserID: 2,
		},
	}

	mockUseCase.On("GetAllChats").Return(chats, nil)

	req := httptest.NewRequest(http.MethodGet, "/chats", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var responseChats []entities.Chat
	json.NewDecoder(resp.Body).Decode(&responseChats)
	assert.Equal(t, len(chats), len(responseChats))
	for i, chat := range chats {
		assert.Equal(t, chat.UserID, responseChats[i].UserID)
	}
}

func TestChatHandler_GetChatByUserID(t *testing.T) {
	app := fiber.New()
	mockUseCase := new(MockChatUseCase)
	handler := api.InitiateChatHandler(mockUseCase)

	app.Get("/chat/:id", handler.GetChatByUserID)

	tests := []struct {
		name         string
		userID       string
		chat         entities.Chat
		expectedCode int
		expectError  bool
	}{
		{
			name:   "Valid User ID",
			userID: "1",
			chat: entities.Chat{
				Model:  gorm.Model{ID: 1},
				UserID: 1,
			},
			expectedCode: fiber.StatusOK,
			expectError:  false,
		},
		{
			name:         "Invalid User ID",
			userID:       "999",
			chat:         entities.Chat{},
			expectedCode: fiber.StatusNotFound,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectError {
				mockUseCase.On("GetChatByUserID", tt.userID).Return(tt.chat, nil).Once()
			} else {
				mockUseCase.On("GetChatByUserID", tt.userID).Return(entities.Chat{}, errors.New("chat not found")).Once()
			}

			req := httptest.NewRequest(http.MethodGet, "/chat/"+tt.userID, nil)
			resp, _ := app.Test(req)

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			if !tt.expectError {
				var responseChat entities.Chat
				json.NewDecoder(resp.Body).Decode(&responseChat)
				assert.Equal(t, tt.chat.UserID, responseChat.UserID)
			}
		})
	}
}
