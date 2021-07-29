package aoc2020

import (
	"fmt"
	"github.com/scottames/adventofcode/2020/go/pkg/airplane"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day05 = 5
)

// Day05 - Day 5 Part 1 & 2
func Day05() error {
	input, err := helpers.ReadInput(year, day05)
	if err != nil {
		return err
	}

	_, seats := airplane.NewBoardingPasses(helpers.StringSplitNewlinesStrings(input))

	// Part 1
	helpers.PrintPart1()
	fmt.Printf("Highest seat ID: %d\n\n", seats.Last())

	// Part 2
	helpers.PrintPart2()
	fmt.Printf("Missing seat(s): %v\n\n", helpers.MissingInts(seats))

	return nil
}
