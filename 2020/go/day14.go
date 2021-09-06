package aoc2020

import (
	"fmt"

	"github.com/scottames/adventofcode/2020/go/pkg/port"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day14 = 14
)

// Day14 - Day 14 Part 1 & 2
func Day14() error {
	input, err := helpers.ReadInput(year, day14)
	if err != nil {
		return err
	}

	prog, err := port.InitializeProgram(helpers.StringSplitNewlinesStrings(input))
	if err != nil {
		return err
	}

	// Part 1
	helpers.PrintPart1()
	fmt.Printf("Sum of program values: %d\n", prog.Sum())

	// Part 2
	helpers.PrintPart2()

	return nil
}
