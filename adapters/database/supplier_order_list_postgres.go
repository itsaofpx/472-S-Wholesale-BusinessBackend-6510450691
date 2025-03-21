package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type SupplierOrderListPostgresRepository struct {
	db *gorm.DB
}

func InitiateSupplierOrderListPostgresRepository(db *gorm.DB) repositories.SupplierOrderListRepository {
	return &SupplierOrderListPostgresRepository{
		db: db,
	}
}

func (s SupplierOrderListPostgresRepository) CreateSupplierOrderList(supplierOrderList entities.SupplierOrderList) (entities.SupplierOrderList, error) {
	query := `INSERT INTO public.supplier_order_lists(
					supplier_id, product_id, quantity, price)
					VALUES ($1, $2, $3, $4)
					RETURNING id, supplier_id, product_id, quantity, price;`

	var createdSupplierOrderList entities.SupplierOrderList

	result := s.db.Raw(query, supplierOrderList.SupplierID, supplierOrderList.ProductID, supplierOrderList.Quantity, supplierOrderList.Price).Scan(&createdSupplierOrderList)

	if result.Error != nil {
		return entities.SupplierOrderList{}, result.Error
	}

	return createdSupplierOrderList, nil
}

func (s SupplierOrderListPostgresRepository) UpdateSupplierOrderList(id int, supplierOrderList entities.SupplierOrderList) (entities.SupplierOrderList, error) {
	query := `UPDATE public.supplier_order_lists
					SET supplier_id = $2, product_id = $3, quantity = $4, price = $5
					WHERE id = $1
					RETURNING id, supplier_id, product_id, quantity, price;`

	var updatedSupplierOrderList entities.SupplierOrderList

	result := s.db.Raw(query, id, supplierOrderList.SupplierID, supplierOrderList.ProductID, supplierOrderList.Quantity, supplierOrderList.Price).Scan(&updatedSupplierOrderList)

	if result.Error != nil {
		return entities.SupplierOrderList{}, result.Error
	}

	return updatedSupplierOrderList, nil
}

func (s SupplierOrderListPostgresRepository) GetSupplierOrderListsBySupplierID(id int) ([]entities.SupplierOrderList, error) {
	query := `SELECT id, supplier_id, product_id, quantity, price
					FROM public.supplier_order_lists
					WHERE supplier_id = $1;`

	var supplierOrderLists []entities.SupplierOrderList

	result := s.db.Raw(query, id).Scan(&supplierOrderLists)

	if result.Error != nil {
		return nil, result.Error
	}

	return supplierOrderLists, nil
}

func (s SupplierOrderListPostgresRepository) GetSupplierOrderListByID(id int) (entities.SupplierOrderList, error) {
	query := `SELECT id, supplier_id, product_id, quantity, price
					FROM public.supplier_order_lists
					WHERE id = $1;`

	var supplierOrderList entities.SupplierOrderList

	result := s.db.Raw(query, id).Scan(&supplierOrderList)

	if result.Error != nil {
		return entities.SupplierOrderList{}, result.Error
	}

	return supplierOrderList, nil
}

func (s SupplierOrderListPostgresRepository) GetAllSupplierOrderLists() ([]entities.SupplierOrderList, error) {
	query := `SELECT id, supplier_id, product_id, quantity, price
					FROM public.supplier_order_lists;`

	var supplierOrderLists []entities.SupplierOrderList

	result := s.db.Raw(query).Scan(&supplierOrderLists)

	if result.Error != nil {
		return nil, result.Error
	}

	return supplierOrderLists, nil
}
