package request

type UpdateTierByUserIDRequest struct {
	ID int `json:"id"`
	Tier int `json:"tier_rank"`
}

type UpdateUserByIDRequest struct {
	FName        string   `json:"f_name"`
	LName        string   `json:"l_name"`
	PhoneNumber  string   `json:"phone_number"`
	Email        string   `json:"email"`
	Address      string   `json:"address"`
}
