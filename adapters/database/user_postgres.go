package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func InitiateUserPostgresRepository(db *gorm.DB) repositories.UserRepository {
	return &UserPostgresRepository{db: db}
}

func (upr *UserPostgresRepository) CreateUser(newUser *entities.User) error {
	query := "INSERT INTO public.users(credential_id, f_name, l_name,phone_number, email, password, status, role, tier_rank, address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	err := upr.db.Exec(query, newUser.CredentialID, newUser.FName, newUser.LName, newUser.PhoneNumber, newUser.Email, newUser.Password, newUser.Status, newUser.Role, newUser.TierRank, newUser.Address)
	if err != nil {
		return err.Error
	}

	return nil
}

func (upr *UserPostgresRepository) GetUserByID(id int) (*entities.User, error) {
	var user *entities.User

	query := "SELECT id, credential_id, f_name, l_name, phone_number, email, password, status, role, tier_rank, address FROM users WHERE id = $1"

	result := upr.db.Raw(query, id).Scan(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (upr *UserPostgresRepository) GetAllUsers() (*[]entities.User, error) {
	query := "SELECT id, credential_id, f_name, l_name, phone_number, email, password, status, role, tier_rank, address FROM users"
	var users *[]entities.User

	result := upr.db.Raw(query).Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil

}

func (upr *UserPostgresRepository) FindUserByEmail(email string) (*entities.User, error) {
	query := "SELECT id, credential_id, f_name, l_name, phone_number, email, password, status, role, tier_rank, address FROM users WHERE email = $1"
	user := &entities.User{}

	result := upr.db.Raw(query, email).Scan(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return user, nil
}


func (upr *UserPostgresRepository) UpdateUserTierByID(req *request.UpdateTierByUserIDRequest, user *entities.User) (*entities.User, error) {
	query := "UPDATE users as u SET tier_rank=$1 WHERE u.id = $2 RETURNING id, credential_id, f_name, l_name, phone_number, email, password, status, role, tier_rank, address;"

	result := upr.db.Raw(query, req.Tier, req.ID).Scan(&user)
	if result.Error != nil {
		return &entities.User{}, result.Error
	}

	return user, nil

}
