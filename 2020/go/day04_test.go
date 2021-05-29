package aoc2020

import (
	"fmt"
	"testing"

	"github.com/scottames/adventofcode/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

const day04SampleBatchFile = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func Test_Day04Part1Sample1(t *testing.T) {
	expected := 2
	actual := len(validBasicPassportsFromStringSlice(helpers.CovertBytesToStrings([]byte(day04SampleBatchFile), "\n\n")))
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day04Part1(t *testing.T) {
	expected := 202
	input, err := helpers.ReadInput(year, day04)
	if err != nil {
		t.Errorf(err.Error())
	}
	actual := len(validBasicPassportsFromStringSlice(helpers.CovertBytesToStrings(input, "\n\n")))
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day04Part2Sample1(t *testing.T) {
	expected := 2
	actual := len(validPassportsFromStringSlice(helpers.CovertBytesToStrings([]byte(day04SampleBatchFile), "\n\n")))
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day04Part2(t *testing.T) {
	expected := 137
	input, err := helpers.ReadInput(year, day04)
	if err != nil {
		t.Errorf(err.Error())
	}
	actual := len(validPassportsFromStringSlice(helpers.CovertBytesToStrings(input, "\n\n")))
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
