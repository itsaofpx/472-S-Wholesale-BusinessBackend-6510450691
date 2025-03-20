package usecases

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/response"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"golang.org/x/crypto/bcrypt"

)

type AuthUsecase interface {
	Login(email string, password string) (*response.AuthResponse, error)
	ChangePassword (req *request.ChangePasswordRequest) error
	Register(user *entities.User) (*response.UserResponse, error)

}

type AuthService struct {
	repo repositories.UserRepository
}

func InitiateAuthService(repo repositories.UserRepository) AuthUsecase {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) Login(email string, password string) (*response.AuthResponse, error) {
	// Fetch user from database
	existUser, err := a.repo.FindUserByEmail(email)
	if err != nil || existUser == nil {
		// Return a generic message to prevent email enumeration attacks
		return nil, errors.New("incorrect email or password")
	}

	// Compare hashed password with input password
	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(password)); err != nil {
		return nil, errors.New("incorrect email or password")
	}

	// Create authentication response
	response := &response.AuthResponse{
		ID:       existUser.ID,
		Role:     existUser.Role,
		TierRank: existUser.TierRank,
	}

	return response, nil
}


func (a *AuthService) Register(user *entities.User) (*response.UserResponse, error) {
	existUser, err := a.repo.FindUserByEmail(user.Email)
	if err == nil && existUser != nil {
		return nil, errors.New("this email is already used")
	}

	if !isValidEmail(user.Email) {
		return nil, errors.New("invalid email")
	}

	if !isValidPhoneNumber(user.PhoneNumber) {
		return nil, errors.New("invalid phone number")
	}

	if !isValidCredentialID(user.CredentialID) {
		return nil, errors.New("invalid credential ID")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("can't hash password")
	}

	user.Password = string(hashedPassword)
	user.Status = "A"
	user.Role = 1
	user.TierRank = 1
	
	id, err := a.repo.CreateUser(user);
	if err != nil {
		return nil, errors.New("cannot create user, try again later")
	}
	parsedID, err :=  strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Parsed ID:", parsedID)
	response := &response.UserResponse{
		ID:           parsedID,
		CredentialID: user.CredentialID,
		FName:        user.FName,
		LName:        user.LName,
		PhoneNumber:  user.PhoneNumber,
		Email:        user.Email,
		Status:       user.Status,
		Role:         user.Role,
		TierRank:     user.TierRank,
		Address:      user.Address,
	}
	return response , nil
}

func (a *AuthService) ChangePassword(req *request.ChangePasswordRequest) error {
	// Find user by email
	existUser, err := a.repo.FindUserByEmail(req.Email)
	if err != nil || existUser == nil {
		return errors.New("user not found")
	}

	fmt.Println(existUser.Password + " user passoword")
	fmt.Println(req.OldPassword + " old password")

	// Compare the old password
	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(req.OldPassword)); err != nil {
		return errors.New("incorrect old password")
	}

	// Hash the new password
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("cannot hash new password")
	}

	// Update the password in the user record
	req.NewPassword = string(hashedNewPassword)
	if err := a.repo.ChangePassword(req); err != nil {
		return errors.New("cannot update password, try again later")
	}

	return nil
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

func isValidPhoneNumber(phone string) bool {
	regex := `^(0(6|8|9)\d{8}|0(2|[3-9])\d{7})$`
	return regexp.MustCompile(regex).MatchString(phone)
}

func isValidCredentialID(credentialID string) bool {
	regex := `^\d{13}$`
	return regexp.MustCompile(regex).MatchString(credentialID)
}
