package engine_parser

import (
	"regexp"
	"strings"

	engine_enums "github.com/miguel-panuto/clear-db/src/engine/internal/enums"
)

func newDbParse(parsedStatement string) (*Command, error) {
	re := regexp.MustCompile(`(?i)new db`)
	dbName := re.ReplaceAllString(parsedStatement, "")
	return &Command{Operation: engine_enums.CREATE_DATABASE, Data: strings.TrimSpace(dbName)}, nil
}
