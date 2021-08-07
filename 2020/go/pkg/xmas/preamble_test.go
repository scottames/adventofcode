package xmas

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/stretchr/testify/assert"
)

const (
	exampleFindFirstBreak = 127
	examplePreamble       = 5
)

var (
	exampleInts = pie.Ints{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	exampleIntsContiguousSet = exampleInts[2:6]
)

func Test_FindFirstBreak(t *testing.T) {
	expected := exampleFindFirstBreak
	actual := FindFirstBreak(exampleInts, examplePreamble)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_FindContiguousSet(t *testing.T) {
	expected := exampleIntsContiguousSet
	actual := FindContiguousSetWithSum(exampleInts, exampleFindFirstBreak)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
