package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
	"errors"
)

type CreditCardPostgresRepository struct {
	db *gorm.DB
}

func InitiateCreditCardPostgresRepository(db *gorm.DB) repositories.CreditCardRepository {
	return &CreditCardPostgresRepository{
		db: db,
	}
}

func (ccr *CreditCardPostgresRepository) CreateCreditCard(cc entities.CreditCard) (entities.CreditCard, error) {
	query := `
		INSERT INTO public.credit_cards(user_id, card_number, card_holder, expiration, security_code)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, card_number, card_holder, expiration, security_code;
	`
	var creditCard entities.CreditCard

	result := ccr.db.Raw(query, cc.UserID, cc.CardNumber, cc.CardHolder, cc.Expiration, cc.SecurityCode).Scan(&creditCard)

	if result.Error != nil {
		return entities.CreditCard{}, result.Error
	}

	return creditCard, nil
}

func (ccr *CreditCardPostgresRepository) UpdateCreditCardByUserID(id int, cc entities.CreditCard) (entities.CreditCard, error) {
	query := `
		UPDATE public.credit_cards 
		SET card_number=$2, card_holder=$3, expiration=$4, security_code=$5
		WHERE user_id = $1 
		RETURNING id, user_id, card_number, card_holder, expiration, security_code;
	`
	var creditCard entities.CreditCard

	result := ccr.db.Raw(query, id, cc.CardNumber, cc.CardHolder, cc.Expiration, cc.SecurityCode).Scan(&creditCard)

	if result.Error != nil {
		return entities.CreditCard{}, result.Error
	}

	return creditCard, nil
}

func (ccr *CreditCardPostgresRepository) GetCreditCardByUserID(id int) (entities.CreditCard, error) {
	var creditCard entities.CreditCard
	result := ccr.db.Where("user_id = ?", id).First(&creditCard)
	if result.Error != nil {
		return entities.CreditCard{}, result.Error
	}
	return creditCard, nil
}


func (ccr *CreditCardPostgresRepository) DeleteCreditCardByUserID(id int) error {
    query := "DELETE FROM public.credit_cards WHERE user_id = $1;"

    result := ccr.db.Exec(query, id)

    if result.Error != nil {
        return result.Error
    }

    return nil
}

func (ccr *CreditCardPostgresRepository) GetCreditCardsByUserID(userID int) ([]entities.CreditCard, error) {
	var creditCards []entities.CreditCard
	result := ccr.db.Where("user_id = ?", userID).Find(&creditCards)

	if result.Error != nil {
		return nil, result.Error
	}

	return creditCards, nil
}

func (ccr *CreditCardPostgresRepository) DeleteByCardNumber(cardNumber string) error {
	result := ccr.db.Where("card_number = ?", cardNumber).Delete(&entities.CreditCard{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("credit card not found")
	}

	return nil
}
