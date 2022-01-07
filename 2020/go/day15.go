package aoc2020

import (
	"fmt"

	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

// Day15 - Day 15 Part 1 & 2
func Day15() error {
	input := pie.Ints{0, 13, 16, 17, 1, 10, 6}

	// Part 1
	helpers.PrintPart1()
	fmt.Printf("2020th number spoken: %d\n", input[0]) // FIXME

	// Part 2
	helpers.PrintPart2()

	return nil
}
