package engine

import (
	"fmt"

	engine_io "github.com/miguel-panuto/clear-db/src/engine/io"
	engine_struct "github.com/miguel-panuto/clear-db/src/engine/struct"
	engine_utils "github.com/miguel-panuto/clear-db/src/engine/utils"
)

func (e *Engine) createTable(data engine_struct.TableCreation) error {
	err := e.selectedDatabase.NewTable(data.DbName, data.Fields)
	if err != nil {
		return err
	}
	go engine_io.UpdateFile(e.selectedDatabase)
	fmt.Println("Table created " + data.DbName)
	return nil
}

func (e *Engine) listTables() {
	e.selectedDatabase.ListTables()
}

func (e *Engine) insert(data engine_struct.RowInsert) error {
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

func (e *Engine) findIn(data engine_struct.FindIn) error {
	table, err := e.selectedDatabase.FindTable(data.TableName)

	if err != nil {
		return err
	}

	findData, err := table.FindIn(data.Columns)

	if err != nil {
		return err
	}

	header := findData[0]
	findData = findData[1:]
	engine_utils.PrintTable(header, findData)
	return nil
}
