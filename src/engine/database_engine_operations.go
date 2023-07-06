package engine

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/miguel-panuto/clear-db/src/database"
	engine_io "github.com/miguel-panuto/clear-db/src/engine/io"
	engine_utils "github.com/miguel-panuto/clear-db/src/engine/utils"
)

func (e *Engine) listDatabases() {
	header := []string{"Name", "Tables"}
	var rows [][]string
	for _, db := range e.databases {
		rows = append(rows, []string{db.Name, strconv.Itoa(db.GetTablesNumber())})
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
	e.databases = append(e.databases, db)
	fmt.Printf("Database created %s \n", db.Name)
	go engine_io.SaveData(db)
	return nil
}

func (e *Engine) foundDatabaseByName(name string) bool {
	found := false
	for _, db := range e.databases {
		if db.Name == name {
			e.selectedDatabase = db
			found = true
			break
		}
	}

	return found
}
