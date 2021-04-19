package aoc2019

import (
	"fmt"
	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

// Day01 - Day 1 Part 1 & 2
func Day01() {
	get, err := helpers.ReadInput(2019, 1)
	helpers.ExitOnError(err)

	input, err := helpers.StringSplitNewlinesToInts(get)
	helpers.ExitOnError(err)

	day01Part1(input)
	day01Part2(input)
}

func day01Part1(input pie.Ints) {
	fmt.Println("--- Part One ---")
	fmt.Println()

	fmt.Println(input.Map(calculateFuel).Sum())
	fmt.Println()
}

func day01Part2(input pie.Ints) {
	fmt.Println("--- Part Two ---")
	fmt.Println()

	// 4905116
	fmt.Println("# Recursion:")
	fmt.Println("  ", sumFuelNeeded(input))

	fmt.Println("\n# Map:")
	fuelNeeded := input.Map(calculateFuelRecursive).Sum()
	fmt.Println("  ", fuelNeeded)
}

func sumFuelNeeded(inputMass pie.Ints) int {
	fuelNeeded := 0

	for _, i := range inputMass {
		fuelNeeded += calculateFuelRecursive(i)
	}

	return fuelNeeded
}

func calculateFuel(mass int) int {
	fuel := mass/3 - 2
	if fuel <= 0 {
		return 0
	}

	return fuel
}

func calculateFuelRecursive(mass int) int {
	fuel := mass/3 - 2
	if fuel <= 0 {
		return 0
	}

	return fuel + calculateFuelRecursive(fuel)
}
