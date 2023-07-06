package engine

import "github.com/miguel-panuto/clear-db/src/database"

type Engine struct {
	selectedDatabase *database.Database
	databases        []*database.Database
}
