package engine_parser

import (
	"errors"
	"regexp"
	"strings"

	engine_enums "github.com/miguel-panuto/clear-db/src/engine/enums"
	engine_struct "github.com/miguel-panuto/clear-db/src/engine/struct"
	"github.com/miguel-panuto/clear-db/src/utils"
)

type Command struct {
	Operation engine_enums.Operations
	Data      interface{}
}

func ParseString(statement string) (*Command, error) {
	parsedStatement := strings.ReplaceAll(statement, "\n", " ")
	lowerStatement := strings.ToLower(parsedStatement)

	if strings.HasPrefix(lowerStatement, "new db") {
		re := regexp.MustCompile(`(?i)new db`)
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

	if strings.HasPrefix(lowerStatement, "new table") {
		re := regexp.MustCompile(`(?i)new table`)
		parsedStatement = re.ReplaceAllString(parsedStatement, "")
		splitedString := utils.Split(parsedStatement, ":")
		tableName := splitedString[0]
		fields := utils.Split(splitedString[1], ",")
		parsedFields := []string{}
		for _, value := range fields {
			parsedFields = append(parsedFields, strings.TrimSpace(value))
		}
		return &Command{Operation: engine_enums.CREATE_TABLE, Data: engine_struct.TableCreation{DbName: tableName, Fields: parsedFields}}, nil
	}

	if strings.TrimSpace(lowerStatement) == "list tables" {
		return &Command{Operation: engine_enums.LIST_TABLES}, nil
	}

	if strings.TrimSpace(lowerStatement) == "exit" {
		return &Command{Operation: engine_enums.EXIT}, nil
	}

	return nil, errors.New("command not found")
}
