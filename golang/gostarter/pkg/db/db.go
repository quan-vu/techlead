package db

import (
	"database/sql"
)

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	return nil, nil
}

func (d *DB) Close() error {
	return nil
}
