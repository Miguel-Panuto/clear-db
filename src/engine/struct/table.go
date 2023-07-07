package engine_struct

type TableCreation struct {
	DbName string
	Fields []string
}

type RowInsert struct {
	TabName string
	Row     []string
}
