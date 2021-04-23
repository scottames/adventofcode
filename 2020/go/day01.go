package aoc2020

import (
	"fmt"
	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const twentyTwenty = 2020

// Day01 - Day 1 Part 1 & 2
func Day01() {
	get, err := helpers.ReadInput(2020, 1)
	helpers.ExitOnError(err)

	input, err := helpers.StringSplitNewlinesToInts(get)
	helpers.ExitOnError(err)

	day01Part1(input)
	day01Part2(input)

}

func day01Part1(input pie.Ints) int {
	fmt.Println("--- Part One ---")

	r := twoNumbersSum2020(input)
	fmt.Println("numbers: ", r)
	fmt.Println("product: ", r.Product())
	fmt.Println()

	return r.Product()
}

func twoNumbersSum2020(list pie.Ints) pie.Ints {
	var result pie.Ints
	bar := func(i int) bool {
		for _, k := range list {
			if i+k == twentyTwenty {
				return true
			}
		}
		return false
	}
	for i := 0; i < 2; i++ {
		result = list.Filter(bar)
	}
	return result
}

func day01Part2(input pie.Ints) int {
	fmt.Println("--- Part Two ---")

	r := threeNumbersSum2020(input)
	fmt.Println("numbers: ", r)
	fmt.Println("product: ", r.Product())
	fmt.Println()

	return r.Product()
}

func threeNumbersSum2020(list pie.Ints) pie.Ints {
	filterIndex := -1
	bar := func(n int) bool {
		filterIndex++
		for i, k := range list {
			x := twentyTwenty - (n + k)
			if filterIndex != i && list.Contains(x) {
				return true
			}
		}
		return false
	}
	return list.Filter(bar)
}
