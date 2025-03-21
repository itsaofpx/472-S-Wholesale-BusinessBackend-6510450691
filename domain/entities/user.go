package entities

type User struct {
	ID           int      `json:"id" gorm:"primaryKey"`
	CredentialID string   `json:"credential_id"`
	FName        string   `json:"f_name"`
	LName        string   `json:"l_name"`
	PhoneNumber  string   `json:"phone_number"`
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Status       string   `json:"status"`
	Role         int      `json:"role"`
	TierRank     int      `json:"tier_rank"`
	TierList     TierList `gorm:"foreignKey:TierRank;references:Tier"`
	Address      string   `json:"address"`
}
