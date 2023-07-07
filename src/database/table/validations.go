package table

import (
	"errors"
	"strings"

	"golang.org/x/exp/slices"
)

func isValidDataType(dataType string) error {
	validDataType := []string{"string", "id", "int", "float", "json", "boolean"}

	isValid := slices.Contains(validDataType, dataType)

	if !isValid {
		return errors.New("data type not suported: " + dataType)
	}

	return nil
}

func isValidProperties(properties []string) error {
	if len(properties) <= 0 || properties[0] == "" {
		return nil
	}

	validProperties := []string{"required", "unique"}

	if len(properties) > len(validProperties)+1 {
		return errors.New("there is more properties then the valids " + strings.Join(properties, " "))
	}

	for _, prop := range properties {
		if !slices.Contains(validProperties, prop) {
			return errors.New("invalid property: " + prop)
		}
	}

	return nil
}
