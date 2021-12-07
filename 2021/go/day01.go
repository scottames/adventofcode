package aoc2021

import (
	"fmt"

	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	year = 2021
)

// Day01 - Day 1 Part 1 & 2
func Day01() error {
	get, err := helpers.ReadInput(year, 1)
	if err != nil {
		return err
	}

	input, err := helpers.StringSplitNewlinesToInts(get)
	if err != nil {
		return err
	}

	helpers.PrintPart1()
	fmt.Println("Number of increases: ", increasesInList(input))

	helpers.PrintPart2()
	fmt.Println("Number of three sequential increases: ", TripleSequenceIncreases(input))

	return nil
}

func increasesInList(ints pie.Ints) int {
	p := ints.First()
	count := 0

	ints.Filter(
		func(i int) bool {
			if i > p {
				count++
			}
			p = i

			return false
		},
	)

	return count
}

func TripleSequenceIncreases(ints pie.Ints) int {
	p := 0
	count := 0
	for i := 0; i < len(ints)-2; i++ {
		s := ints[i : i+3].Sum()
		if i > 0 && s > p {
			count++
		}
		p = s
	}

	return count
}
