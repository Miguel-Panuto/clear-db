package database

import (
	"github.com/miguel-panuto/clear-db/src/database/table"
)

func NewDatabase(name string) *Database {
	database := Database{Name: name, Tables: []table.Table{}}
	return &database
}
