package response

type CreditCardResponse struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	CardNumber     string `json:"card_number"`
	CardHolder     string `json:"card_holder"`
	Expiration     string `json:"expiration"`
	SecurityCode   string `json:"security_code"`
}
