package table

import (
	"errors"
	"strings"

	"github.com/miguel-panuto/clear-db/src/utils"
)

func indexOf(row []string, field string, originalIndex int) int {
	for i, el := range row {
		if !strings.Contains(el, "::") {
			return originalIndex
		}

		column := utils.TrimSplit(el, "::")[0]
		if column == field {
			return i
		}
	}

	return -1
}

func (t *Table) InsertNewRow(row []string) error {
	parsedRows := []interface{}{}
	for i, value := range t.Fields {
		index := indexOf(row, value.name, i)
		if index < 0 {
			return errors.New("column not finded")
		}

		parsedValue, err := getValueType(row[index], value.data_type)
		if err != nil {
			return err
		}
		parsedRows = append(parsedRows, parsedValue)
	}

	*t.Rows = append(*t.Rows, parsedRows)

	return nil
}

func (t *Table) InsertFromReader(row []interface{}) error {
	*t.Rows = append(*t.Rows, row)

	return nil
}
