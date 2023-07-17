package engine_parser

import (
	"errors"
	"regexp"
	"strings"

	domain "github.com/miguel-panuto/clear-db/src/domain/struct"
	engine_enums "github.com/miguel-panuto/clear-db/src/engine/internal/enums"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func insertTable(parsedStatement string) (*Command, error) {
	re := regexp.MustCompile(`(?i)insert`)
	parsedStatement = re.ReplaceAllString(parsedStatement, "")

	splitedString := utils.SplitFirst(parsedStatement, ":")

	tableName := strings.TrimSpace(splitedString[0])

	if strings.Contains(splitedString[1], "::") != utils.ContainsMany(splitedString[1], "{", "}") {
		return nil, errors.New("not valid statement for parsing: :: and {} must appear together")
	}

	rawRow := splitedString[1]
	if utils.ContainsMany(splitedString[1], "{", "}") {
		rawRow = utils.SubString(rawRow, "{", "}")
	}
	row := utils.TrimSplit(rawRow, ",")
	return &Command{
		Operation: engine_enums.INSERT_INTO,
		Data:      domain.RowInsert{TableName: tableName, Row: row}}, nil
}
