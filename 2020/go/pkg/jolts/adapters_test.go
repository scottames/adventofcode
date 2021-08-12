package jolts

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/stretchr/testify/assert"
)

const (
	exampleOneMultiples  = 35
	exampleOneVariations = 8
	exampleTwoMultiples  = 220
	exampleTwoVariations = 19208
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
	exampleOneArrangements = pie.Ints{
		1,
		1,
		1,
		1,
		2,
		4,
		4,
		4,
		8,
		8,
		8,
		8,
		8,
	}
	exampleOneChain = Chain{
		adapters:     exampleOne,
		arrangements: exampleOneArrangements,
		ordered:      orderAdapters(exampleOne),
		diffs: map[int]int{
			1: 7,
			3: 5,
		},
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
	exampleTwoArrangements = pie.Ints{
		1,
		1,
		2,
		4,
		7,
		7,
		7,
		14,
		28,
		49,
		49,
		49,
		49,
		98,
		196,
		196,
		196,
		392,
		392,
		392,
		392,
		784,
		1568,
		2744,
		2744,
		2744,
		2744,
		2744,
		2744,
		5488,
		10976,
		19208,
		19208,
	}
	exampleTwoChain = Chain{
		adapters:     exampleTwo,
		arrangements: exampleTwoArrangements,
		ordered:      orderAdapters(exampleTwo),
		diffs: map[int]int{
			1: 22,
			3: 10,
		},
	}
)

func Test_ChainExampleOne(t *testing.T) {
	expected := &exampleOneChain
	actual := NewChain(exampleOne)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ChainExampleTwo(t *testing.T) {
	expected := &exampleTwoChain
	actual := NewChain(exampleTwo)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ChainMultiplesExampleOne(t *testing.T) {
	expected := exampleOneMultiples
	actual, err := NewChain(exampleOne).Multiples()
	if err != nil {
		t.Error(err)
	}
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
func Test_ChainMultiplesExampleTwo(t *testing.T) {
	expected := exampleTwoMultiples
	actual, err := NewChain(exampleTwo).Multiples()
	if err != nil {
		t.Error(err)
	}
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ChainVariationNumExampleOne(t *testing.T) {
	expected := exampleOneVariations
	actual := NewChain(exampleOne).DistinctArrangements()
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_ChainVariationNumExampleTwo(t *testing.T) {
	expected := exampleTwoVariations
	actual := NewChain(exampleTwo).DistinctArrangements()
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
