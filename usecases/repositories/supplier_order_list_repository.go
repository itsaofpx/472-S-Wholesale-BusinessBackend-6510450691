package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type SupplierOrderListRepository interface {
	CreateSupplierOrderList(sr entities.SupplierOrderList) (entities.SupplierOrderList, error)
	UpdateSupplierOrderList(id int, supplierOrderList entities.SupplierOrderList) (entities.SupplierOrderList, error)
	GetSupplierOrderListByID(id int) (entities.SupplierOrderList, error)
	GetSupplierOrderListsBySupplierID(id int) ([]entities.SupplierOrderList, error)
	GetAllSupplierOrderLists() ([]entities.SupplierOrderList, error)
}
