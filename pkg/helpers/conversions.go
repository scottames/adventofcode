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

func CovertBytesToStrings(input []byte, split string) pie.Strings {
	return strings.Split(strings.TrimSpace(string(input)), split)
}

func StringSplitCommaToInts(csv []byte) (pie.Ints, error) {
	return covertToInts(csv, ",")
}

func StringSplitNewlinesToInts(input []byte) (pie.Ints, error) {
	return covertToInts(input, "\n")
}

// StringSplitNewlinesStrings returns a set of strings split on a single newline
func StringSplitNewlinesStrings(input []byte) pie.Strings {
	return strings.Split(strings.TrimSpace(string(input)), "\n")
}

// StringSplitNewlinesNewlinesStrings returns a set of strings split on two newlines
func StringSplitNewlinesNewlinesStrings(input []byte) pie.Strings {
	return strings.Split(strings.TrimSpace(string(input)), "\n\n")
}

func StringSplitCommaToStrings(csv []byte) pie.Strings {
	return CovertBytesToStrings(csv, ",")
}
