package table

import "strings"

func createProperties(rawProperties string) []string {
	properties := strings.ToLower(rawProperties)
	properties = strings.TrimSpace(properties)

	return strings.Split(properties, " ")
}
