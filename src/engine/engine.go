package engine

import (
	"errors"

	"github.com/miguel-panuto/clear-db/src/database"
	engine_enums "github.com/miguel-panuto/clear-db/src/engine/enums"
	engine_parser "github.com/miguel-panuto/clear-db/src/engine/parser"
)

type engine struct {
	selectedDatabase database.Database
	databases        []database.Database
}

func NewEngine() *engine {
	return &engine{databases: []database.Database{}}
}

func (e *engine) RunStatement(statement string) error {
	cmd, err := engine_parser.ParseString(statement)

	if err != nil {
		return err
	}

	switch cmd.Operation {
	case engine_enums.CREATE_DATABASE:
		if _, ok := cmd.Data.(string); !ok {
			return errors.New("not entered valid name")
		}
		e.createDatabase(cmd.Data.(string))
		return nil

	case engine_enums.LIST_DATABASES:
		e.listDatabases()
		return nil

	case engine_enums.USE_DATABASE:
		if _, ok := cmd.Data.(string); !ok {
			return errors.New("not entered valid name")
		}
		if !e.foundDatabaseByName(cmd.Data.(string)) {
			return errors.New("database not founded")
		}
		return nil

	default:
		return errors.New("command not found")
	}
}
