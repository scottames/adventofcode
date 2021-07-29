package aoc2020

import (
	"fmt"

	"github.com/scottames/adventofcode/2020/go/pkg/handheld"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day08 = 8
)

// Day08 - Day 8 Part 1 & 2
func Day08() error {
	input, err := helpers.ReadInput(year, day08)
	if err != nil {
		return err
	}

	// Part 1
	helpers.PrintPart1()
	result := handheld.NewInstructions(string(input)).Execute()
	fmt.Printf(
		"Accumulator value: %d\n\n",
		result,
	)

	// Part 2
	helpers.PrintPart2()
	fmt.Printf(
		"Accumulator value: %d\n\n",
		handheld.NewInstructions(string(input)).Fix(),
	)
	return nil
}
