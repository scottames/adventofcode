package aoc2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scottames/adventofcode/2021/go/pkg/coordinates"
	"github.com/scottames/adventofcode/pkg/helpers"
)

// Day02 - Day 2 Part 1 & 2
func Day02() error {

	input, err := helpers.ReadInput(year, 2)
	if err != nil {
		return err
	}

	moves, err := newMoves(helpers.StringSplitNewlinesStrings(input))
	if err != nil {
		return err
	}

	part1Result, err := moves.navigate(directionWithoutAim)
	if err != nil {
		return err
	}

	helpers.PrintPart1()
	fmt.Println(part1Result.X * helpers.Absolute(part1Result.Y))

	part2Result, err := moves.navigate(directionWithAim)
	if err != nil {
		return err
	}

	helpers.PrintPart2()
	fmt.Println(part2Result.X * helpers.Absolute(part2Result.Y))

	return nil
}

// part 1
func directionWithoutAim(m move, aim int) (coordinates.Point, int) {
	switch m.direction {
	case forward:
		return coordinates.New(m.units, 0), 0
	case down:
		return coordinates.New(0, -m.units), 0
	case up:
		return coordinates.New(0, m.units), 0
	}

	return coordinates.New(0, 0), 0
}

// part 2
func directionWithAim(m move, aim int) (coordinates.Point, int) {
	switch m.direction {
	case forward:
		return coordinates.New(m.units, -(aim * m.units)), aim
	case down:
		return coordinates.New(0, 0), aim + m.units
	case up:
		return coordinates.New(0, 0), aim - m.units
	}

	return coordinates.New(0, 0), 0
}

const (
	forward direction = "forward"
	down    direction = "down"
	up      direction = "up"
)

type direction string

type move struct {
	direction direction
	units     int
}

func newMove(s string) (move, error) {
	result := move{}
	step := strings.Split(s, " ")
	if len(step) < 2 {
		return result, fmt.Errorf("expected 2 values, got less")
	}

	units, err := strconv.Atoi(step[1])
	if err != nil {
		return result, fmt.Errorf("invalid unit '%s'", step[1])
	}

	return move{
		direction: direction(step[0]),
		units:     units,
	}, nil
}

type moves []move

func newMoves(ss []string) (moves, error) {
	moves := []move{}
	for i, s := range ss {
		move, err := newMove(s)
		if err != nil {
			return nil, fmt.Errorf("unexpected input at %s (#%d) - %w", s, i, err)
		}

		moves = append(moves, move)
	}

	return moves, nil
}

func (m moves) navigate(
	fn func(m move, aim int) (coordinates.Point, int),
) (
	coordinates.Point,
	error,
) {
	coord := coordinates.New(0, 0)
	aim := 0
	for _, move := range m {
		d, a := fn(move, aim)

		coord = coord.Add(d)
		aim = a
	}

	return coord, nil
}
