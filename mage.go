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
	"github.com/scottames/cmder"
)

// TODO: make this more dynamic -> merge into Dayer interface
var aoc map[int]map[int]func() error = map[int]map[int]func() error{
	2019: {
		1: aoc2019.Day01,
		2: aoc2019.Day02,
		3: aoc2019.Day03,
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
	},
	2021: {
		1: aoc2021.Day01,
		2: aoc2021.Day02,
	},
}

type Go mg.Namespace

// Run | run a given year / day in Golang
func (Go) Run(year int, day int) error {
	if year == 2021 {
		d, ok := aoc2021.Days[day]
		if !ok {
			return fmt.Errorf("Invalid Year/Date combination - %d/%d not yet implemented", year, day)
		}

		input, err := helpers.ReadInput(year, day)
		if err != nil {
			return err
		}

		err = d.Parse(input)
		if err != nil {
			return err
		}

		err = d.Part1()
		if err != nil {
			return err
		}

		err = d.Part2()
		if err != nil {
			return err
		}

		fmt.Println(d)

		return nil
	}

	fn, ok := aoc[year][day]
	if !ok {
		return fmt.Errorf("Invalid Year/Date combination - %d/%d not yet implemented", year, day)
	}

	return fn()
}

func (Go) Test() error {
	return cmder.New("go", "test", "./...").Run()
}

type Rust mg.Namespace

// Run | run a given year / day in Rust
func (Rust) Run(year int, day int) {
	// TODO: replace with proper AOC helper commands to call year/day
	manifestPath := fmt.Sprintf("--manifest-path=%d/rust/Cargo.toml", year)
	cmder.New("cargo", "run", manifestPath).Run()
}

type Py mg.Namespace

// Run | run a given year / day in Python
func (Py) Run(year int, day int) error {
	file := fmt.Sprintf("%d/python/day%02d.py", year, day)

	if !helpers.FileExists(file) {
		return fmt.Errorf("Python '%d' day '%d' not found", year, day)
	}

	return cmder.New("python3", file).Run()
}
