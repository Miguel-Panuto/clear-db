package table

import (
	"fmt"

	"github.com/miguel-panuto/clear-db/src/domain"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func (t *Table) validateColumn(w []domain.Where) ([]int, error) {
	var indexes []int = make([]int, len(w))

	for _, v := range w {
		indexOf := t.indexOfColumn(v.Column)
		if indexOf < 0 {
			return indexes, fmt.Errorf("there is column named as %s", v.Column)
		}
		indexes = append(indexes, indexOf)
	}
	return indexes, nil
}

func (t *Table) findWithoutConditions(columns []string, row []interface{}) []string {
	line := []string{}
	for _, col := range columns {
		i := t.indexOfColumn(col)
		line = append(line, utils.InterfaceToString(row[i]))
	}
	return line
}

func (t *Table) FindIn(columns []string, w []domain.Where) ([][]string, error) {
	if len(columns) <= 0 {
		columns = t.GetFields()
	}

	indexesWhere, err := t.validateColumn(w)

	if err != nil {
		return [][]string{}, err
	}

	lines := [][]string{columns}

	for _, row := range *t.Rows {
		var line []string
		if len(indexesWhere) <= 0 {
			line = t.findWithoutConditions(columns, row)
		}
		lines = append(lines, line)
	}

	return lines, nil
}
