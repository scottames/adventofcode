package aoc2020

import (
	"fmt"

	"github.com/scottames/adventofcode/2020/go/pkg/xmas"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day09    = 9
	preamble = 25
)

// Day09 - Day 9 Part 1 & 2
func Day09() error {
	ints, err := helpers.ParseLinesToInts(year, day09)
	if err != nil {
		return err
	}

	// Part 1
	helpers.PrintPart1()
	p1 := xmas.FindFirstBreak(ints, preamble)
	fmt.Printf(
		"No sum: %d\n\n",
		p1,
	)

	// Part 2
	helpers.PrintPart2()
	p2 := xmas.SumSmallestAndLargestInts(
		xmas.FindContiguousSetWithSum(ints, p1),
	)
	fmt.Printf(
		"P2 Sum: %d\n\n",
		p2,
	)
	return nil
}
