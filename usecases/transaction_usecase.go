package usecases

import (
	"time"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type TransactionUsecase interface {
	CreateTransaction(t entities.Transaction) (entities.Transaction, error)
	GetTransactionById(id int) (entities.Transaction, error)
	GetTransactionByOrderId(orderId int) (entities.Transaction, error)
	GetAllTransactions() ([]entities.Transaction, error)
	UpdateTransaction(id int, t entities.Transaction) (entities.Transaction, error)
	DeleteTransaction(id int) error
}

type TransactionService struct {
	repo repositories.TransactionRepository
}

func InitiateTransactionService(repo repositories.TransactionRepository) TransactionUsecase {
	return &TransactionService{
		repo: repo,
	}
}

func (ts *TransactionService) CreateTransaction(t entities.Transaction) (entities.Transaction, error) {
	t.T_time_stamp = time.Now()
	createdTransaction, err := ts.repo.CreateTransaction(t)

	if err != nil {
		return entities.Transaction{}, err
	}

	return createdTransaction, nil
}

func (ts *TransactionService) GetTransactionById(id int) (entities.Transaction, error) {
	t, err := ts.repo.GetTransactionById(id)
	if err != nil {
		return entities.Transaction{}, err
	}
	return t, nil
}

func (ts *TransactionService) GetAllTransactions() ([]entities.Transaction, error) {
	t_list, err := ts.repo.GetAllTransactions()
	if err != nil {
		return []entities.Transaction{}, err
	}
	return t_list, nil
}

func (ts *TransactionService) UpdateTransaction(id int, updatedTransaction entities.Transaction) (entities.Transaction, error) {
	updatedTrans, err := ts.repo.UpdateTransaction(id, updatedTransaction)
	if err != nil {
		return entities.Transaction{}, err
	}
	return updatedTrans, nil
}

func (ts *TransactionService) GetTransactionByOrderId(orderId int) (entities.Transaction, error) {
	t, err := ts.repo.GetTransactionByOrderId(orderId)
	if err != nil {
		return entities.Transaction{}, err
	}
	return t, nil
}

func (ts *TransactionService) DeleteTransaction(id int) error {
	err := ts.repo.DeleteTransaction(id)
	if err != nil {
		return err
	}
	return nil
}
