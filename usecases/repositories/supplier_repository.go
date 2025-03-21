package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type SupplierRepository interface {
	CreateSupplier(s entities.Supplier) (entities.Supplier, error)
	UpdateSupplier(id int, s entities.Supplier) (entities.Supplier, error)
	GetSupplierByID(id int) (entities.Supplier, error)
	GetAllSuppliers() ([]entities.Supplier, error)
}
