package aoc2021

import "github.com/jedib0t/go-pretty/v6/table"

// TODO: if this sticks around, put it in a separate package
func NewTableRounded() table.Writer {
	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)

	return t
}
