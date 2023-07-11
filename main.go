package main

import (
	"github.com/miguel-panuto/clear-db/src/cli"
	"github.com/miguel-panuto/clear-db/src/database/table"
)

func main() {
	table.StartIds(10)
	cli.StartCli()
}
