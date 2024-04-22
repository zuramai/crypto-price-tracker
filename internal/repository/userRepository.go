package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func (repo *UserRepository) FindAll() {
}
