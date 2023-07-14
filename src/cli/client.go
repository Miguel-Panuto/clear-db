package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cli_print "github.com/miguel-panuto/clear-db/src/cli/print"
	"github.com/miguel-panuto/clear-db/src/engine"
)

func StartCli() {
	engine := engine.NewEngine()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		command, _ := reader.ReadString(';')
		command = strings.TrimSpace(command)
		value, err := engine.RunStatement(strings.ReplaceAll(command, ";", ""))
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		rows, ok := value.([][]string)
		if ok {
			cli_print.PrintRows(rows)
		}
	}
}
