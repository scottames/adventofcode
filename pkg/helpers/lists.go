package helpers

import "github.com/elliotchance/pie/pie"

// MissingInts returns a new slice of Ints that are missing from the given range if Ints
func MissingInts(ints pie.Ints) pie.Ints {
	missing := pie.Ints{}
	ints = ints.Sort()
	for i := ints.First(); i <= ints.Last(); i++ {
		if !ints.Contains(i) {
			missing = missing.Append(i)
		}
	}
	return missing
}
