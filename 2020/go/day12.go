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

	actions, err := ferry.NewActions(helpers.CovertBytesToStrings(input, "\n"))
	if err != nil {
		return err
	}

	// Part 1
	helpers.PrintPart1()
	fmt.Printf(
		"Manhattan distance: %d\n\n",
		actions.ManhattanDistance(),
	)

	// Part 2
	helpers.PrintPart2()

	waypoints, err := ferry.NewWaypoints(helpers.CovertBytesToStrings(input, "\n"))
	if err != nil {
		return err
	}
	md, err := waypoints.ManhattanDistance()
	if err != nil {
		return err
	}
	fmt.Printf(
		"Manhattan distance: %d\n\n",
		md,
	)

	return nil
}
