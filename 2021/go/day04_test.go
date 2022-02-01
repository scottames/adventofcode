package aoc2021

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/2021/go/pkg/bingo"
	"github.com/stretchr/testify/assert"
)

const (
	day04Example string = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`
)

var (
	exampleBoards = bingo.Boards{
		0: &bingo.Board{
			Number: 0,
			Marks:  pie.Ints{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24},
			Rows: []*bingo.Row{
				{Numbers: pie.Ints{22, 13, 17, 11, 0}, Marked: 3},
				{Numbers: pie.Ints{8, 2, 23, 4, 24}, Marked: 4},
				{Numbers: pie.Ints{21, 9, 14, 16, 7}, Marked: 4},
				{Numbers: pie.Ints{6, 10, 3, 18, 5}, Marked: 1},
				{Numbers: pie.Ints{1, 12, 20, 15, 19}, Marked: 0},
			},
		},
		1: &bingo.Board{
			Number: 1,
			Marks:  pie.Ints{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24},
			Rows: []*bingo.Row{
				{Numbers: pie.Ints{3, 15, 0, 2, 22}, Marked: 2},
				{Numbers: pie.Ints{9, 18, 13, 17, 5}, Marked: 3},
				{Numbers: pie.Ints{19, 8, 7, 25, 23}, Marked: 2},
				{Numbers: pie.Ints{20, 11, 10, 24, 4}, Marked: 3},
				{Numbers: pie.Ints{14, 21, 16, 12, 6}, Marked: 2},
			},
		},
		2: &bingo.Board{
			Number: 2,
			Marks:  pie.Ints{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24},
			Rows: []*bingo.Row{
				{Numbers: pie.Ints{14, 21, 17, 24, 4}, Marked: 5},
				{Numbers: pie.Ints{10, 16, 15, 9, 19}, Marked: 1},
				{Numbers: pie.Ints{18, 8, 23, 26, 20}, Marked: 1},
				{Numbers: pie.Ints{22, 11, 13, 6, 5}, Marked: 2},
				{Numbers: pie.Ints{2, 0, 12, 3, 7}, Marked: 3},
			},
		},
	}
	exampleNumbersDrawn = pie.Ints{
		7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1,
	}
	exampleWinningBoardNumber       = 2
	exampleWinningNumber            = 24
	exampleWinningBoardSumRemaining = 188
	testDay04                       = Day04{
		input:                    []byte(day04Example),
		boards:                   exampleBoards,
		numbersDrawn:             exampleNumbersDrawn,
		winningBoardNumber:       &exampleWinningBoardNumber,
		winningBoardSumRemaining: &exampleWinningBoardSumRemaining,
		winningNumber:            &exampleWinningNumber,
	}
)

func Test_Day04(t *testing.T) {
	expectedDay04 := testDay04
	actualDay04 := Day04{}
	err := actualDay04.Parse([]byte(day04Example))
	if err != nil {
		t.Errorf(err.Error())
	}

	err = actualDay04.Part1()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = actualDay04.Part2()
	if err != nil {
		t.Errorf(err.Error())
	}

	msg := fmt.Sprintf("Expected %v. Got %v.", expectedDay04, actualDay04)
	assert.Equal(t, expectedDay04, actualDay04, msg)

	expectedScore := 4512
	actualScore := *actualDay04.winningBoardSumRemaining * *actualDay04.winningNumber

	msg = fmt.Sprintf("Expected %v. Got %v.", expectedScore, actualScore)
	assert.Equal(t, expectedScore, actualScore, msg)
}
