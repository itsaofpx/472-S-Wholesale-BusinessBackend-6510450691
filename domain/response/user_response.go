package response

type UserResponse struct {
	ID           int    `json:"id"`
	CredentialID string `json:"credential_id"`
	FName        string `json:"f_name"`
	LName        string `json:"l_name"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	Role         int    `json:"role"`
	TierRank     int    `json:"tier_rank"`
	Address      string `json:"address"`
}

type GetUserResponse UserResponse

type GetUsersResponse struct {
	Users []GetUserResponse
}
