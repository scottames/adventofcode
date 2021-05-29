package aoc2020

import (
	"fmt"
	"github.com/scottames/adventofcode/2020/go/pkg/trajectory"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day03 = 3
)

var day3Part2Coordinates = trajectory.NewCoordinateses(
	[][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}...)

// Day03 - Day 3 Part 1 & 2
func Day03() error {
	input, err := helpers.ReadInput(year, day03)
	if err != nil {
		return err
	}

	var hill = trajectory.NewHill(input)

	// Part 1
	helpers.PrintPart1()
	fmt.Printf("Number of trees encountered: %d\n\n",
		trajectory.NewSlope(hill, trajectory.NewCoordinates(3, 1)).
			Run().
			TreesEncountered(),
	)

	// Part 2
	helpers.PrintPart2()
	fmt.Printf("Product of number of trees encountered: %d\n\n",
		trajectory.NewSlopes(hill, day3Part2Coordinates...).
			Run().
			TreesEncountered().
			Product(),
	)

	return nil
}
