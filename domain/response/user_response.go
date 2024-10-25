package response

type UserResponse struct {
	ID           int    `json:"id"`
	CredentialID string `json:"credential_id"`
	Name         string `json:"name"`
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
