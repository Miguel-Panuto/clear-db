package table

func (t *Table) InsertRow(row []string) error {
	*t.Rows = append(*t.Rows, row)

	return nil
}
