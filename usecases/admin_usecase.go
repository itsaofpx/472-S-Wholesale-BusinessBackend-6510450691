package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"golang.org/x/crypto/bcrypt"
)
type AdminUsecase interface {
	InitializeAdmin() error
}

type AdminService struct {
	repo repositories.UserRepository
}

func InitiateAdminService( repo repositories.UserRepository) AdminUsecase {
	return &AdminService{
		repo: repo,
	}
}

func (a *AdminService) InitializeAdmin() error {
	admin_email := "admin@admin.com"
	admin_password := "admin1234"

	existUser, err := a.repo.FindUserByEmail(admin_email)
	if err != nil {
		return err
	}
	if existUser != nil {
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin_password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	existUser = &entities.User{
		CredentialID: "admin",
		FName: "Admin",
		LName: "",
		Email: admin_email,
		Password: string(hashedPassword),
		Status: "A",
		Role: 2,
		TierRank: 1,
	}
	_ , err = a.repo.CreateUser(existUser)
	if err != nil {
		return err
	}
	return nil
}