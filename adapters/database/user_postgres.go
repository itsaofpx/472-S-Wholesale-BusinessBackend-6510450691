package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
	"errors"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func InitiateUserPostgresRepository(db *gorm.DB) repositories.UserRepository {
	return &UserPostgresRepository{db: db}
}

func (upr *UserPostgresRepository) CreateUser(newUser *entities.User) (string, error) {
	query := `
		INSERT INTO public.users 
		(credential_id, f_name, l_name, phone_number, email, password, status, role, tier_rank, address) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		RETURNING id
	`

	var userID string
	err := upr.db.Raw(query, newUser.CredentialID, newUser.FName, newUser.LName, newUser.PhoneNumber, newUser.Email, newUser.Password, newUser.Status, newUser.Role, newUser.TierRank, newUser.Address).Scan(&userID).Error
	
	if err != nil {
		return "0", err
	}
	return userID, nil
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


func (upr *UserPostgresRepository) UpdateUserByID(id int, user *request.UpdateUserByIDRequest) (*entities.User, error) {
	query := `
		UPDATE users 
		SET f_name=$1, l_name=$2, phone_number=$3, email=$4, address=$5
		WHERE id = $6 
		RETURNING id, credential_id, f_name, l_name, phone_number, email, password, status, role, tier_rank, address;
	`


	var updatedUser entities.User
	result := upr.db.Raw(query, user.FName, user.LName, user.PhoneNumber, user.Email, user.Address, id).Scan(&updatedUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &updatedUser, nil
}


func (upr *UserPostgresRepository) UpdateUserTierByID(req *request.UpdateTierByUserIDRequest, user *entities.User) (*entities.User, error) {
	query := "UPDATE users as u SET tier_rank=$1 WHERE u.id = $2 RETURNING id, credential_id, f_name, l_name, phone_number, email, password, status, role, tier_rank, address;"

	result := upr.db.Raw(query, req.Tier, req.ID).Scan(&user)
	if result.Error != nil {
		return &entities.User{}, result.Error
	}

	return user, nil

}

func (upr *UserPostgresRepository) ChangePassword(req *request.ChangePasswordRequest) error {
	// ค้นหาผู้ใช้จาก email
	var user entities.User
	query := "SELECT id, credential_id, f_name, l_name, phone_number, email, password, status, role, tier_rank, address FROM users WHERE email = $1"
	result := upr.db.Raw(query, req.Email).Scan(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	// อัพเดตข้อมูลผู้ใช้ด้วยรหัสผ่านใหม่
	updateQuery := "UPDATE users SET password = $1 WHERE email = $2"
	updateResult := upr.db.Exec(updateQuery, req.NewPassword, req.Email)

	if updateResult.Error != nil {
		return updateResult.Error
	}

	return nil
}