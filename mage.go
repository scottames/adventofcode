//go:build mage
// +build mage

package main

import (
	"fmt"

	aoc2019 "github.com/scottames/adventofcode/2019/go"
	aoc2020 "github.com/scottames/adventofcode/2020/go"
	aoc2021 "github.com/scottames/adventofcode/2021/go"
	"github.com/scottames/adventofcode/pkg/helpers"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// TODO: make this more dynamic
var aoc map[int]map[int]func() error = map[int]map[int]func() error{
	2019: {
		1: aoc2019.Day01,
		2: aoc2019.Day02,
		3: aoc2019.Day03,
		// 4:  aoc2019.Day04,
		// 5:  aoc2019.Day05,
		// 6:  aoc2019.Day06,
		// 7:  aoc2019.Day07,
		// 8:  aoc2019.Day08,
		// 9:  aoc2019.Day09,
		// 10: aoc2019.Day10,
		// 11: aoc2019.Day11,
		// 12: aoc2019.Day12,
		// 13: aoc2019.Day13,
		// 14: aoc2019.Day14,
		// 15: aoc2019.Day15,
		// 16: aoc2019.Day16,
		// 17: aoc2019.Day17,
		// 18: aoc2019.Day18,
		// 19: aoc2019.Day19,
		// 20: aoc2019.Day20,
		// 21: aoc2019.Day21,
		// 22: aoc2019.Day22,
		// 23: aoc2019.Day23,
		// 24: aoc2019.Day24,
		// 25: aoc2019.Day25,
	},
	2020: {
		1:  aoc2020.Day01,
		2:  aoc2020.Day02,
		3:  aoc2020.Day03,
		4:  aoc2020.Day04,
		5:  aoc2020.Day05,
		6:  aoc2020.Day06,
		7:  aoc2020.Day07,
		8:  aoc2020.Day08,
		9:  aoc2020.Day09,
		10: aoc2020.Day10,
		11: aoc2020.Day11,
		12: aoc2020.Day12,
		13: aoc2020.Day13,
		14: aoc2020.Day14,
		// 15: aoc2020.Day15,
		// 16: aoc2020.Day16,
		// 17: aoc2020.Day17,
		// 18: aoc2020.Day18,
		// 19: aoc2020.Day19,
		// 20: aoc2020.Day20,
		// 21: aoc2020.Day21,
		// 22: aoc2020.Day22,
		// 23: aoc2020.Day23,
		// 24: aoc2020.Day24,
		// 25: aoc2020.Day25,
	},
	2021: {
		1: aoc2021.Day01,
		2: aoc2021.Day02,
	},
}

type Go mg.Namespace

// Run | run a given year / day in Golang
func (Go) Run(year int, day int) error {
	fn, ok := aoc[year][day]
	if !ok {
		return fmt.Errorf("Invalid Year/Date combination - %d/%d not yet implemented", year, day)
	}

	return fn()
}

func (Go) Test() error {
	return sh.RunV("go", "test", "./...")
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
func (Python) Run(year int, day int) error {
	file := fmt.Sprintf("%d/python/day%02d.py", year, day)

	if !helpers.FileExists(file) {
		return fmt.Errorf("Python '%d' day '%d' not found", year, day)
	}

	return sh.RunV("python3", file)
}
