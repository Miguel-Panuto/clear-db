package engine_parser

import (
	"errors"
	"regexp"
	"strings"

	enums_engine "github.com/miguel-panuto/clear-db/src/engine/enums"
)

type Command struct {
	Operation enums_engine.Operations
	Data      interface{}
}

func ParseString(statement string) (*Command, error) {
	parsedStatement := strings.ReplaceAll(statement, "\n", " ")
	if strings.HasPrefix(strings.ToLower(parsedStatement), "create database") {
		re := regexp.MustCompile(`(?i)create database`)
		dbName := re.ReplaceAllString(parsedStatement, "")
		return &Command{Operation: enums_engine.CREATE_DATABASE, Data: dbName}, nil
	}
	if strings.HasPrefix(strings.ToLower(parsedStatement), "list dbs") {
		return &Command{Operation: enums_engine.LIST_DATABASES}, nil
	}
	if strings.HasPrefix(strings.ToLower(parsedStatement), "use") {
		re := regexp.MustCompile(`(?i)create database`)
		dbName := re.ReplaceAllString(parsedStatement, "")
		return &Command{Operation: enums_engine.USE_DATABASE, Data: dbName}, nil
	}

	return nil, errors.New("command not found")
}
