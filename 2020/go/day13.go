package aoc2020

import (
	"fmt"

	"github.com/scottames/adventofcode/2020/go/pkg/bus"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day13 = 13
)

// Day13- Day 13 Part 1 & 2
func Day13() error {
	input, err := helpers.ReadInput(year, day13)
	if err != nil {
		return err
	}

	schedule, err := bus.ReadSchedule(helpers.StringSplitNewlinesStrings(input))
	if err != nil {
		return err
	}

	next := schedule.NextDepartureID()
	wait := next - schedule.EarliestDeparture()

	// Part 1
	helpers.PrintPart1()
	fmt.Printf("Bus id '%d' departs in '%d' min with product: %d\n\n", next, wait, next*wait)

	// Part 2
	helpers.PrintPart2()
	fmt.Printf("Earliest timestamp with offset alignment: %d\n\n", schedule.EarliestTimestampOffsetAlignment())

	return nil
}
