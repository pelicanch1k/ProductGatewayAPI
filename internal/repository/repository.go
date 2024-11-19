package repository

import (
	"github.com/jmoiron/sqlx"
	product "github.com/pelicanch1k/rest-api"
)

type Products interface {
	Create(product product.Product) (int, error)
	Update(product product.Product, id int) error
	Delete(id int) error
	Get(id int) (product.Product, error)
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
