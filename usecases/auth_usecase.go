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
	Login(email string, password string) (err error, res *response.AuthResponse)
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

func (a *AuthService) Login(email string, password string) (err error, res *response.AuthResponse) {
	existUser, err := a.repo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found"), nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(password)); err != nil {
		return errors.New("password doesn't match"), nil
	}
	
	var response *response.AuthResponse = &response.AuthResponse{ID: existUser.ID, Role: existUser.Role, TierRank: existUser.TierRank}

	return nil, response
}

func (a *AuthService) Register(user *entities.User) error {
	existUser, err := a.repo.FindUserByEmail(user.Email)

	if err != nil || existUser != nil {
		return errors.New("this email is already used")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(user.Email) {
		return errors.New("invalid email")
	}

	if !regexp.MustCompile(`^(0(6|8|9)\d{8}|0(2|[3-9])\d{7})$`).MatchString(user.PhoneNumber) {
		return errors.New("invalid phone number")
	}

	if !regexp.MustCompile(`^\d{13}$`).MatchString(user.CredentialID) {
		return errors.New("invalid credential id")
	}


	

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.Status = "A"
	user.Role = 1
	user.TierRank = 1

	if err := a.repo.CreateUser(user); err != nil {
		return errors.New("can not create user, try again later")
	}

	return nil

}
