//+build mage

package main

import (
	"fmt"
	aoc2019 "github.com/scottames/adventofcode/2019/go"
	"github.com/scottames/adventofcode/2020/go"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// TODO: make this more dynamic
var aoc map[int]map[int]func() error = map[int]map[int]func() error{
	2019: {
		01: aoc2019.Day01,
		02: aoc2019.Day02,
	},
	2020: {
		01: aoc2020.Day01,
		02: aoc2020.Day02,
	},
}

type Go mg.Namespace

func (Go) Run(year int, day int) {
	fn, ok := aoc[year][day]
	if !ok {
		fmt.Println("Invalid Year/Date combination.")
		return
	}

	fn()
}

type Rust mg.Namespace

func (Rust) Run(year int, day int) {
	// TODO: replace with proper AOC helper commands to call year/day
	manifestPath := fmt.Sprintf("--manifest-path=%d/rust/Cargo.toml", year)
	sh.Run("cargo", "run", manifestPath)
}
