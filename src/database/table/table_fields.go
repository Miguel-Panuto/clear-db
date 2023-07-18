package table

import "strings"

func (t *Table) GetFieldsString() string {
	names := ""
	dataTypes := ""

	for i, value := range t.Fields {
		dataTypes += value.data_type
		names += value.name

		if len(value.properties) > 0 {
			dataTypes += "-" + strings.Join(value.properties, "-")
		}
		if i+1 < len(t.Fields) {
			names += ";"
			dataTypes += ";"
		}
	}
	return names + "\n" + dataTypes
}

func (t *Table) GetFields() []string {
	arr := make([]string, len(t.Fields))

	for i := range t.Fields {
		arr[i] = t.Fields[i].name
	}

	return arr
}
