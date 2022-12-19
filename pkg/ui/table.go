package ui

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func Table(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(data[0])

	for _, v := range data[1:] {
		table.Append(v)
	}
	table.Render()
}
