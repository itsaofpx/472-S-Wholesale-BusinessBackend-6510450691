package request

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	ID           int      `json:"id" gorm:"primaryKey"`
	CredentialID string   `json:"credential_id"`
	FName        string   `json:"f_name"`
	LName        string   `json:"l_name"`
	PhoneNumber  string   `json:"phone_number"`
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Address      string   `json:"address"`
}

type ChangePasswordRequest struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
