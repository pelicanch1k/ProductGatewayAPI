package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pelicanch1k/ProductGatewayAPI/structs"
)

type Auth interface {
	CreateUser(user structs.User) (int, error)
	GetUserId(username, password_hash string) (structs.User, error)
}

type Products interface {
	Create(product structs.Product) (int, error)
	Update(product structs.Product, id int) error
	Delete(id int) error
	Get(id int) (structs.Product, error)
	GetAll() ([]structs.Product, error)
}

type Repository struct {
	Products
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Products: NewProductsPostgres(db),
		Auth:     NewAuthRepository(db),
	}
}
