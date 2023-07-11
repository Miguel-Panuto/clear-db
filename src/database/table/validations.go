package table

import (
	"errors"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var validDataType []string = []string{"string", "id", "int", "float", "boolean"}

func isValidDataType(dataType string) error {

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

func getValueType(value string, dataType string) (interface{}, error) {
	if err := isValidDataType(dataType); err != nil {
		return nil, err
	}

	switch dataType {
	case "string":
		return value, nil
	case "int":
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return strconv.Itoa(v), nil

	case "float":
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case "id":
		return getId(), nil

	case "boolean":
		if value == "true" {
			return true, nil
		}

		if value == "false" {
			return false, nil
		}

		return nil, errors.New("not valid value for boolean")

	default:
		return nil, errors.New("unknown dataType")
	}
}
