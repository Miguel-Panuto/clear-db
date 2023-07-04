package engine_parser

import (
	"errors"
	"regexp"
	"strings"

	engine_enums "github.com/miguel-panuto/clear-db/src/engine/enums"
	engine_struct "github.com/miguel-panuto/clear-db/src/engine/struct"
)

type Command struct {
	Operation engine_enums.Operations
	Data      interface{}
}

func ParseString(statement string) (*Command, error) {
	parsedStatement := strings.ReplaceAll(statement, "\n", " ")
	lowerStatement := strings.ToLower(parsedStatement)

	if strings.HasPrefix(lowerStatement, "create database") {
		re := regexp.MustCompile(`(?i)create database`)
		dbName := re.ReplaceAllString(parsedStatement, "")
		return &Command{Operation: engine_enums.CREATE_DATABASE, Data: strings.TrimSpace(dbName)}, nil
	}

	if strings.HasPrefix(lowerStatement, "list dbs") {
		return &Command{Operation: engine_enums.LIST_DATABASES}, nil
	}

	if strings.HasPrefix(lowerStatement, "use") {
		re := regexp.MustCompile(`(?i)use`)
		dbName := re.ReplaceAllString(parsedStatement, "")
		return &Command{Operation: engine_enums.USE_DATABASE, Data: strings.TrimSpace(dbName)}, nil
	}

	if strings.HasPrefix(lowerStatement, "create table") {
		re := regexp.MustCompile(`(?i)create table`)
		parsedStatement = re.ReplaceAllString(parsedStatement, "")
		splitedString := strings.Split(parsedStatement, "(")
		dbName := splitedString[0]
		fields := strings.Split(splitedString[1], ",")
		parsedFields := []string{}
		for i, value := range fields {
			if strings.Contains(value, ")") {
				parsedFields = append(parsedFields, strings.TrimSpace(strings.Replace(value, ")", "", 1)))
				break
			}

			if i+1 == len(fields) {
				return nil, errors.New("error on trying to parse columns name, table not closed with )")
			}
			parsedFields = append(parsedFields, strings.TrimSpace(value))
		}
		return &Command{Operation: engine_enums.CREATE_TABLE, Data: engine_struct.TableCreation{DbName: dbName, Fields: parsedFields}}, nil
	}

	if lowerStatement == "exit" {
		return &Command{Operation: engine_enums.EXIT}, nil
	}

	return nil, errors.New("command not found")
}
