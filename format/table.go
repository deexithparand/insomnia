package format

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func TableFormat(headers []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)

	for _, row := range rows {
		for _, value := range row {
			fmt.Printf("%s ", value)
		}
		fmt.Println("\n")
		table.Append(row)
	}
	table.Render() // Send output
}
