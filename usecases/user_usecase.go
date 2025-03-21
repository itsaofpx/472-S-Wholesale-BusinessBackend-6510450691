package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type UserUseCase interface {
	GetUserByID(id int) (*entities.User, error)
	GetAllUsers() (*[]entities.User, error)
	UpdateUserByID(id int, user *request.UpdateUserByIDRequest) (*entities.User, error)
	UpdateTierByUserID(req *request.UpdateTierByUserIDRequest) (*entities.User, error)
}

type UserService struct {
	repo repositories.UserRepository
}


func InitiateUserService(repo repositories.UserRepository) UserUseCase {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetUserByID(id int) (*entities.User, error) {
	user, err := us.repo.GetUserByID(id)
	if user == nil {
		return &entities.User{}, nil
	}
	if err != nil {
		return &entities.User{}, err
	}

	return user, nil
}

func (us *UserService) GetAllUsers() (*[]entities.User, error) {
	users, err := us.repo.GetAllUsers()
	if users == nil {
		return nil, nil
	}

	if err != nil {
		return &[]entities.User{}, err
	}

	return users, nil
}

func (us *UserService) UpdateUserByID(id int, user *request.UpdateUserByIDRequest) (*entities.User, error) {
	exist, err := us.repo.GetUserByID(id)
	
	if exist == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	updated, err := us.repo.UpdateUserByID(id, user)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (us *UserService) UpdateTierByUserID(req *request.UpdateTierByUserIDRequest) (*entities.User, error) {
	exist, err := us.repo.GetUserByID(req.ID)

	if exist == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	updated, err := us.repo.UpdateUserTierByID(req, exist)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
