package aoc2020

import (
	"fmt"
	"github.com/scottames/adventofcode/2020/go/pkg/trajectory"
	"github.com/scottames/adventofcode/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const day03SampleSlope = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func Test_Day03Part1Sample1(t *testing.T) {
	expected := 7
	actual := trajectory.NewSlope(
		trajectory.NewHill([]byte(day03SampleSlope)),
		trajectory.NewCoordinates(3, 1)).
		Run().TreesEncountered()
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day03Part1(t *testing.T) {
	expected := 262
	input, err := helpers.ReadInput(year, day03)
	if err != nil {
		t.Errorf(err.Error())
	}
	actual := trajectory.NewSlope(
		trajectory.NewHill(input), trajectory.NewCoordinates(3, 1)).
		Run().
		TreesEncountered()
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day03Part2Sample1(t *testing.T) {
	var hill = trajectory.NewHill([]byte(day03SampleSlope))
	expected := 336
	actual := trajectory.NewSlopes(hill, day3Part2Coordinates...).
		Run().
		TreesEncountered().
		Product()
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Day03Part2(t *testing.T) {
	input, err := helpers.ReadInput(year, day03)
	if err != nil {
		t.Errorf(err.Error())
	}
	var hill = trajectory.NewHill(input)
	expected := 2698900776
	actual := trajectory.NewSlopes(hill, day3Part2Coordinates...).
		Run().
		TreesEncountered().
		Product()
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
