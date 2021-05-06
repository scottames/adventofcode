package aoc2020

import (
	"fmt"
	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const sum = 2020

// Day01 - Day 1 Part 1 & 2
func Day01() error {
	get, err := helpers.ReadInput(2020, 1)
	if err != nil {
		return err
	}

	input, err := helpers.StringSplitNewlinesToInts(get)
	if err != nil {
		return err
	}

	day01Part1(input)
	day01Part2(input)

	return nil

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
	fi := -1
	bar := func(n int) bool {
		fi++
		for i, k := range list {
			if fi != i && n+k == sum {
				return true
			}
		}
		return false
	}
	return list.Filter(bar)
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
	fi := -1
	bar := func(n int) bool {
		fi++
		for i, k := range list {
			x := sum - (n + k)
			if fi != i && list.Contains(x) {
				return true
			}
		}
		return false
	}
	return list.Filter(bar)
}
