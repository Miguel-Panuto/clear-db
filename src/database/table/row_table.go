package table

import (
	"errors"
	"strings"

	"github.com/miguel-panuto/clear-db/src/domain"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func (f *Field) validateAndReturn(row []string, originalIndex int) (string, error) {
	for _, el := range row {
		if !strings.Contains(el, "::") {
			return row[originalIndex], nil
		}
		values := utils.TrimSplit(el, "::")
		if values[0] == f.name {
			return values[1], nil
		}

		if utils.ContainsInside(f.properties, "required") {
			return "", errors.New("required field must be filled")
		}
	}

	return "null", nil
}

func (t *Table) findValue(i int, value interface{}) bool {
	for _, row := range *t.Rows {
		if row[i] == value {
			return true
		}
	}
	return false
}

func (t *Table) InsertNewRow(row []string) error {
	parsedRows := []interface{}{}
	for i, f := range t.Fields {
		el, err := f.validateAndReturn(row, i)
		if err != nil {
			return err
		}

		parsedValue, err := getValueType(el, f.data_type)
		if utils.ContainsInside(f.properties, "unique") {
			if t.findValue(i, parsedValue) {
				return errors.New("unique value must be unique")
			}
		}
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

func (t *Table) FindIn(columns []string, w []domain.Where) ([][]string, error) {
	if len(columns) <= 0 {
		columns = t.GetFields()
	}

	for _, value := range w {
		if !utils.ContainsInside(t.GetFields(), value.Column) {
			return [][]string{}, errors.New("column not exists")
		}
	}

	lines := [][]string{columns}

	for _, row := range *t.Rows {
		line := []string{}
		for _, col := range columns {
			i := t.GetFieldIndex(col)
			line = append(line, utils.InterfaceToString(row[i]))
		}
		lines = append(lines, line)
	}

	return lines, nil
}
