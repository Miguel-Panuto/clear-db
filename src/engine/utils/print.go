package engine_utils

import (
	"fmt"
	"strings"
)

func PrintTable(header []string, rows [][]string) {
	columnLengths := make([]int, len(header))

	for i, h := range header {
		columnLengths[i] = len(h)
		for _, row := range rows {
			if len(row[i]) > columnLengths[i] {
				columnLengths[i] = len(row[i])
			}
		}
	}

	printRow := func(row []string) {
		for i, col := range row {
			fmt.Print("| ", col, strings.Repeat(" ", columnLengths[i]-len(col)), " ")
		}
		fmt.Println("|")
	}

	printRow(header)

	for _, row := range rows {
		printRow(row)
	}
}
