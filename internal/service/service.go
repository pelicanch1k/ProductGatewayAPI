package service

import (
	product "github.com/pelicanch1k/rest-api"
	"github.com/pelicanch1k/rest-api/internal/repository"
)

type Products interface {
	Create(id int, name string, price float64) error
	Update(id int, name string, price float64) error
	Delete(id int) error
	Get(id int) product.Product
	GetAll() []product.Product
}

type Service struct {
	Products
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
