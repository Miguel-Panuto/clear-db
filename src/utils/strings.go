package utils

import (
	"fmt"
	"strings"
)

func Split(s string, sep string) []string {
	if len(s) == 0 {
		return []string{}
	}

	return strings.Split(s, sep)
}

func TrimSplit(s string, sep string) []string {
	arr := Split(s, sep)

	for i := range arr {
		arr[i] = strings.TrimSpace(arr[i])
	}

	return arr
}

func ContainsMany(s string, values ...string) bool {
	contains := true

	for _, el := range values {
		if !strings.Contains(s, el) {
			contains = false
		}
	}

	return contains
}

func SubString(s string, from string, until string) string {
	fromIndex := strings.Index(s, from) + 1
	untilIndex := strings.Index(s, until)
	substring := s[fromIndex:untilIndex]
	return substring
}

func VerifyLowerPrefix(s string, prefix string) bool {
	aux := strings.ToLower(s)

	aux = strings.TrimSpace(aux)

	return strings.HasPrefix(aux, prefix)
}

func VerifyLower(s string, command string) bool {
	aux := strings.ToLower(s)

	aux = strings.TrimSpace(aux)

	return aux == command
}

func MakeStringArr(i []any) []string {
	arr := []string{}
	for _, value := range i {
		arr = append(arr, fmt.Sprint(value))
	}

	return arr
}
