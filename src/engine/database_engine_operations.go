package engine

import (
	"errors"
	"fmt"
	"strings"

	"github.com/miguel-panuto/clear-db/src/database"
	engine_io "github.com/miguel-panuto/clear-db/src/engine/io"
	engine_utils "github.com/miguel-panuto/clear-db/src/engine/utils"
)

func (e *Engine) listDatabases() {
	header := []string{"Name"}
	var rows [][]string
	for _, db := range e.databases {
		rows = append(rows, []string{db})
	}
	engine_utils.PrintTable(header, rows)
}

func (e *Engine) createDatabase(dbName string) error {
	found := e.foundDatabaseByName(dbName)
	if len(dbName) < 2 {
		return errors.New("database name length not accepted")
	}

	if found {
		return errors.New("database already created")
	}

	db := database.NewDatabase(strings.TrimSpace(dbName))
	e.databases = append(e.databases, db.Name)
	fmt.Printf("Database created %s \n", db.Name)
	go engine_io.SaveData(db)
	return nil
}

func (e *Engine) useDb(name string) error {
	if !e.foundDatabaseByName(name) {
		return errors.New("database not finded")
	}

	db, err := engine_io.ReadDatabase(name)

	if err != nil {
		return err
	}

	e.selectedDatabase = db
	return nil
}

func (e *Engine) foundDatabaseByName(name string) bool {
	found := false
	for _, db := range e.databases {
		if db == name {
			found = true
			break
		}
	}

	return found
}
