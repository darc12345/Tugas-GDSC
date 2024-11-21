package repository

import "database/sql"

type DBrepo struct {
	db *sql.DB
}

func NewRepoDB(dbHandler *sql.DB) DBrepo {
	return DBrepo{db: dbHandler}
}
