package engine_parser

import (
	"errors"
	"regexp"
	"strings"

	engine_enums "github.com/miguel-panuto/clear-db/src/engine/enums"
)

type Command struct {
	Operation engine_enums.Operations
	Data      interface{}
}

func ParseString(statement string) (*Command, error) {
	parsedStatement := strings.ReplaceAll(statement, "\n", " ")
	lowerStatement := strings.ToLower(parsedStatement)

	if strings.HasPrefix(lowerStatement, "new db") {
		return newDbParse(parsedStatement)
	}

	if strings.HasPrefix(lowerStatement, "list dbs") {
		return &Command{Operation: engine_enums.LIST_DATABASES}, nil
	}

	if strings.HasPrefix(lowerStatement, "use") {
		re := regexp.MustCompile(`(?i)use`)
		dbName := re.ReplaceAllString(parsedStatement, "")
		return &Command{Operation: engine_enums.USE_DATABASE, Data: strings.TrimSpace(dbName)}, nil
	}

	if strings.HasPrefix(lowerStatement, "new table") {
		return newTable(parsedStatement)
	}

	if strings.TrimSpace(lowerStatement) == "list tables" {
		return &Command{Operation: engine_enums.LIST_TABLES}, nil
	}

	if strings.HasPrefix(strings.TrimSpace(lowerStatement), "insert") {
		return insertTable(parsedStatement)
	}

	if strings.TrimSpace(lowerStatement) == "exit" {
		return &Command{Operation: engine_enums.EXIT}, nil
	}

	return nil, errors.New("command not found")
}
