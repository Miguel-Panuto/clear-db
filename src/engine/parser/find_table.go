package engine_parser

import (
	"errors"
	"regexp"
	"strings"

	engine_enums "github.com/miguel-panuto/clear-db/src/engine/enums"
	engine_struct "github.com/miguel-panuto/clear-db/src/engine/struct"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func verifyContainsBraces(s string) bool {
	return utils.IsBetween(s, "{", "}")
}

func findInTable(statement string) (*Command, error) {
	re := regexp.MustCompile(`(?i)in`)
	if !re.MatchString(statement) {
		return nil, errors.New("not found in operator")
	}

	re = regexp.MustCompile(`(?i)find`)
	statement = re.ReplaceAllString(statement, "")

	splited := utils.MultipleSplit(statement, "in", "where")

	if strings.Contains(splited[0], "{") != strings.Contains(splited[0], "}") {
		return nil, errors.New("pair of brackets not correctly closed {}")
	}

	if len(splited[1]) <= 1 {
		return nil, errors.New("table name was not inserted")
	}

	columns := []string{}
	where := []engine_struct.Where{}
	if utils.ContainsMany(splited[0], "{", "}") {
		if !verifyContainsBraces(splited[0]) {
			return nil, errors.New("when has {} mas have to be closed before in")
		}
		columns = utils.SubSplit(splited[0], "{", "}", ",")
	}

	if len(splited) >= 3 {
		if !verifyContainsBraces(splited[2]) {
			return nil, errors.New("when has {} mas have to be closed after where")
		}
		rawWhere := utils.SubSplit(splited[2], "{", "}", ",")

		// for _, value := range rawWhere {

		// }
	}

	return &Command{
		Operation: engine_enums.FIND_IN,
		Data:      engine_struct.FindIn{TableName: splited[1], Columns: columns, Where: where},
	}, nil
}
