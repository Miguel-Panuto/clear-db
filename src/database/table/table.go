package table

import (
	"errors"
	"strings"
)

type Table struct {
	Name   string
	Fields []Field
	Rows   []string
}

type Field struct {
	name      string
	data_type string
	required  bool
}

func createField(columns []string) ([]Field, error) {
	var fields []Field = make([]Field, len(columns))
	for i, value := range columns {
		splitedField := strings.Split(value, " ")
		if len(splitedField) < 1 {
			return nil, errors.New("data_type not setted")
		}

		dataType := strings.TrimSpace(strings.ToLower(splitedField[1]))
		err := isValidDataType(dataType)

		if err != nil {
			return nil, err
		}

		fields[i] = Field{
			name:      splitedField[0],
			data_type: dataType,
			required:  strings.Contains("required", strings.ToLower(value)),
		}
	}
	return fields, nil
}

func NewTable(name string, columns []string) (*Table, error) {
	fields, err := createField(columns)

	if err != nil {
		return nil, err
	}

	newTable := Table{Name: name, Fields: fields, Rows: []string{}}
	return &newTable, nil
}
