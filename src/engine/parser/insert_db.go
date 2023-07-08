package engine_parser

import (
	"regexp"
	"strings"

	engine_enums "github.com/miguel-panuto/clear-db/src/engine/enums"
	engine_struct "github.com/miguel-panuto/clear-db/src/engine/struct"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func insertTable(parsedStatement string) (*Command, error) {
	re := regexp.MustCompile(`(?i)insert`)
	parsedStatement = re.ReplaceAllString(parsedStatement, "")
	splitedString := utils.Split(parsedStatement, ":")

	tableName := strings.TrimSpace(splitedString[0])
	rawRow := utils.SubString(splitedString[1], "{", "}")
	row := utils.TrimSplit(rawRow, ",")
	return &Command{
		Operation: engine_enums.INSERT_INTO,
		Data:      engine_struct.RowInsert{TabName: tableName, Row: row}}, nil
}
