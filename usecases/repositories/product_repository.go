package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type ProductRepository interface {
	CreateProduct(p entities.Product) error
	GetProductByID(id int) (entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
}
