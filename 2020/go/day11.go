package aoc2020

import (
	"fmt"

	"github.com/scottames/adventofcode/2020/go/pkg/seats"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day11 = 11
)

// Day11 - Day 11 Part 1 & 2
func Day11() error {
	input, err := helpers.ReadInput(year, day11)
	helpers.ExitOnError(err)

	lines := helpers.CovertBytesToStrings(input, "\n")

	// Part 1
	helpers.PrintPart1()
	fmt.Printf(
		"Occupied seats: %d\n\n",
		seats.NewArrangement(lines).NextUntilMatchingAdjacent().OccupiedSeats(),
	)

	// Part 2
	helpers.PrintPart2()
	fmt.Printf(
		"Occupied seats: %d\n\n",
		seats.NewArrangement(lines).NextUntilMatchingEightDirections().OccupiedSeats(),
	)

	return nil
}
