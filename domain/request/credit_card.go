package request

type CreateCreditCardRequest struct {
	UserID        string `json:"user_id" validate:"required,email"`
	CardNumber   string `json:"card_number" validate:"required,len=16"`
	CardHolder   string `json:"card_holder" validate:"required"`
	Expiration   string `json:"expiration" validate:"required"`
	SecurityCode string `json:"security_code" validate:"required,len=3"`
}

type UpdateCreditCardRequest struct {
	CardNumber   string `json:"card_number" validate:"required,len=16"`
	CardHolder   string `json:"card_holder" validate:"required"`
	Expiration   string `json:"expiration" validate:"required"`
	SecurityCode string `json:"security_code" validate:"required,len=3"`
}
