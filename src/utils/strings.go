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
