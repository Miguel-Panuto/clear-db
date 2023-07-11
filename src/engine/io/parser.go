package engine_io

import (
	"fmt"
	"strings"

	"github.com/miguel-panuto/clear-db/src/database"
	"github.com/miguel-panuto/clear-db/src/utils"
)

func parseToDatabase(s string) *database.Database {
	lines := utils.Split(s, "\n")
	dbName := lines[0]
	tables := strings.Split(strings.Replace(s, dbName+"\n", "", 1), "\n;;\n")
	db := database.NewDatabase(dbName)

	for _, rawTable := range tables {
		splitedRawTable := utils.Split(rawTable, "\n")
		tableName := splitedRawTable[0]
		fieldNames := strings.Split(splitedRawTable[1], ";")
		fieldAtt := strings.Split(splitedRawTable[2], ";")
		columns := []string{}

		for i := range fieldNames {
			columns = append(columns, fieldNames[i]+" "+strings.ReplaceAll(fieldAtt[i], "-", " "))
		}

		if err := db.NewTable(tableName, columns); err != nil {
			fmt.Println(err)
		}
		table, _ := db.FindTable(tableName)
		for i := 3; i < len(splitedRawTable); i++ {
			splitedString := utils.Split(splitedRawTable[i], ";")
			table.InsertFromReader(utils.ToInterfaceArr(splitedString))
		}
	}

	return db
}
