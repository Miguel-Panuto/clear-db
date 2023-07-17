package engine

import (
	"fmt"

	"github.com/miguel-panuto/clear-db/src/domain"
	engine_io "github.com/miguel-panuto/clear-db/src/engine/internal/io"
)

func (e *Engine) createTable(data domain.TableCreation) error {
	err := e.selectedDatabase.NewTable(data.DbName, data.Fields)
	if err != nil {
		return err
	}
	go engine_io.UpdateFile(e.selectedDatabase)
	fmt.Println("Table created " + data.DbName)
	return nil
}

func (e *Engine) insert(data domain.RowInsert) error {
	table, err := e.selectedDatabase.FindTable(data.TableName)

	if err != nil {
		return err
	}

	if err := table.InsertNewRow(data.Row); err != nil {
		return err
	}
	fmt.Println("new row was inserted")
	go engine_io.UpdateFile(e.selectedDatabase)

	return nil
}

func (e *Engine) findIn(data domain.FindIn) ([][]string, error) {
	table, err := e.selectedDatabase.FindTable(data.TableName)

	if err != nil {
		return [][]string{}, err
	}

	findData, err := table.FindIn(data.Columns, data.Where)

	if err != nil {
		return [][]string{}, err
	}

	return findData, nil
}
