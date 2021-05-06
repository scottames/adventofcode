package aoc2019

import (
	"fmt"
	"testing"

	"github.com/scottames/adventofcode/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

const (
	day3example1 = "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"
	day3example2 = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
)

func TestManhattanDistance_Day3Part1Example1(t *testing.T) {
	exampleOne := day3example1
	expected := 159
	actual := breakInputToWires([]byte(exampleOne)).findPoint(closestIntersectToStart)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestManhattanDistance_Day3Part1Example2(t *testing.T) {
	exampleOne := day3example2
	expected := 135
	actual := breakInputToWires([]byte(exampleOne)).findPoint(closestIntersectToStart)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestManhattanDistance_Day3Part1(t *testing.T) {
	input, err := helpers.ReadInput(2019, 3)
	if err != nil {
		assert.Error(t, err)
	}
	expected := 1211
	actual := breakInputToWires(input).findPoint(closestIntersectToStart)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestShortestDistanceToIntersect_Day3Part2Example1(t *testing.T) {
	exampleOne := day3example1
	expected := 610
	actual := breakInputToWires([]byte(exampleOne)).findPoint(shortestDistanceToIntersect)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestShortestDistanceToIntersect_Day3Part2Example2(t *testing.T) {
	exampleOne := day3example2
	expected := 410
	actual := breakInputToWires([]byte(exampleOne)).findPoint(shortestDistanceToIntersect)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestShortestDistanceToIntersect(t *testing.T) {
	input, err := helpers.ReadInput(2019, 3)
	if err != nil {
		assert.Error(t, err)
	}
	expected := 101386
	actual := breakInputToWires(input).findPoint(shortestDistanceToIntersect)
	msg := fmt.Sprintf("Expected %d. Got %d.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
