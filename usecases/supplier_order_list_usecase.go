package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type SupplierOrderListUsecase interface {
	CreateSupplierOrderList(sr entities.SupplierOrderList) (entities.SupplierOrderList, error)
	UpdateSupplierOrderList(id int, supplierOrderList entities.SupplierOrderList) (entities.SupplierOrderList, error)
	GetSupplierOrderListByID(id int) (entities.SupplierOrderList, error)
	GetSupplierOrderListsBySupplierID(id int) ([]entities.SupplierOrderList, error)
	GetAllSupplierOrderLists() ([]entities.SupplierOrderList, error)
}

type SupplierOrderListService struct {
	repo repositories.SupplierOrderListRepository
}

func InitiateSupplierOrderListService(repo repositories.SupplierOrderListRepository) SupplierOrderListUsecase {
	return &SupplierOrderListService{
		repo: repo,
	}
}

func (sos *SupplierOrderListService) CreateSupplierOrderList(sr entities.SupplierOrderList) (entities.SupplierOrderList, error) {
	createdSupplierOrderList, err := sos.repo.CreateSupplierOrderList(sr)

	if err != nil {
		return entities.SupplierOrderList{}, err
	}

	return createdSupplierOrderList, nil
}

func (sos *SupplierOrderListService) UpdateSupplierOrderList(id int, supplierOrderList entities.SupplierOrderList) (entities.SupplierOrderList, error) {
	updateSupplierOrderList, err := sos.repo.UpdateSupplierOrderList(id, supplierOrderList)

	if err != nil {
		return entities.SupplierOrderList{}, err
	}

	return updateSupplierOrderList, nil
}

func (sos *SupplierOrderListService) GetSupplierOrderListByID(id int) (entities.SupplierOrderList, error) {
	supplierOrderList, err := sos.repo.GetSupplierOrderListByID(id)

	if err != nil {
		return entities.SupplierOrderList{}, err
	}

	return supplierOrderList, nil
}

func (sos *SupplierOrderListService) GetSupplierOrderListsBySupplierID(id int) ([]entities.SupplierOrderList, error) {
	supplierOrderLists, err := sos.repo.GetSupplierOrderListsBySupplierID(id)

	if err != nil {
		return nil, err
	}

	return supplierOrderLists, nil
}

func (sos *SupplierOrderListService) GetAllSupplierOrderLists() ([]entities.SupplierOrderList, error) {
	supplierOrderLists, err := sos.repo.GetAllSupplierOrderLists()

	if err != nil {
		return nil, err
	}

	return supplierOrderLists, nil
}
