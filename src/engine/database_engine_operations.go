package engine

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/miguel-panuto/clear-db/src/database"
	engine_io "github.com/miguel-panuto/clear-db/src/engine/io"
)

func (e *Engine) listDatabases() {
	const baseLenName int = 4
	const baseLenQty int = 6
	maxLenName := baseLenName
	maxLenQty := baseLenQty

	for _, db := range e.databases {
		if len(db.Name) > maxLenName {
			maxLenName = len(db.Name)
		}
		stringQty := len(strconv.Itoa(db.GetTablesNumber()))
		if stringQty > maxLenQty {
			maxLenQty = stringQty
		}
	}

	blankSpacesName := strings.Repeat(" ", maxLenName-baseLenName)
	blankSpacesQty := strings.Repeat(" ", maxLenQty-baseLenQty)
	printStatement := "| Name " + blankSpacesName + "| Tables " + blankSpacesQty + "|"
	printStatement = strings.ReplaceAll(printStatement, "{blankSpacesQty}", blankSpacesQty)
	fmt.Println(printStatement)
	for _, value := range e.databases {
		blankSpacesName = strings.Repeat(" ", maxLenName-len(value.Name))
		blankSpacesQty = strings.Repeat(" ", maxLenQty-len(strconv.Itoa(value.GetTablesNumber())))
		printStatement =
			"| " +
				value.Name +
				blankSpacesName +
				" | " +
				strconv.Itoa(value.GetTablesNumber()) +
				blankSpacesQty +
				" |"

		fmt.Println(printStatement)
	}
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
