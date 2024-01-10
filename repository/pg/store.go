package pg

import (
	"database/sql"

	rep "main/repository"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgreSqlStore struct {
	db *sql.DB
}

func NewPostgreSqlStore(db *sql.DB) rep.DataStoreInterface {
	return &PostgreSqlStore{
		db: db,
	}
}
