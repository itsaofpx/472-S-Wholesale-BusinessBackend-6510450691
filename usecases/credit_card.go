package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type CreditCardUseCase interface {
	CreateCreditCard(cc entities.CreditCard) (entities.CreditCard, error)
	GetCreditCardByUserID(id int) (entities.CreditCard, error)
	UpdateCreditCardByUserID(id int, cc entities.CreditCard) (entities.CreditCard, error)
	DeleteCreditCardByUserID(id int) error
	GetCreditCardsByUserID(userID int) ([]entities.CreditCard, error)
	DeleteByCardNumber(cardNumber string) error // เพิ่มฟังก์ชันนี้

}

type CreditCardService struct {
	repo repositories.CreditCardRepository
}

func InitiateCreditCardService(repo repositories.CreditCardRepository) CreditCardUseCase {
	return &CreditCardService{
		repo: repo,
	}
}

func (ccs *CreditCardService) CreateCreditCard(cc entities.CreditCard) (entities.CreditCard, error) {
	creditCard, err := ccs.repo.CreateCreditCard(cc)
	if err != nil {
		return entities.CreditCard{}, err
	}
	return creditCard, nil
}

func (ccs *CreditCardService) GetCreditCardByUserID(id int) (entities.CreditCard, error) {
	creditCard, err := ccs.repo.GetCreditCardByUserID(id)
	if err != nil {
		return entities.CreditCard{}, err
	}
	return creditCard, nil
}

func (ccs *CreditCardService) UpdateCreditCardByUserID(id int, cc entities.CreditCard) (entities.CreditCard, error) {
	updatedCreditCard, err := ccs.repo.UpdateCreditCardByUserID(id, cc)
	if err != nil {
		return entities.CreditCard{}, err
	}
	return updatedCreditCard, nil
}

func (ccs *CreditCardService) DeleteCreditCardByUserID(id int) error {
	err := ccs.repo.DeleteCreditCardByUserID(id)
	if err != nil {
		return err
	}
	return nil
}

func (ccs *CreditCardService) GetCreditCardsByUserID(userID int) ([]entities.CreditCard, error) {
	creditCards, err := ccs.repo.GetCreditCardsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return creditCards, nil
}
func (ccs *CreditCardService) DeleteByCardNumber(cardNumber string) error {
	err := ccs.repo.DeleteByCardNumber(cardNumber)
	if err != nil {
		return err
	}
	return nil
}