package aoc2020

import (
	"fmt"
	"github.com/scottames/adventofcode/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

var day01SampleData = `1721
979
366
299
675
1456`

func Test_Day01Part1Sample1(t *testing.T) {
	data, err := helpers.StringSplitNewlinesToInts([]byte(day01SampleData))
	if err != nil {
		t.Errorf(err.Error())
	}

	expected := 514579
	actual := day01Part1(data)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day01Part2Sample1(t *testing.T) {
	data, err := helpers.StringSplitNewlinesToInts([]byte(day01SampleData))
	if err != nil {
		t.Errorf(err.Error())
	}

	expected := 241861950
	actual := day01Part2(data)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
