package engine

import (
	engine_io "github.com/miguel-panuto/clear-db/src/engine/io"
	engine_struct "github.com/miguel-panuto/clear-db/src/engine/struct"
)

func (e *Engine) createTable(data engine_struct.TableCreation) error {
	err := e.selectedDatabase.NewTable(data.DbName, data.Fields)
	if err != nil {
		return err
	}
	go engine_io.UpdateFile(e.selectedDatabase)
	return nil
}

func (e *Engine) listTables() {
	e.selectedDatabase.ListTables()
}
