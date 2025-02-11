package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type CreditCardRepository interface {
	CreateCreditCard(cc entities.CreditCard) (entities.CreditCard, error)
	GetCreditCardByEmail(email string) (entities.CreditCard, error)
	UpdateCreditCardByEmail(email string, cc entities.CreditCard) (entities.CreditCard, error)
	DeleteCreditCardByEmail(email string) error
}
