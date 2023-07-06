package database

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

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
	const baseLenName int = 4
	const baseLenQty int = 4

	maxLenName := baseLenName
	maxLenQty := baseLenQty
	for _, table := range db.Tables {
		if len(table.Name) > maxLenName {
			maxLenName = len(table.Name)
		}

		stringQty := len(strconv.Itoa(len(table.Rows)))
		if stringQty > maxLenQty {
			maxLenQty = stringQty
		}
	}

	blankSpacesName := strings.Repeat(" ", maxLenName-baseLenName)
	blankSpacesQty := strings.Repeat(" ", maxLenQty-baseLenQty)

	printStatement := "| Name " + blankSpacesName + "| Rows " + blankSpacesQty + "|"
	fmt.Println(printStatement)
	for _, value := range db.Tables {
		blankSpacesName = strings.Repeat(" ", maxLenName-len(value.Name))
		blankSpacesQty = strings.Repeat(" ", maxLenQty-len(strconv.Itoa(len(value.Rows))))
		printStatement =
			"| " +
				value.Name +
				blankSpacesName +
				" | " +
				strconv.Itoa(len(value.Rows)) +
				blankSpacesQty +
				" |"

		fmt.Println(printStatement)
	}
}

func (db *Database) GetTablesNumber() int {
	return len(db.Tables)
}
