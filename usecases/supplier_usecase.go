package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type SupplierUsecase interface {
	CreateSupplier(s entities.Supplier) (entities.Supplier, error)
	UpdateSupplier(id int, s entities.Supplier) (entities.Supplier, error)
	GetSupplierByID(id int) (entities.Supplier, error)
	GetAllSuppliers() ([]entities.Supplier, error)
}

type SupplierService struct {
	repo repositories.SupplierRepository
}

func InitiateSupplierService(repo repositories.SupplierRepository) SupplierUsecase {
	return &SupplierService{
		repo: repo,
	}
}

func (ss *SupplierService) CreateSupplier(s entities.Supplier) (entities.Supplier, error) {
	cs, err := ss.repo.CreateSupplier(s)

	if err != nil {
		return entities.Supplier{}, err
	}

	return cs, nil
}

func (ss *SupplierService) UpdateSupplier(id int, s entities.Supplier) (entities.Supplier, error) {
	updateSupplier, err := ss.repo.UpdateSupplier(id, s)

	if err != nil {
		return entities.Supplier{}, err
	}

	return updateSupplier, nil
}

func (ss *SupplierService) GetSupplierByID(id int) (entities.Supplier, error) {
	supplier, err := ss.repo.GetSupplierByID(id)

	if err != nil {
		return entities.Supplier{}, err
	}

	return supplier, nil
}

func (ss *SupplierService) GetAllSuppliers() ([]entities.Supplier, error) {
	supplierList, err := ss.repo.GetAllSuppliers()

	if err != nil {
		return nil, err
	}

	return supplierList, nil
}
