package request

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	ID           int      `json:"id" gorm:"primaryKey"`
	CredentialID string   `json:"credential_id"`
	Name         string   `json:"name"`
	PhoneNumber  string   `json:"phone_number"`
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Address      string   `json:"address"`
}