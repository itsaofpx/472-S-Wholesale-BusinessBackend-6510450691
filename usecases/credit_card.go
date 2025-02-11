package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type CreditCardUseCase interface {
	CreateCreditCard(cc entities.CreditCard) (entities.CreditCard, error)
	GetCreditCardByEmail(email string) (entities.CreditCard, error)
	UpdateCreditCardByEmail(email string, cc entities.CreditCard) (entities.CreditCard, error)
	DeleteCreditCardByEmail(email string) error
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

func (ccs *CreditCardService) GetCreditCardByEmail(email string) (entities.CreditCard, error) {
	creditCard, err := ccs.repo.GetCreditCardByEmail(email)
	if err != nil {
		return entities.CreditCard{}, err
	}
	return creditCard, nil
}

func (ccs *CreditCardService) UpdateCreditCardByEmail(email string, cc entities.CreditCard) (entities.CreditCard, error) {
	updatedCreditCard, err := ccs.repo.UpdateCreditCardByEmail(email, cc)
	if err != nil {
		return entities.CreditCard{}, err
	}
	return updatedCreditCard, nil
}

func (ccs *CreditCardService) DeleteCreditCardByEmail(email string) error {
	err := ccs.repo.DeleteCreditCardByEmail(email)
	if err != nil {
		return err
	}
	return nil
}
