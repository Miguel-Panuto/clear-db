package database

import "github.com/miguel-panuto/clear-db/src/database/table"

type Database struct {
	Name   string
	Tables []table.Table
}
