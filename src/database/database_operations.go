package database

import (
	"errors"
	"fmt"

	"github.com/miguel-panuto/clear-db/src/database/table"
)

func (db *Database) isThereAnyTable(name string) bool {
	for _, table := range db.Tables {
		if table.Name == name {
			return true
		}
	}
	return false
}

func (db *Database) NewTable(name string, columns []string) error {
	if db.isThereAnyTable(name) {
		return errors.New("table already created")
	}

	newTable, err := table.NewTable(name, columns)
	if err != nil {
		return err
	}

	db.Tables = append(db.Tables, *newTable)
	return nil
}

func (db *Database) ListTables() {
	fmt.Println(db.Tables)
}

func (db *Database) GetTablesNumber() int {
	return len(db.Tables)
}
