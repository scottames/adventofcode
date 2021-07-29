package aoc2020

import (
	"fmt"

	"github.com/scottames/adventofcode/2020/go/pkg/luggage"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day07     = 7
	shinyGold = "shiny gold"
)

// Day07 - Day 7 Part 1 & 2
func Day07() error {
	input, err := helpers.ReadInput(year, day07)
	if err != nil {
		return err
	}

	bagsSS := helpers.StringSplitNewlinesStrings(input)
	bags := luggage.NewBags(bagsSS)

	// Part 1
	helpers.PrintPart1()
	fmt.Printf(
		"Number of bags that can hold '%s': %d\n\n",
		shinyGold,
		bags.NumBagsCanContainBag(shinyGold),
	)

	// Part 2
	helpers.PrintPart2()
	fmt.Printf(
		"Number of bags that '%s' holds: %d\n\n",
		shinyGold,
		bags.NumBagHolds(shinyGold),
	)

	return nil
}
