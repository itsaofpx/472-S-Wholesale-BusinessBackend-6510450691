package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type SupplierPostgresRepository struct {
	db *gorm.DB
}

func InitiateSupplierPostgresRepository(db *gorm.DB) repositories.SupplierRepository {
	return &SupplierPostgresRepository{
		db: db,
	}
}

func (spr SupplierPostgresRepository) CreateSupplier(supplier entities.Supplier) (entities.Supplier, error) {
	query := `INSERT INTO public.suppliers(
				name, contract_person, email)
				VALUES ($1, $2, $3)
				RETURNING id, name, contract_person, email;`

	result := spr.db.Raw(query, supplier.Name, supplier.Contract_person, supplier.Email).Scan(&supplier)

	if result.Error != nil {
		return entities.Supplier{}, result.Error
	}

	return supplier, nil
}

func (spr SupplierPostgresRepository) UpdateSupplier(id int, supplier entities.Supplier) (entities.Supplier, error) {
	query := `UPDATE public.suppliers
				SET name = $1, contract_person = $2, email = $3
				WHERE id = $4
				RETURNING id, name, contract_person, email;`

	result := spr.db.Raw(query, supplier.Name, supplier.Contract_person, supplier.Email, id).Scan(&supplier)

	if result.Error != nil {
		return entities.Supplier{}, result.Error
	}

	return supplier, nil
}

func (spr SupplierPostgresRepository) GetSupplierByID(id int) (entities.Supplier, error) {
	query := `SELECT id, name, contract_person, email
				FROM public.suppliers
				WHERE id = $1;`

	var supplier entities.Supplier

	result := spr.db.Raw(query, id).Scan(&supplier)

	if result.Error != nil {
		return entities.Supplier{}, result.Error
	}

	return supplier, nil
}

func (spr SupplierPostgresRepository) GetAllSuppliers() ([]entities.Supplier, error) {
	query := `SELECT id, name, contract_person, email
				FROM public.suppliers;`

	var suppliers []entities.Supplier

	result := spr.db.Raw(query).Scan(&suppliers)

	if result.Error != nil {
		return nil, result.Error
	}

	return suppliers, nil
}
