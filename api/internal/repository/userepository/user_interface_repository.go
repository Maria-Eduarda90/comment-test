package userepository

import (
	"api/internal/database/sqlc"
	"database/sql"
)

func NewUserRepository(db *sql.DB, q *sqlc.Queries) UserRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct{
	db		*sql.DB
	queries *sqlc.Queries
}

type UserRepository interface {
	CreateUser() error
}