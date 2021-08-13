package trajectory

import (
	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

// NewHill returns a pointer to pie.Strings which represent the given Hill
// with any empty strings filtered out
func NewHill(b []byte) *pie.Strings {
	if len(b) == 0 {
		return &pie.Strings{}
	}
	rows := helpers.StringSplitNewlinesStrings(b)
	rows = rows.Filter(func(s string) bool {
		return s != ""
	})
	return &rows
}
