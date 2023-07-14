package database

import (
	"errors"
	"strconv"

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

func (db *Database) ListTables() [][]string {
	header := []string{"Name", "Rows"}
	rows := [][]string{header}
	for _, table := range db.Tables {
		rows = append(rows, []string{table.Name, strconv.Itoa(len(*table.Rows))})
	}
	return rows
}

func (db *Database) FindTable(tableName string) (*table.Table, error) {
	if !db.isThereAnyTable(tableName) {
		return nil, errors.New("no table founded")
	}

	var table table.Table
	for _, value := range db.Tables {
		if value.Name == tableName {
			table = value
			break
		}
	}
	return &table, nil
}

func (db *Database) GetTablesNumber() int {
	return len(db.Tables)
}
