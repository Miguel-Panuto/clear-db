package engine

import (
	"errors"
	"fmt"
	"os"

	"github.com/miguel-panuto/clear-db/src/domain"
	engine_enums "github.com/miguel-panuto/clear-db/src/engine/internal/enums"
	engine_io "github.com/miguel-panuto/clear-db/src/engine/internal/io"
	engine_parser "github.com/miguel-panuto/clear-db/src/engine/internal/parser"
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

func (e *Engine) RunStatement(statement string) (interface{}, error) {
	cmd, err := engine_parser.ParseString(statement)

	if err != nil {
		return nil, err
	}

	switch cmd.Operation {
	case engine_enums.CREATE_DATABASE:
		dbName, _ := cmd.Data.(string)

		return nil, e.createDatabase(dbName)

	case engine_enums.LIST_DATABASES:

		return e.listDatabases(), nil

	case engine_enums.USE_DATABASE:
		dbName, _ := cmd.Data.(string)

		if err := e.useDb(dbName); err != nil {
			return nil, errors.New("database not founded")
		}

		fmt.Println("Using database " + cmd.Data.(string))
		return nil, nil

	case engine_enums.CREATE_TABLE:
		tableCreation, _ := cmd.Data.(domain.TableCreation)

		if err := e.isSelectedDatabase(); err != nil {
			return nil, err
		}

		err := e.createTable(tableCreation)

		if err != nil {
			fmt.Println(err)
		}
		return nil, nil

	case engine_enums.LIST_TABLES:
		if err := e.isSelectedDatabase(); err != nil {
			return nil, err
		}
		tables := e.selectedDatabase.ListTables()
		return tables, nil

	case engine_enums.INSERT_INTO:
		rowInsert, _ := cmd.Data.(domain.RowInsert)

		if err := e.isSelectedDatabase(); err != nil {
			return nil, err
		}

		if err := e.insert(rowInsert); err != nil {
			return nil, err
		}
		return nil, nil

	case engine_enums.FIND_IN:
		if err := e.isSelectedDatabase(); err != nil {
			return nil, err
		}

		findIn, _ := cmd.Data.(domain.FindIn)

		return e.findIn(findIn)

	case engine_enums.EXIT:
		os.Exit(0)
		return nil, nil

	default:
		return nil, err
	}
}
