package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func createProduct(db *sqlx.DB) {
	scheme := `
CREATE TABLE IF NOT EXISTS products (
   id serial PRIMARY KEY not null unique,
   name varchar(255) not null unique,
   price integer not null default 0
)
`

	db.MustExec(scheme)
}

func createUser(db *sqlx.DB) {
	scheme := `
create table if not exists users (
    id serial PRIMARY KEY not null unique,
    username varchar(20) not null unique,
    password_hash varchar(255) not null
)
`
	db.MustExec(scheme)
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	createProduct(db)
	createUser(db)

	return db, nil
}
