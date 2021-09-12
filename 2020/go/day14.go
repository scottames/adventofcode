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

	prog := port.NewProgram(helpers.StringSplitNewlinesStrings(input))

	// Part 1
	helpers.PrintPart1()
	v1, err := prog.V1().Run()
	if err != nil {
		return err
	}
	fmt.Printf("Sum of program values: %d\n", v1.Sum())

	// Part 2
	helpers.PrintPart2()
	v2, err := prog.V2().Run()
	if err != nil {
		return err
	}
	fmt.Printf("Sum of program values: %d\n", v2.Sum())

	return nil
}
