package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type CreditCardPostgresRepository struct {
	db *gorm.DB
}

func InitiateCreditCardPostgresRepository(db *gorm.DB) repositories.CreditCardRepository {
	return &CreditCardPostgresRepository{db: db}
}

func (ccr *CreditCardPostgresRepository) CreateCreditCard(cc entities.CreditCard) (entities.CreditCard, error) {
	query := `
		INSERT INTO public.credit_cards(email, card_number, card_holder, expiration, security_code)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, email, card_number, card_holder, expiration, security_code;
	`
	var creditCard entities.CreditCard

	result := ccr.db.Raw(query, cc.Email, cc.CardNumber, cc.CardHolder, cc.Expiration, cc.SecurityCode).Scan(&creditCard)

	if result.Error != nil {
		return entities.CreditCard{}, result.Error
	}

	return creditCard, nil
}

func (ccr *CreditCardPostgresRepository) UpdateCreditCardByEmail(email string, cc entities.CreditCard) (entities.CreditCard, error) {
	query := `
		UPDATE public.credit_cards 
		SET card_number=$2, card_holder=$3, expiration=$4, security_code=$5
		WHERE email = $1 
		RETURNING id, email, card_number, card_holder, expiration, security_code;
	`
	var creditCard entities.CreditCard

	result := ccr.db.Raw(query, email, cc.CardNumber, cc.CardHolder, cc.Expiration, cc.SecurityCode).Scan(&creditCard)

	if result.Error != nil {
		return entities.CreditCard{}, result.Error
	}

	return creditCard, nil
}

func (ccr *CreditCardPostgresRepository) GetCreditCardByEmail(email string) (entities.CreditCard, error) {
	var creditCard entities.CreditCard
	result := ccr.db.Where("email = ?", email).First(&creditCard)
	if result.Error != nil {
		return entities.CreditCard{}, result.Error
	}
	return creditCard, nil
}


func (ccr *CreditCardPostgresRepository) DeleteCreditCardByEmail(email string) error {
	query := "DELETE FROM public.credit_cards WHERE email = $1;"

	result := ccr.db.Exec(query, email)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
