package repositories

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/response"
)

type AuthRepository interface {
	Login(email string, password string) error
	Register(entities.User) (*response.UserResponse, error)
}
