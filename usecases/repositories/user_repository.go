package repositories

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
)

type UserRepository interface {
	CreateUser(user *entities.User) (string, error)
	GetUserByID(id int) (*entities.User, error)
	GetAllUsers() (*[]entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	// UpdateUserTierByID(id int,*entities.User) (*entities.User, error)
	UpdateUserByID(id int, user *request.UpdateUserByIDRequest) (*entities.User, error)
	UpdateUserTierByID(req *request.UpdateTierByUserIDRequest,user *entities.User) (*entities.User, error)
	ChangePassword(req *request.ChangePasswordRequest) error
}
