package aoc2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elliotchance/pie/pie"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/scottames/adventofcode/2021/go/pkg/bingo"
	"github.com/scottames/adventofcode/pkg/helpers"
	"github.com/scottames/adventofcode/pkg/style"
)

type Day04 struct {
	input                    []byte
	boards                   bingo.Boards
	numbersDrawn             pie.Ints
	winningNumber            *int
	winningBoardNumber       *int
	winningBoardSumRemaining *int
}

func init() {
	Days[4] = &Day04{}
}

func (d *Day04) String() string {
	p1 := NewTableRounded()
	p1.AppendHeader(table.Row{"Part 1"})
	if d.winningNumber != nil &&
		d.winningBoardNumber != nil &&
		d.winningBoardSumRemaining != nil {
		p1.AppendRows([]table.Row{
			{style.Blue("Winning Number 🙌"), *d.winningNumber},
			{style.Yellow("Winning Board  🎉"), *d.winningBoardNumber},
			{style.Magenta("Remaining Sum of Board"), *d.winningBoardSumRemaining},
		})
		p1.AppendFooter(table.Row{
			style.Green("Score"), *d.winningBoardSumRemaining * *d.winningNumber,
		})
	} else {
		p1.AppendRows([]table.Row{
			{style.Red("No winner found"), "😢"},
		})
	}

	return p1.Render()
}

func (d *Day04) Parse(input []byte) error {
	if input == nil {
		return fmt.Errorf("empty input value, unable to parse")
	}

	d.input = input
	data := helpers.StringSplitNewlinesNewlinesStrings(input)
	if len(data) <= 2 {
		return fmt.Errorf("unable to parse input")
	}

	var err error
	d.numbersDrawn, err = d.parseNumbersDrawn(data[0])
	if err != nil {
		return err
	}

	d.boards, err = bingo.ParseBoards(data[1:])
	if err != nil {
		return err
	}

	return nil
}

// 11312 --> too high

func (d *Day04) Part1() error {
	for _, num := range d.numbersDrawn {
		winner := d.boards.Mark(num)
		if winner != nil {
			sum, err := d.boards.SumBoardRemaining(*winner)
			if err != nil {
				return err
			}

			d.winningNumber = &num
			d.winningBoardNumber = winner
			d.winningBoardSumRemaining = sum

			return nil
		}
	}

	return nil
}

func (d *Day04) Part2() error {

	return nil
}

func (d *Day04) parseNumbersDrawn(s string) (pie.Ints, error) {
	if d.input == nil {
		return nil, fmt.Errorf("empty input value, unable to parse numbers drawn")
	}

	ints := pie.Ints{}
	ss := strings.Split(s, ",")
	for _, is := range ss {
		i, err := strconv.Atoi(is)
		if err != nil {
			return nil, fmt.Errorf("unable to parse value in numbers drawn '%s': %w", is, err)
		}
		ints = ints.Append(i)
	}

	return ints, nil
}
