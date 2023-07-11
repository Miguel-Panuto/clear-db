package table

import (
	"errors"
	"strings"

	"github.com/miguel-panuto/clear-db/src/utils"
)

func (f *Field) validateAndReturn(row []string) *string {
	defaultValue := "null"
	for _, el := range row {
		if !strings.Contains(el, "::") {
			return &el
		}

		values := utils.TrimSplit(el, "::")
		if values[0] == f.name {
			return &values[1]
		}

		for _, prop := range f.properties {
			if prop == "required" {
				return nil
			}
		}
	}

	return &defaultValue
}

func (t *Table) InsertNewRow(row []string) error {
	parsedRows := []interface{}{}
	for _, value := range t.Fields {
		el := value.validateAndReturn(row)
		if el == nil {
			return errors.New("column not finded")
		}

		parsedValue, err := getValueType(*el, value.data_type)
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
