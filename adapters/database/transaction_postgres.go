package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type TransactionPostgresRepository struct {
    db *gorm.DB
}

func InitiateTransactionPostGresRepository(db *gorm.DB) repositories.TransactionRepository {
	return &TransactionPostgresRepository{
		db: db,
	}
}


func (tpr *TransactionPostgresRepository) CreateTransaction(t entities.Transaction) (entities.Transaction, error) {
    query := "INSERT INTO public.transactions(t_time_stamp, t_net_price, t_image_url, order_id) VALUES ($1, $2, $3, $4) RETURNING id, t_time_stamp, t_net_price, t_image_url, order_id;"

    var transaction entities.Transaction

    result := tpr.db.Raw(query, t.T_time_stamp, t.T_net_price, t.T_image_url, t.OrderId).Scan(&transaction)

    if result.Error != nil {
        return entities.Transaction{}, result.Error
    }

    return transaction, nil
}

func (tpr *TransactionPostgresRepository) UpdateTransaction(id int, t entities.Transaction) (entities.Transaction, error) {
    query := "UPDATE public.transactions SET t_time_stamp=$2, t_net_price=$3, t_image_url=$4, order_id=$5 WHERE id = $1 RETURNING id, t_time_stamp, t_net_price, t_image_url, order_id;"

    var transaction entities.Transaction

    result := tpr.db.Raw(query, id, t.T_time_stamp, t.T_net_price, t.T_image_url, t.OrderId).Scan(&transaction)

    if result.Error != nil {
        return entities.Transaction{}, result.Error
    }

    return transaction, nil
}

func (tpr *TransactionPostgresRepository) GetTransactionById(id int) (entities.Transaction, error) {
    query := "SELECT id, t_time_stamp, t_net_price, t_image_url, order_id FROM public.transactions WHERE id = $1;"

    var transaction entities.Transaction

    result := tpr.db.Raw(query, id).Scan(&transaction)

    if result.Error != nil {
        return entities.Transaction{}, result.Error
    }

    return transaction, nil
}

func (tpr *TransactionPostgresRepository) GetAllTransactions() ([]entities.Transaction, error) {
    query := "SELECT id, t_time_stamp, t_net_price, t_image_url, order_id FROM public.transactions;"

    var transactions []entities.Transaction

    result := tpr.db.Raw(query).Scan(&transactions)

    if result.Error != nil {
        return nil, result.Error
    }

    return transactions, nil
}

func (tpr *TransactionPostgresRepository) GetTransactionByOrderId(orderId int) (entities.Transaction, error) {
    query := "SELECT id, t_time_stamp, t_net_price, t_image_url, order_id FROM public.transactions WHERE order_id = $1;"

    var transaction entities.Transaction

    result := tpr.db.Raw(query, orderId).Scan(&transaction)

    if result.Error != nil {
        return entities.Transaction{}, result.Error
    }

    return transaction, nil
}

func (tpr *TransactionPostgresRepository) DeleteTransaction(id int) error {
    query := "DELETE FROM public.transactions WHERE id = $1;"

    result := tpr.db.Exec(query, id)

    if result.Error != nil {
        return result.Error
    }    

    return nil
}

