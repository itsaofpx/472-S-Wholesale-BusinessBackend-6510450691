package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type TransactionRepository interface {
	CreateTransaction(t entities.Transaction) (entities.Transaction, error)
	GetTransactionById(id int) (entities.Transaction, error)
	GetTransactionByOrderId(orderId int) (entities.Transaction, error)
	GetAllTransactions() ([]entities.Transaction, error)
	UpdateTransaction(id int, t entities.Transaction) (entities.Transaction, error)
	DeleteTransaction(id int) error
}
