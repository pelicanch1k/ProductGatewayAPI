package repository

import (
	"github.com/jmoiron/sqlx"
	product "github.com/pelicanch1k/rest-api"
)

type Products interface {
	Create(id int, name string, price float64) error
	Update(id int, name string, price float64) error
	Delete(id int) error
	Get(id int) product.Product
	GetAll() []product.Product
}

type Repository struct {
	Products
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Products: NewProductsPostgres(db),
	}
}
