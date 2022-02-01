package bingo

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/stretchr/testify/assert"
)

var (
	exampleBoard0 = &Board{
		Marks:  nil,
		Number: 0,
		Rows: []*Row{
			{
				Numbers: pie.Ints{22, 13, 17, 11, 0},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{8, 2, 23, 4, 24},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{21, 9, 14, 16, 7},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{6, 10, 3, 18, 5},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{1, 12, 20, 15, 19},
				Marked:  0,
			},
		},
	}

	exampleBoard1 = &Board{
		Marks:  nil,
		Number: 0,
		Rows: []*Row{
			{
				Numbers: pie.Ints{3, 15, 0, 2, 22},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{9, 18, 13, 17, 5},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{19, 8, 7, 25, 23},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{20, 11, 10, 24, 4},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{14, 21, 16, 12, 6},
				Marked:  0,
			},
		},
	}

	exampleBoard2 = &Board{
		Marks:  nil,
		Number: 0,
		Rows: []*Row{
			{
				Numbers: pie.Ints{14, 21, 17, 24, 4},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{10, 16, 15, 9, 19},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{18, 8, 23, 26, 20},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{22, 11, 13, 6, 5},
				Marked:  0,
			},
			{
				Numbers: pie.Ints{2, 0, 12, 3, 7},
				Marked:  0,
			},
		},
	}

	exampleTestBoards = map[string]*Board{
		`22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19`: exampleBoard0,
		` 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6`: exampleBoard1,
		`14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`: exampleBoard2,
	}

	exampleTestBoardsSum = map[int]*Board{
		300: exampleBoard0,
		324: exampleBoard1,
		325: exampleBoard2,
	}
)

func Test_ParseBoard(t *testing.T) {
	for string, expected := range exampleTestBoards {
		actual, err := ParseBoard(string)
		if err != nil {
			t.Errorf(err.Error())
		}

		msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
		assert.Equal(t, expected, actual, msg)
	}
}

func Test_BoardSum(t *testing.T) {
	for expected, board := range exampleTestBoardsSum {
		actual := board.Sum()

		msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
		assert.Equal(t, expected, actual, msg)
	}
}
