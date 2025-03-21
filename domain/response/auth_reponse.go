package response

type AuthResponse struct {
	ID int `json:"id"`
	Role int `json:"role"`
	TierRank int `json:"tier_rank"`
}
