package table

import (
	"strings"

	"github.com/miguel-panuto/clear-db/src/utils"
)

func createProperties(rawProperties string) []string {
	properties := strings.ToLower(rawProperties)
	properties = strings.TrimSpace(properties)

	return utils.Split(properties, " ")
}
