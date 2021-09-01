package helpers

import (
	"strconv"
	"strings"

	"github.com/elliotchance/pie/pie"
)

func covertToInts(input []byte, split string) (pie.Ints, error) {
	s := strings.Split(strings.TrimSpace(string(input)), split)
	ints := make([]int, len(s))
	for i, n := range s {
		int, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		ints[i] = int
	}
	return ints, nil
}

// CovertBytesToStrings returns a set of strings split from the input
func CovertBytesToStrings(input []byte, split string) pie.Strings {
	return strings.Split(strings.TrimSpace(string(input)), split)
}

// StringSplitCommaToInts returns a set of ints from a given csv
func StringSplitCommaToInts(csv []byte) (pie.Ints, error) {
	return covertToInts(csv, ",")
}

// StringSplitNewlinesToInts returns a set of ints split on a single newline
func StringSplitNewlinesToInts(input []byte) (pie.Ints, error) {
	return covertToInts(input, "\n")
}

// StringSplitNewlinesStrings returns a set of strings split on a single newline
// with leading and trailing whitespace removed
func StringSplitNewlinesStrings(input []byte) pie.Strings {
	return strings.Split(strings.TrimSpace(string(input)), "\n")
}

// StringSplitNewlinesNewlinesStrings returns a set of strings split on two newlines
// with leading and trailing whitespace removed
func StringSplitNewlinesNewlinesStrings(input []byte) pie.Strings {
	return strings.Split(strings.TrimSpace(string(input)), "\n\n")
}

// StringSplitCommaToStrings returns a set of strings from a given csv
func StringSplitCommaToStrings(csv []byte) pie.Strings {
	return CovertBytesToStrings(csv, ",")
}
