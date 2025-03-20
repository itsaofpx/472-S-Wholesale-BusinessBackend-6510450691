package repositories

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/domain/response"
)

type AuthRepository interface {
	Login(email string, password string) error
	ChangePassword(req *request.ChangePasswordRequest) error
	Register(entities.User) (*response.UserResponse, error)

}
