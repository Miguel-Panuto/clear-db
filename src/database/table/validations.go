package table

import (
	"errors"
)

func isValidDataType(dataType string) error {
	validDataType := []string{"string", "id", "int", "float", "json", "boolean"}

	for _, value := range validDataType {
		if dataType == value {
			return nil
		}
	}

	return errors.New("data type not suported")
}
