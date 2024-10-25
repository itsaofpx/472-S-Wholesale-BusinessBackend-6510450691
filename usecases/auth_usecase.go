package usecases

import (
	"errors"
	"fmt"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(email string, password string) error
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

func (a *AuthService) Login(email string, password string) error {
	existUser, err := a.repo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(password)); err != nil {
		return errors.New("password doesn't match")
	}

	fmt.Println(existUser)

	return nil
}

func (a *AuthService) Register(user *entities.User) error {
	existUser, err := a.repo.FindUserByEmail(user.Email)
	
	if err != nil || existUser != nil {
		return errors.New("this email is already used")
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
