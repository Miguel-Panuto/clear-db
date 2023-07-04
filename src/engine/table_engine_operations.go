package engine

import engine_struct "github.com/miguel-panuto/clear-db/src/engine/struct"

func (e *Engine) createTable(data engine_struct.TableCreation) error {
	for _, value := range data.Fields {
		println(value)
	}
	return nil
}
