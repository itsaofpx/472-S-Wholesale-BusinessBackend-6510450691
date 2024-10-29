package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type ProductRepository interface {
	CreateProduct(p entities.Product) (entities.Product, error)
	GetProductByID(id int) (entities.Product, error)
	GetProductByFilter(name string, minprice float64, maxprice float64) ([]entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
	UpdateProduct(id int, p entities.Product) (entities.Product, error)
}
