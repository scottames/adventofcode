//+build mage

package main

import (
	"fmt"

	aoc2019 "github.com/scottames/adventofcode/2019/go"
	aoc2020 "github.com/scottames/adventofcode/2020/go"

	//   "strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// TODO: make this more dynamic
var aoc map[int]map[int]func() error = map[int]map[int]func() error{
	2019: {
		1: aoc2019.Day01,
		2: aoc2019.Day02,
		3: aoc2019.Day03,
	},
	2020: {
		1: aoc2020.Day01,
		2: aoc2020.Day02,
		3: aoc2020.Day03,
		4: aoc2020.Day04,
		5: aoc2020.Day05,
		6: aoc2020.Day06,
		7: aoc2020.Day07,
		8: aoc2020.Day08,
	},
}

type Go mg.Namespace

// Run | run a given year / day in Golang
func (Go) Run(year int, day int) {
	fn, ok := aoc[year][day]
	if !ok {
		fmt.Println("Invalid Year/Date combination.")
		return
	}

	fn()
}

type Rust mg.Namespace

// Run | run a given year / day in Rust
func (Rust) Run(year int, day int) {
	// TODO: replace with proper AOC helper commands to call year/day
	manifestPath := fmt.Sprintf("--manifest-path=%d/rust/Cargo.toml", year)
	sh.RunV("cargo", "run", manifestPath)
}

type Python mg.Namespace

// Run | run a given year / day in Python
func (Python) Run(year int, day int) {
	sh.RunV("python3", fmt.Sprintf("%d/python/day%02d.py", year, day))
}
