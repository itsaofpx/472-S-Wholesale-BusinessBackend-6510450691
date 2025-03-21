package tests

import (
	"errors"
	"testing"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/mocks"
	"github.com/ppwlsw/sa-project-backend/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	// Create a hashed password for testing
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Test cases
	testCases := []struct {
		name           string
		email          string
		password       string
		mockUser       *entities.User
		mockError      error
		expectedError  bool
		expectedErrMsg string
	}{
		{
			name:     "Success",
			email:    "john.doe@example.com",
			password: "password123",
			mockUser: &entities.User{
				ID:       1,
				Email:    "john.doe@example.com",
				Password: string(hashedPassword),
				Role:     1,
				TierRank: 1,
			},
			mockError:      nil,
			expectedError:  false,
			expectedErrMsg: "",
		},
		{
			name:           "User Not Found",
			email:          "nonexistent@example.com",
			password:       "password123",
			mockUser:       nil,
			mockError:      errors.New("user not found"),
			expectedError:  true,
			expectedErrMsg: "incorrect email or password",
		},
		{
			name:     "Incorrect Password",
			email:    "john.doe@example.com",
			password: "wrongpassword",
			mockUser: &entities.User{
				ID:       1,
				Email:    "john.doe@example.com",
				Password: string(hashedPassword),
				Role:     1,
				TierRank: 1,
			},
			mockError:      nil,
			expectedError:  true,
			expectedErrMsg: "incorrect email or password",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := new(mocks.MockUserRepository)
			mockRepo.On("FindUserByEmail", tc.email).Return(tc.mockUser, tc.mockError)

			// Create service with mock repository
			authService := usecases.InitiateAuthService(mockRepo)

			// Call the method being tested
			response, err := authService.Login(tc.email, tc.password)

			// Assertions
			if tc.expectedError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErrMsg, err.Error())
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, tc.mockUser.ID, response.ID)
				assert.Equal(t, tc.mockUser.Role, response.Role)
				assert.Equal(t, tc.mockUser.TierRank, response.TierRank)
			}

			// Verify that the expected methods were called
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestRegister(t *testing.T) {
	// Test cases
	testCases := []struct {
		name           string
		user           *entities.User
		mockExistUser  *entities.User
		mockFindError  error
		mockCreateID   string
		mockCreateErr  error
		expectedError  bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			user: &entities.User{
				CredentialID: "1234567890123",
				FName:        "John",
				LName:        "Doe",
				PhoneNumber:  "0812345678",
				Email:        "john.doe@example.com",
				Password:     "password123",
				Address:      "123 Main St",
			},
			mockExistUser:  nil,
			mockFindError:  errors.New("user not found"),
			mockCreateID:   "1",
			mockCreateErr:  nil,
			expectedError:  false,
			expectedErrMsg: "",
		},
		{
			name: "Email Already Used",
			user: &entities.User{
				Email: "existing@example.com",
			},
			mockExistUser: &entities.User{
				ID:    1,
				Email: "existing@example.com",
			},
			mockFindError:  nil,
			mockCreateID:   "",
			mockCreateErr:  nil,
			expectedError:  true,
			expectedErrMsg: "this email is already used",
		},
		{
			name: "Invalid Email",
			user: &entities.User{
				Email: "invalid-email",
			},
			mockExistUser:  nil,
			mockFindError:  errors.New("user not found"),
			mockCreateID:   "",
			mockCreateErr:  nil,
			expectedError:  true,
			expectedErrMsg: "invalid email",
		},
		{
			name: "Invalid Phone Number",
			user: &entities.User{
				Email:       "valid@example.com",
				PhoneNumber: "123", // Invalid phone number
			},
			mockExistUser:  nil,
			mockFindError:  errors.New("user not found"),
			mockCreateID:   "",
			mockCreateErr:  nil,
			expectedError:  true,
			expectedErrMsg: "invalid phone number",
		},
		{
			name: "Invalid Credential ID",
			user: &entities.User{
				Email:        "valid@example.com",
				PhoneNumber:  "0812345678",
				CredentialID: "123", // Invalid credential ID
			},
			mockExistUser:  nil,
			mockFindError:  errors.New("user not found"),
			mockCreateID:   "",
			mockCreateErr:  nil,
			expectedError:  true,
			expectedErrMsg: "invalid credential ID",
		},
		{
			name: "Create User Error",
			user: &entities.User{
				CredentialID: "1234567890123",
				FName:        "John",
				LName:        "Doe",
				PhoneNumber:  "0812345678",
				Email:        "john.doe@example.com",
				Password:     "password123",
				Address:      "123 Main St",
			},
			mockExistUser:  nil,
			mockFindError:  errors.New("user not found"),
			mockCreateID:   "",
			mockCreateErr:  errors.New("database error"),
			expectedError:  true,
			expectedErrMsg: "cannot create user, try again later",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := new(mocks.MockUserRepository)
			mockRepo.On("FindUserByEmail", tc.user.Email).Return(tc.mockExistUser, tc.mockFindError)
			
			// Only set up CreateUser expectation if we expect it to be called
			if tc.mockExistUser == nil && tc.user.Email == "john.doe@example.com" {
				mockRepo.On("CreateUser", mock.AnythingOfType("*entities.User")).Return(tc.mockCreateID, tc.mockCreateErr)
			}

			// Create service with mock repository
			authService := usecases.InitiateAuthService(mockRepo)

			// Call the method being tested
			response, err := authService.Register(tc.user)

			// Assertions
			if tc.expectedError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErrMsg, err.Error())
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 1, response.ID) // Assuming ID is 1 from the mock
				assert.Equal(t, tc.user.FName, response.FName)
				assert.Equal(t, tc.user.LName, response.LName)
				assert.Equal(t, tc.user.Email, response.Email)
				assert.Equal(t, "A", response.Status)
				assert.Equal(t, 1, response.Role)
				assert.Equal(t, 1, response.TierRank)
			}

			// Verify that the expected methods were called
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestChangePassword(t *testing.T) {
	// Create a hashed password for testing
	oldPassword := "oldPassword123"
	hashedOldPassword, _ := bcrypt.GenerateFromPassword([]byte(oldPassword), bcrypt.DefaultCost)

	// Test cases
	testCases := []struct {
		name           string
		req            *request.ChangePasswordRequest
		mockUser       *entities.User
		mockFindError  error
		mockChangeErr  error
		expectedError  bool
		expectedErrMsg string
	}{
		{
			name: "Success",
			req: &request.ChangePasswordRequest{
				Email:       "john.doe@example.com",
				OldPassword: "oldPassword123",
				NewPassword: "newPassword123",
			},
			mockUser: &entities.User{
				ID:       1,
				Email:    "john.doe@example.com",
				Password: string(hashedOldPassword),
			},
			mockFindError:  nil,
			mockChangeErr:  nil,
			expectedError:  false,
			expectedErrMsg: "",
		},
		{
			name: "User Not Found",
			req: &request.ChangePasswordRequest{
				Email: "nonexistent@example.com",
			},
			mockUser:       nil,
			mockFindError:  errors.New("user not found"),
			mockChangeErr:  nil,
			expectedError:  true,
			expectedErrMsg: "user not found",
		},
		{
			name: "Incorrect Old Password",
			req: &request.ChangePasswordRequest{
				Email:       "john.doe@example.com",
				OldPassword: "wrongPassword",
				NewPassword: "newPassword123",
			},
			mockUser: &entities.User{
				ID:       1,
				Email:    "john.doe@example.com",
				Password: string(hashedOldPassword),
			},
			mockFindError:  nil,
			mockChangeErr:  nil,
			expectedError:  true,
			expectedErrMsg: "incorrect old password",
		},
		{
			name: "Database Error",
			req: &request.ChangePasswordRequest{
				Email:       "john.doe@example.com",
				OldPassword: "oldPassword123",
				NewPassword: "newPassword123",
			},
			mockUser: &entities.User{
				ID:       1,
				Email:    "john.doe@example.com",
				Password: string(hashedOldPassword),
			},
			mockFindError:  nil,
			mockChangeErr:  errors.New("database error"),
			expectedError:  true,
			expectedErrMsg: "cannot update password, try again later",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := new(mocks.MockUserRepository)
			mockRepo.On("FindUserByEmail", tc.req.Email).Return(tc.mockUser, tc.mockFindError)
			
			// Only set up ChangePassword expectation if we expect it to be called
			if tc.mockUser != nil && tc.req.OldPassword == "oldPassword123" {
				mockRepo.On("ChangePassword", mock.AnythingOfType("*request.ChangePasswordRequest")).Return(tc.mockChangeErr)
			}

			// Create service with mock repository
			authService := usecases.InitiateAuthService(mockRepo)

			// Call the method being tested
			err := authService.ChangePassword(tc.req)

			// Assertions
			if tc.expectedError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErrMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}

			// Verify that the expected methods were called
			mockRepo.AssertExpectations(t)
		})
	}
}
