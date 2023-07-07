package engine_parser

import (
	"errors"
	"regexp"
	"strings"

	engine_enums "github.com/miguel-panuto/clear-db/src/engine/enums"
	"github.com/miguel-panuto/clear-db/src/utils"
)

type Command struct {
	Operation engine_enums.Operations
	Data      interface{}
}

func ParseString(statement string) (*Command, error) {
	parsedStatement := strings.ReplaceAll(statement, "\n", " ")

	if utils.VerifyLowerPrefix(parsedStatement, "new db") {
		return newDbParse(parsedStatement)
	}

	if utils.VerifyLower(parsedStatement, "list dbs") {
		return &Command{Operation: engine_enums.LIST_DATABASES}, nil
	}

	if utils.VerifyLowerPrefix(parsedStatement, "use") {
		re := regexp.MustCompile(`(?i)use`)
		dbName := re.ReplaceAllString(parsedStatement, "")
		return &Command{Operation: engine_enums.USE_DATABASE, Data: strings.TrimSpace(dbName)}, nil
	}

	if utils.VerifyLowerPrefix(parsedStatement, "new table") {
		return newTable(parsedStatement)
	}

	if utils.VerifyLower(parsedStatement, "list tables") {
		return &Command{Operation: engine_enums.LIST_TABLES}, nil
	}

	if utils.VerifyLowerPrefix(parsedStatement, "insert") {
		return insertTable(parsedStatement)
	}

	if utils.VerifyLower(parsedStatement, "exit") {
		return &Command{Operation: engine_enums.EXIT}, nil
	}

	return nil, errors.New("command not found")
}
