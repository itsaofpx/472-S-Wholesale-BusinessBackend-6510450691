package repositories

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
)

type AuthRepository interface {
	Login(email string, password string) error
	Register(entities.User) error
	ChangePassword(req *request.ChangePasswordRequest) error
}
