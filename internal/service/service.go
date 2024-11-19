package service

import (
	product "github.com/pelicanch1k/rest-api"
	"github.com/pelicanch1k/rest-api/internal/repository"
)

type Products interface {
	Create(product product.Product) (int, error)
	Update(product product.Product, id int) error
	Delete(id int) error
	Get(id int) (product.Product, error)
	GetAll() []product.Product
}

type Service struct {
	Products
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Products: NewProductsService(repo),
	}
}
