package service

import (
	"database/sql"
	"main/repository"
)

type ServiceDB struct {
	RepoDB repository.DBrepo
}

func NewServiceDB(db *sql.DB) ServiceDB {
	return ServiceDB{RepoDB: repository.NewRepoDB(db)}
}
