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

func createField(colums []string) ([]Field, error) {
	var fields []Field = make([]Field, len(colums))
	for i, value := range colums {
		splitedField := strings.Split(value, " ")
		if len(splitedField) < 1 {
			return nil, errors.New("data_type not setted")
		}
		fields[i] = Field{
			name:      splitedField[0],
			data_type: splitedField[1],
			required:  strings.Contains("required", strings.ToLower(value)),
		}
	}
	return fields, nil
}

func NewTable(name string, colums []string) (*Table, error) {
	fields, err := createField(colums)

	if err != nil {
		return nil, err
	}

	newTable := Table{Name: name, Fields: fields, Rows: []string{}}
	return &newTable, nil
}