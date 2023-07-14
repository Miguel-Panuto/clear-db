package engine_parser

import (
	"errors"
	"regexp"
	"strings"

	engine_enums "github.com/miguel-panuto/clear-db/src/engine/enums"
	engine_struct "github.com/miguel-panuto/clear-db/src/engine/struct"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func findInTable(statement string) (*Command, error) {
	re := regexp.MustCompile(`(?i)in`)
	if !re.MatchString(statement) {
		return nil, errors.New("not found in operator")
	}

	re = regexp.MustCompile(`(?i)find`)
	statement = re.ReplaceAllString(statement, "")

	splited := utils.TrimSplit(statement, "in")

	if strings.Contains(splited[0], "{") != strings.Contains(splited[0], "}") {
		return nil, errors.New("pair of brackets not correctly closed {}")
	}

	if len(splited[1]) <= 1 {
		return nil, errors.New("table name was not inserted")
	}

	columns := []string{}
	if utils.ContainsMany(splited[0], "{", "}") {
		if !utils.IsBetween(splited[0], "{", "}") {
			return nil, errors.New("when has {} mas have to be closed before in")
		}
		columns = utils.SubSplit(splited[0], "{", "}", ",")
	}

	return &Command{
		Operation: engine_enums.FIND_IN,
		Data:      engine_struct.FindIn{TableName: splited[1], Columns: columns},
	}, nil
}
