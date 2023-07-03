package table

import (
	"errors"
	"strings"
)

type Table struct {
	Name   string
	Fields []field
	Rows   []string
}

type field struct {
	name      string
	data_type string
	required  bool
}

func createField(colums []string) ([]field, error) {
	var fields []field = make([]field, len(colums))
	for i, value := range colums {
		splitedField := strings.Split(value, " ")
		if len(splitedField) < 1 {
			return nil, errors.New("data_type not setted")
		}
		fields[i] = field{
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
