package repositories

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
)

type AuthRepository interface {
	Login(email string, password string) error
	Register(entities.User) error
}
