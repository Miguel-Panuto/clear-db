package database

import (
	"errors"

	"github.com/miguel-panuto/clear-db/src/database/table"
)

func (db *Database) isThereAnyTable(name string) bool {
	for _, table := range db.tables {
		if table.Name == name {
			return true
		}
	}
	return false
}

func (db *Database) NewTable(name string, colums []string) error {
	if db.isThereAnyTable(name) {
		return errors.New("table already created")
	}

	newTable, err := table.NewTable(name, colums)
	if err != nil {
		return err
	}

	db.tables = append(db.tables, *newTable)
	return nil
}

func (*Database) ListTables(name string, colums []string) {

}

func (e *Database) GetTablesNumber() int {
	return len(e.tables)
}
