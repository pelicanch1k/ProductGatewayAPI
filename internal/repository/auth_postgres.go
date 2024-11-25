package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pelicanch1k/ProductGatewayAPI/structs"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user structs.User) (int, error) {
	query := "INSERT INTO users (username, password_hash) values ($1, $2) RETURNING id"
	row := r.db.QueryRow(query, user.Username, user.Password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUserId(username, password_hash string) (structs.User, error) {
	var user structs.User
	query := "SELECT * FROM users WHERE username = $1 and password_hash = $2 LIMIT 1"
	err := r.db.Get(&user, query, username, password_hash)

	return user, err
}
