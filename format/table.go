package format

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func TableFormat(headers []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)

	for _, row := range rows {
		table.Append(row)
	}
	table.Render() // Send output
}
