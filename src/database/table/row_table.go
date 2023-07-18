package table

func (t *Table) indexOfColumn(name string) int {
	for i := range t.Fields {
		if name == t.Fields[i].name {
			return i
		}
	}
	return -1
}
