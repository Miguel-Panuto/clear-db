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
		str := fmt.Sprintf("%v", el)
		arr = append(arr, str)
	}

	return arr
}
