package engine

import (
	"errors"
	"fmt"
	"os"

	engine_enums "github.com/miguel-panuto/clear-db/src/engine/enums"
	engine_io "github.com/miguel-panuto/clear-db/src/engine/io"
	engine_parser "github.com/miguel-panuto/clear-db/src/engine/parser"
	engine_struct "github.com/miguel-panuto/clear-db/src/engine/struct"
)

func NewEngine() *Engine {
	dbs, _ := engine_io.LoadDatabases()
	if dbs != nil {
		return &Engine{databases: dbs}
	}
	return &Engine{databases: []string{}}
}

func (e *Engine) isSelectedDatabase() error {
	if e.selectedDatabase == nil {
		return errors.New("there is none database beeing in use")
	}
	return nil
}

func (e *Engine) RunStatement(statement string) error {
	cmd, err := engine_parser.ParseString(statement)

	if err != nil {
		return err
	}

	switch cmd.Operation {
	case engine_enums.CREATE_DATABASE:
		dbName, _ := cmd.Data.(string)

		return e.createDatabase(dbName)

	case engine_enums.LIST_DATABASES:
		e.listDatabases()
		return nil

	case engine_enums.USE_DATABASE:
		dbName, _ := cmd.Data.(string)

		if err := e.useDb(dbName); err != nil {
			return errors.New("database not founded")
		}

		fmt.Println("Using database " + cmd.Data.(string))
		return nil

	case engine_enums.CREATE_TABLE:
		tableCreation, _ := cmd.Data.(engine_struct.TableCreation)

		if err := e.isSelectedDatabase(); err != nil {
			return err
		}

		err := e.createTable(tableCreation)

		if err != nil {
			fmt.Println(err)
		}
		return nil

	case engine_enums.LIST_TABLES:
		e.listTables()
		return nil

	case engine_enums.INSERT_INTO:
		rowInsert, _ := cmd.Data.(engine_struct.RowInsert)

		if err := e.isSelectedDatabase(); err != nil {
			return err
		}

		if err := e.insert(rowInsert); err != nil {
			return err
		}
		return nil

	case engine_enums.FIND_IN:
		if err := e.isSelectedDatabase(); err != nil {
			return err
		}

		findIn, _ := cmd.Data.(engine_struct.FindIn)

		if err := e.findIn(findIn); err != nil {
			return err
		}

		return nil

	case engine_enums.EXIT:
		os.Exit(0)
		return nil

	default:
		return err
	}
}
