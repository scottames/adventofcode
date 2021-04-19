package aoc2019

import (
	"fmt"
	"github.com/scottames/adventofcode/pkg/helpers"
	"os"
)

// TODO: Use type / struct to solve

// example
var example = []int{
	1, 9, 10, 3,
	2, 3, 11, 0,
	99, 30, 40, 50,
}

// Day02 - https://adventofcode.com/2019/day/2
func Day02() {
	day02Part1()
	day02Part2()
}

// day02Part1 - AOC 2019 Day 2 Part 1
func day02Part1() {
	fmt.Println("--- 2019 Day 2 Part 1 ---")
	fmt.Println()

	get, err := helpers.ReadInput(2019, 2)
	helpers.ExitOnError(err)

	input, err := helpers.StringSplitCommaToInts(get)
	helpers.ExitOnError(err)

	exampleOutput := computer(example, nil, nil)
	fmt.Println("# Example\n\n", exampleOutput)
	fmt.Println()

	// part 1
	//
	// Once you have a working computer, the first step is to
	// restore the gravity assist program (your puzzle input)
	// to the "1202 program alarm" state it had just before the
	// last computer caught fire. To do this, before running
	// the program, replace position 1 with the value 12 and
	noun := 12 // noun
	// replace position 2 with the value 2.
	verb := 2 // verb
	//
	// intCode[0] --> 4090701
	output := computer(input, &noun, &verb)
	fmt.Println(output)
	fmt.Println()
}

// day02Part2 - AOC 2019 Day 2 Part 2
func day02Part2() {
	fmt.Println("--- 2019 Day 2 Part 2 ---")
	fmt.Println()

	get, err := helpers.ReadInput(2019, 2)
	helpers.ExitOnError(err)

	input, err := helpers.StringSplitCommaToInts(get)
	helpers.ExitOnError(err)

	// part 2: What is 100 * noun + verb that produces out of 19690720?
	//
	// To complete the gravity assist, you need to determine
	//  what pair of inputs produces the output 19690720
	desired := 19690720
	n, v, err := completeManeuver(input, desired)
	if err != nil {
		fmt.Println("Error! ", err)
		os.Exit(1)
	}
	fmt.Printf(" noun: %d\n verb: %d\n desired output: %d\n\n", n, v, desired)
	fmt.Printf("100 * %d + %d = %d\n", n, v, 100*n+v)
}

// computer takes a memory input and returns a calculated output
//  without modifying the input value
func computer(memInput []int, nounFix *int, verbFix *int) int {
	mem := make([]int, len(memInput))
	copy(mem, memInput)

	if nounFix != nil {
		mem[1] = *nounFix
	}
	if verbFix != nil {
		mem[2] = *verbFix
	}

	iter(mem)

	return mem[0]
}

func iter(mem []int) {
	const step = 4
	const nounPosition = 1
	const verbPosition = 2
	var position = 0
	for {
		addr := mem[position+3]
		opCode := mem[position]
		switch opCode {
		case 1:
			mem[addr] = helpers.AddInt(mem[mem[position+nounPosition]], mem[mem[position+verbPosition]])
			position += step
		case 2:
			mem[addr] = helpers.MultiplyInt(mem[mem[position+nounPosition]], mem[mem[position+verbPosition]])
			position += step
		default:
			return
		}
	}
}

// completeManeuver determines the appropriate noun and verb
//  needed to obtain a desired output
func completeManeuver(mem []int, desiredOutput int) (int, int, error) {

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			result := computer(mem, &i, &j)
			if result == desiredOutput {
				return i, j, nil
			}
		}

	}

	return 0, 0, fmt.Errorf("unable to calculate the nount & verb necessary to complete maneuver")
}
