package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type CreditCardRepository interface {
	CreateCreditCard(cc entities.CreditCard) (entities.CreditCard, error)
	GetCreditCardByUserID(id int) (entities.CreditCard, error)
	UpdateCreditCardByUserID(id int, cc entities.CreditCard) (entities.CreditCard, error)
	DeleteCreditCardByUserID(id int) error
	GetCreditCardsByUserID(userID int) ([]entities.CreditCard, error)
	DeleteByCardNumber(cardNumber string) error

}
