package aoc2021

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

const (
	day03Example string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
)

func Test_Day03(t *testing.T) {
	expected := Day03{
		input:     []byte(day03Example),
		report:    helpers.StringSplitNewlinesStrings([]byte(day03Example)),
		colLen:    5,
		gamma:     pie.Strings{"1", "0", "1", "1", "0"},
		epsilon:   pie.Strings{"0", "1", "0", "0", "1"},
		o2rating:  23,
		co2rating: 10,
	}

	actual := Day03{}
	err := actual.Parse([]byte(day03Example))
	if err != nil {
		t.Errorf(err.Error())
	}

	err = actual.Part1()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = actual.Part2()
	if err != nil {
		t.Errorf(err.Error())
	}

	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
