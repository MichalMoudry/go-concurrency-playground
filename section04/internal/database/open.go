package database

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func OpenDb(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
