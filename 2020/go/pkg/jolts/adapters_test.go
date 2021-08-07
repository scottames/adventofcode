package jolts_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/2020/go/pkg/jolts"
	"github.com/stretchr/testify/assert"
)

const (
	exampleOneMultiples = 35
	exampleTwoMultiples = 220
)

var (
	exampleOne = pie.Ints{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}
	exampleOneChainDiffs = jolts.ChainDiffs{
		1: 7,
		3: 5,
	}
	exampleTwo = pie.Ints{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	exampleTwoChainDiffs = jolts.ChainDiffs{
		1: 22,
		3: 10,
	}
)

func Test_ChainDiffsExampleOne(t *testing.T) {
	expected := exampleOneChainDiffs
	actual := jolts.NewChainDiffs(exampleOne)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ChainDiffsExampleTwo(t *testing.T) {
	expected := exampleTwoChainDiffs
	actual := jolts.NewChainDiffs(exampleTwo)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ChainDiffsMultiplesExampleOne(t *testing.T) {
	expected := exampleOneMultiples
	actual, err := jolts.NewChainDiffs(exampleOne).Multiples()
	if err != nil {
		t.Error(err)
	}
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
func Test_ChainDiffsMultiplesExampleTwo(t *testing.T) {
	expected := exampleTwoMultiples
	actual, err := jolts.NewChainDiffs(exampleTwo).Multiples()
	if err != nil {
		t.Error(err)
	}
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
