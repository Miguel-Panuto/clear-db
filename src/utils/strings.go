package utils

import "strings"

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
