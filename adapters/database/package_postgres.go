package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type PackagePostgresRepository struct {
	db *gorm.DB
}

func InitiatePackagePostgresRepository(db *gorm.DB) repositories.PackageRepository {
	return &PackagePostgresRepository{
		db: db,
	}
}

func (ppr *PackagePostgresRepository) CreatePackage(p entities.Package) (entities.Package, error) {
	query := "INSERT INTO public.packages(shipment_id, product_id, quantity, time_stamp, tracking_no) VALUES ($1, $2, $3, $4, $5) RETURNING id, shipment_id, product_id, quantity, time_stamp;"

	var i_package entities.Package

	result := ppr.db.Raw(query, p.ShipmentID, p.ProductID, p.Quantity, p.Time_stamp, p.TrackingNo).Scan(&i_package)

	if result.Error != nil {
		return entities.Package{}, result.Error
	}

	return i_package, nil
}

func (ppr *PackagePostgresRepository) GetPackageByID(id int) (entities.Package, error) {
	query := "SELECT id, shipment_id, product_id, quantity, time_stamp, tracking_no FROM public.packages WHERE id = $1;"

	var i_package entities.Package

	result := ppr.db.Raw(query, id).Scan(&i_package)

	if result.Error != nil {
		return entities.Package{}, result.Error
	}

	return i_package, nil
}

func (ppr *PackagePostgresRepository) GetAllPackagesByShipmentID(id int) ([]entities.Package, error) {
	query := "SELECT id, shipment_id, product_id, quantity, time_stamp, tracking_no FROM public.packages WHERE shipment_id = $1;"

	var packages []entities.Package

	result := ppr.db.Raw(query, id).Scan(&packages)

	if result.Error != nil {
		return nil, result.Error
	}

	return packages, nil
}

func (ppr *PackagePostgresRepository) GetAllPackages() ([]entities.Package, error) {
	query := "SELECT id, shipment_id, product_id, quantity, time_stamp, tracking_no FROM public.packages;"

	var packages []entities.Package

	result := ppr.db.Raw(query).Scan(&packages)

	if result.Error != nil {
		return nil, result.Error
	}

	return packages, nil
}
