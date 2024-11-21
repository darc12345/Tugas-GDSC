package controller

import (
	"database/sql"
	"main/model"
	"main/service"
)

type ControllerDB struct {
	serviceDB service.ServiceDB
}

func NewControllerDB(db *sql.DB) ControllerDB {
	return ControllerDB{serviceDB: service.NewServiceDB(db)}
}

type getBookStructure struct {
	Data model.Book `json:"data"`
}
type getBooksStructure struct {
	Data []model.Book `json:"data"`
}
type postBookStructure struct {
	Message string     `json:"message"`
	Data    model.Book `json:"data"`
}
type deleteBookStructure struct {
	Message string `json:"message"`
}
type messageStructure struct {
	Message string `json:"message"`
}
type putBookStructure struct {
	Message string     `json:"message"`
	Data    model.Book `json:"data"`
}
