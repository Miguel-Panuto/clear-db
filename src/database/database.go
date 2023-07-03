package database

import (
	"github.com/miguel-panuto/clear-db/src/database/table"
)

type Database struct {
	Name   string
	tables []table.Table
}

// func loadDatabases() {}

func NewDatabase(name string) *Database {
	database := Database{Name: name, tables: []table.Table{}}
	return &database
}
