package helpers

import (
	"github.com/elliotchance/pie/pie"
	"strconv"
	"strings"
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

func StringSplitNewlinesStrings(input []byte) pie.Strings {
	return strings.Split(string(input), "\n")
}

func StringSplitCommaToStrings(csv []byte) pie.Strings {
	return CovertBytesToStrings(csv, ",")
}
