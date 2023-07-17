package engine_struct

type TableCreation struct {
	DbName string
	Fields []string
}

type RowInsert struct {
	TableName string
	Row       []string
}

type FindIn struct {
	Columns   []string
	Where     []Where
	TableName string
}

type Where struct {
	Column   string
	Operator string
	Value    string
}
