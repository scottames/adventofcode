package aoc2020

import (
	"fmt"
	"github.com/scottames/adventofcode/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const day02SampleData = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func Test_Day02Part1Sample1(t *testing.T) {
	expected := []int{2}
	actual, err := validatePassword(helpers.CovertBytesToStrings([]byte(day02SampleData), "\n"), validatorPart1)
	if err != nil {
		t.Errorf(err.Error())
	}
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day02Part2Sample1(t *testing.T) {
	expected := []int{1}
	actual, err := validatePassword(helpers.CovertBytesToStrings([]byte(day02SampleData), "\n"), validatorPart2)
	if err != nil {
		t.Errorf(err.Error())
	}
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day02Part1(t *testing.T) {
	get, err := helpers.ReadInput(2020, 2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expected := []int{645}
	actual, err := validatePassword(helpers.CovertBytesToStrings(get, "\n"), validatorPart1)
	if err != nil {
		t.Errorf(err.Error())
	}
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day02Part2(t *testing.T) {
	get, err := helpers.ReadInput(2020, 2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expected := []int{737}
	actual, err := validatePassword(helpers.CovertBytesToStrings(get, "\n"), validatorPart2)
	if err != nil {
		t.Errorf(err.Error())
	}
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
