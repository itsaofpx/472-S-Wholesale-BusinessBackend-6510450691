package usecases

import (
	"errors"
	"regexp"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/response"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(email string, password string) (*response.AuthResponse, error)
	Register(user *entities.User) error
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


func (a *AuthService) Register(user *entities.User) error {
	existUser, err := a.repo.FindUserByEmail(user.Email)
	if err == nil && existUser != nil {
		return errors.New("this email is already used")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email")
	}

	if !isValidPhoneNumber(user.PhoneNumber) {
		return errors.New("invalid phone number")
	}

	if !isValidCredentialID(user.CredentialID) {
		return errors.New("invalid credential ID")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("can't hash password")
	}

	user.Password = string(hashedPassword)
	user.Status = "A"
	user.Role = 1
	user.TierRank = 1

	if err := a.repo.CreateUser(user); err != nil {
		return errors.New("cannot create user, try again later")
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
