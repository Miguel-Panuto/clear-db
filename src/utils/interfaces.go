package utils

import "fmt"

func ToInterfaceArr(args []string) []interface{} {
	arr := []interface{}{}

	for _, el := range args {
		arr = append(arr, el)
	}

	return arr
}

func ToStringArr(args []interface{}) []string {
	arr := []string{}

	for _, el := range args {
		arr = append(arr, InterfaceToString(el))
	}

	return arr
}

func InterfaceToString(el interface{}) string {
	return fmt.Sprintf("%v", el)
}
