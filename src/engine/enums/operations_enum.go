package engine_enums

type Operations int

const (
	CREATE_DATABASE Operations = iota
	LIST_DATABASES
	USE_DATABASE
	CREATE_TABLE
)
