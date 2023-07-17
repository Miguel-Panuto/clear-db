package engine_parser

import (
	"errors"
	"regexp"
	"strings"

	domain "github.com/miguel-panuto/clear-db/src/domain/struct"
	engine_enums "github.com/miguel-panuto/clear-db/src/engine/internal/enums"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func verifyContainsBraces(s string) bool {
	s = strings.TrimSpace(s)
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
	where := []domain.Where{}
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

		for _, value := range rawWhere {
			splitedWhere := utils.SplitByOperators(value, "=", ">", ">=", "<=", "contains", "stw", "edw")
			if len(splitedWhere) < 3 {
				return nil, errors.New("failed to convert to where statement")
			}
			where = append(where, domain.Where{
				Column:   splitedWhere[0],
				Operator: splitedWhere[1],
				Value:    splitedWhere[2],
			})
		}
	}

	return &Command{
		Operation: engine_enums.FIND_IN,
		Data:      domain.FindIn{TableName: splited[1], Columns: columns, Where: where},
	}, nil
}
