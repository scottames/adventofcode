package aoc2020

import (
	"fmt"

	"github.com/scottames/adventofcode/2020/go/pkg/jolts"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day10 = 10
)

// Day10 - Day 10 Part 1 & 2
func Day10() error {
	ints, err := helpers.ParseLinesToInts(year, day10)
	helpers.ExitOnError(err)

	// Part 1
	helpers.PrintPart1()
	chain := jolts.NewChain(ints)
	multiples, err := chain.Multiples()
	helpers.ExitOnError(err)
	fmt.Printf(
		"Ones * Threes: %d\n\n",
		multiples,
	)

	// Part 2
	helpers.PrintPart2()
	fmt.Printf(
		"Distinct Arrangements: %d\n\n",
		chain.DistinctArrangements(),
	)

	return nil
}
