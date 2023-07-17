package engine_parser

import (
	"regexp"
	"strings"

	domain "github.com/miguel-panuto/clear-db/src/domain/struct"
	engine_enums "github.com/miguel-panuto/clear-db/src/engine/internal/enums"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func newTable(parsedStatement string) (*Command, error) {
	re := regexp.MustCompile(`(?i)new table`)
	parsedStatement = re.ReplaceAllString(parsedStatement, "")
	splitedString := utils.Split(parsedStatement, ":")
	tableName := splitedString[0]
	fields := utils.Split(splitedString[1], ",")
	parsedFields := []string{}
	for _, value := range fields {
		parsedFields = append(parsedFields, strings.TrimSpace(value))
	}
	return &Command{Operation: engine_enums.CREATE_TABLE, Data: domain.TableCreation{DbName: tableName, Fields: parsedFields}}, nil
}
