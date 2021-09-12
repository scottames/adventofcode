package aoc2020

import (
	"fmt"
	"strings"

	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	day06 = 6
)

// Day06 - Day 6 Part 1 & 2
func Day06() error {
	input, err := helpers.ReadInput(year, day06)
	if err != nil {
		return err
	}

	groups := helpers.StringSplitNewlinesNewlinesStrings(input)

	// Part 1
	helpers.PrintPart1()
	fmt.Printf("Sum of counts: %d\n", sumCounts(groups))
	fmt.Println()

	// Part 2
	helpers.PrintPart2()
	fmt.Printf("Sum of counts: %d\n", sumCountsEveryone(groups))
	fmt.Println()

	return nil
}

// TODO: redo this solution in more efficient & cleaner way
func sumCounts(groups pie.Strings) int {
	answers := pie.Ints{}
	for _, group := range groups {
		trimmed := uniqueString(strings.Replace(group, "\n", "", -1))
		answers = answers.Append(len(trimmed))
	}
	return answers.Sum()
}

func sumCountsEveryone(groups pie.Strings) int {
	var y pie.Strings
	for _, group := range groups {
		var a map[rune]int = make(map[rune]int)
		people := strings.Split(group, "\n")
		for _, person := range people {
			for _, answer := range person {
				a[answer]++
			}
		}
		for answer, number := range a {
			if number == len(people) {
				y = y.Append(string(answer))
			}
		}
	}
	return y.Len()
}

func uniqueString(str string) map[string]int {
	result := map[string]int{}
	for _, c := range str {
		s := string(c)
		result[s]++

	}
	return result
}
