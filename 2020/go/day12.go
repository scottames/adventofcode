package aoc2020

import (
	"fmt"

	"github.com/scottames/adventofcode/2020/go/pkg/ferry"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day12 = 12
)

// Day12 - Day 12 Part 1 & 2
func Day12() error {
	input, err := helpers.ReadInput(year, day12)
	helpers.ExitOnError(err)

	a1, err := ferry.NewActions(
		helpers.CovertBytesToStrings(input, "\n"),
		// refactor of part 1 inspired by @viking66:
		// https://github.com/viking66/adventofcode/blob/master/src/Day12.hs
		nil, false,
	)
	if err != nil {
		return err
	}
	md1, err := a1.ManhattanDistance()
	if err != nil {
		return err
	}

	// Part 1
	helpers.PrintPart1()
	fmt.Printf(
		"Manhattan distance: %d\n\n",
		md1,
	)

	// Part 2
	helpers.PrintPart2()

	a2, err := ferry.NewActions(
		helpers.CovertBytesToStrings(input, "\n"),
		&ferry.Point{X: 10, Y: 1}, true,
	)
	if err != nil {
		return err
	}
	md2, err := a2.ManhattanDistance()
	if err != nil {
		return err
	}
	fmt.Printf(
		"Manhattan distance: %d\n\n",
		md2,
	)

	return nil
}
