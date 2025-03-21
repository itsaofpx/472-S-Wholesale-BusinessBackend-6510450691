package tests

import (
	"errors"
	"testing"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/usecases"
	"github.com/ppwlsw/sa-project-backend/mocks"
	"github.com/stretchr/testify/assert"
	_"github.com/stretchr/testify/mock"
)

func TestGetUserByID(t *testing.T) {
	// Test cases
	testCases := []struct {
		name          string
		userID        int
		mockUser      *entities.User
		mockError     error
		expectedUser  *entities.User
		expectedError error
	}{
		{
			name:   "Success",
			userID: 1,
			mockUser: &entities.User{
				ID:           1,
				CredentialID: "1234567890123",
				FName:        "John",
				LName:        "Doe",
				Email:        "john.doe@example.com",
				PhoneNumber:  "0812345678",
				Status:       "A",
				Role:         1,
				TierRank:     1,
				Address:      "123 Main St",
			},
			mockError:     nil,
			expectedUser:  &entities.User{ID: 1, FName: "John", LName: "Doe"},
			expectedError: nil,
		},
		{
			name:          "User Not Found",
			userID:        999,
			mockUser:      nil,
			mockError:     nil,
			expectedUser:  &entities.User{},
			expectedError: nil,
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := new(mocks.MockUserRepository)
			mockRepo.On("GetUserByID", tc.userID).Return(tc.mockUser, tc.mockError)

			// Create service with mock repository
			userService := usecases.InitiateUserService(mockRepo)

			// Call the method being tested
			user, err := userService.GetUserByID(tc.userID)

			// Assertions
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			if tc.mockUser == nil {
				assert.Equal(t, tc.expectedUser, user)
			} else {
				assert.Equal(t, tc.mockUser.ID, user.ID)
				assert.Equal(t, tc.mockUser.FName, user.FName)
				assert.Equal(t, tc.mockUser.LName, user.LName)
			}

			// Verify that the expected methods were called
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	// Test cases
	testCases := []struct {
		name           string
		mockUsers      *[]entities.User
		mockError      error
		expectedUsers  *[]entities.User
		expectedError  error
		expectedResult bool
	}{
		{
			name: "Success",
			mockUsers: &[]entities.User{
				{ID: 1, FName: "John", LName: "Doe"},
				{ID: 2, FName: "Jane", LName: "Smith"},
			},
			mockError:      nil,
			expectedUsers:  &[]entities.User{{ID: 1}, {ID: 2}},
			expectedError:  nil,
			expectedResult: true,
		},
		{
			name:           "No Users Found",
			mockUsers:      nil,
			mockError:      nil,
			expectedUsers:  nil,
			expectedError:  nil,
			expectedResult: true,
		},
		{
			name:           "Database Error",
			mockUsers:      &[]entities.User{},
			mockError:      errors.New("database error"),
			expectedUsers:  &[]entities.User{},
			expectedError:  errors.New("database error"),
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := new(mocks.MockUserRepository)
			mockRepo.On("GetAllUsers").Return(tc.mockUsers, tc.mockError)

			// Create service with mock repository
			userService := usecases.InitiateUserService(mockRepo)

			// Call the method being tested
			users, err := userService.GetAllUsers()

			// Assertions
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			if tc.mockUsers == nil {
				assert.Nil(t, users)
			} else if tc.expectedResult {
				assert.Equal(t, len(*tc.mockUsers), len(*users))
				for i, user := range *users {
					assert.Equal(t, (*tc.mockUsers)[i].ID, user.ID)
				}
			}

			// Verify that the expected methods were called
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateUserByID(t *testing.T) {
	// Test cases
	testCases := []struct {
		name          string
		userID        int
		updateReq     *request.UpdateUserByIDRequest
		mockExistUser *entities.User
		mockUpdated   *entities.User
		mockError     error
		expectedError error
	}{
		{
			name:   "Success",
			userID: 1,
			updateReq: &request.UpdateUserByIDRequest{
				FName:       "Updated",
				LName:       "User",
				PhoneNumber: "0987654321",
				Email:       "updated@example.com",
				Address:     "456 New St",
			},
			mockExistUser: &entities.User{ID: 1, FName: "John", LName: "Doe"},
			mockUpdated: &entities.User{
				ID:          1,
				FName:       "Updated",
				LName:       "User",
				PhoneNumber: "0987654321",
				Email:       "updated@example.com",
				Address:     "456 New St",
			},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "User Not Found",
			userID:        999,
			updateReq:     &request.UpdateUserByIDRequest{},
			mockExistUser: nil,
			mockUpdated:   nil,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Database Error",
			userID:        2,
			updateReq:     &request.UpdateUserByIDRequest{},
			mockExistUser: &entities.User{ID: 2},
			mockUpdated:   nil,
			mockError:     errors.New("database error"),
			expectedError: errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := new(mocks.MockUserRepository)
			mockRepo.On("GetUserByID", tc.userID).Return(tc.mockExistUser, nil)
			
			if tc.mockExistUser != nil {
				mockRepo.On("UpdateUserByID", tc.userID, tc.updateReq).Return(tc.mockUpdated, tc.mockError)
			}

			// Create service with mock repository
			userService := usecases.InitiateUserService(mockRepo)

			// Call the method being tested
			user, err := userService.UpdateUserByID(tc.userID, tc.updateReq)

			// Assertions
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			if tc.mockExistUser == nil {
				assert.Nil(t, user)
			} else if tc.mockUpdated != nil {
				assert.Equal(t, tc.mockUpdated.ID, user.ID)
				assert.Equal(t, tc.mockUpdated.FName, user.FName)
				assert.Equal(t, tc.mockUpdated.LName, user.LName)
			}

			// Verify that the expected methods were called
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTierByUserID(t *testing.T) {
	// Test cases
	testCases := []struct {
		name          string
		updateReq     *request.UpdateTierByUserIDRequest
		mockExistUser *entities.User
		mockUpdated   *entities.User
		mockError     error
		expectedError error
	}{
		{
			name: "Success",
			updateReq: &request.UpdateTierByUserIDRequest{
				ID:   1,
				Tier: 2,
			},
			mockExistUser: &entities.User{ID: 1, TierRank: 1},
			mockUpdated:   &entities.User{ID: 1, TierRank: 2},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name: "User Not Found",
			updateReq: &request.UpdateTierByUserIDRequest{
				ID:   999,
				Tier: 2,
			},
			mockExistUser: nil,
			mockUpdated:   nil,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name: "Database Error",
			updateReq: &request.UpdateTierByUserIDRequest{
				ID:   2,
				Tier: 3,
			},
			mockExistUser: &entities.User{ID: 2, TierRank: 1},
			mockUpdated:   nil,
			mockError:     errors.New("database error"),
			expectedError: errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := new(mocks.MockUserRepository)
			mockRepo.On("GetUserByID", tc.updateReq.ID).Return(tc.mockExistUser, nil)
			
			if tc.mockExistUser != nil {
				mockRepo.On("UpdateUserTierByID", tc.updateReq, tc.mockExistUser).Return(tc.mockUpdated, tc.mockError)
			}

			// Create service with mock repository
			userService := usecases.InitiateUserService(mockRepo)

			// Call the method being tested
			user, err := userService.UpdateTierByUserID(tc.updateReq)

			// Assertions
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			if tc.mockExistUser == nil {
				assert.Nil(t, user)
			} else if tc.mockUpdated != nil {
				assert.Equal(t, tc.mockUpdated.ID, user.ID)
				assert.Equal(t, tc.mockUpdated.TierRank, user.TierRank)
			}

			// Verify that the expected methods were called
			mockRepo.AssertExpectations(t)
		})
	}
}
