package bingo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elliotchance/pie/pie"
)

const (
	BoardSize int = 5
)

type Row struct {
	Numbers pie.Ints
	Marked  int
}

func (r *Row) Append(number int) error {
	if len(r.Numbers) >= BoardSize {
		return fmt.Errorf(
			"row '%v' full (%d), unable to append '%d'", r.Numbers, BoardSize, number,
		)
	}

	r.Numbers = r.Numbers.Append(number)

	return nil
}

type Board struct {
	Marks  pie.Ints
	Number int
	Rows   []*Row
}

func NewBoard() *Board {
	return &Board{}
}

func ParseBoard(s string) (*Board, error) {
	board := NewBoard()

	for _, row := range strings.Split(s, "\n") {
		err := board.AppendRow(row)
		if err != nil {
			return nil, err
		}
	}

	return board, nil
}

func (board *Board) AppendRow(row string) error {
	if len(board.Rows) >= BoardSize {
		return fmt.Errorf(
			"board #%d: rows full (%d), unable to append '%s'",
			board.Number,
			BoardSize,
			row,
		)
	}

	r := Row{}
	for _, s := range strings.Fields(row) {
		squareNumber, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		err = r.Append(squareNumber)
		if err != nil {
			return err
		}
	}

	board.Rows = append(board.Rows, &r)

	return nil
}

func (board *Board) Mark(number int) *int {
	for _, row := range board.Rows {
		if row.Numbers.Contains(number) {
			row.Marked++
			board.Marks = board.Marks.Append(number)
			if row.Marked >= BoardSize {
				winningBoard := board.Number
				return &winningBoard
			}
		}
	}

	return nil
}

func (board *Board) Sum() int {
	sum := 0
	for _, row := range board.Rows {
		sum += row.Numbers.Sum()
	}
	return sum
}

func (board *Board) SumRemaining() int {
	return board.Sum() - board.Marks.Sum()
}

type Boards map[int]*Board

func NewBoards() Boards {
	return make(Boards)
}

func ParseBoards(boardsAsStringSlice []string) (Boards, error) {
	boards := NewBoards()
	for i, boardString := range boardsAsStringSlice {
		board, err := ParseBoard(boardString)
		if err != nil {
			return nil, err
		}

		board.Number = i
		err = boards.Add(board)
		if err != nil {
			return nil, err
		}
	}

	return boards, nil
}

func (boards Boards) Add(board *Board) error {
	if _, ok := boards[board.Number]; ok {
		return fmt.Errorf("board number '%d' already exists in boards", board.Number)
	}

	boards[board.Number] = board

	return nil
}

func (boards Boards) Mark(number int) *int {
	for i := 0; i < len(boards); i++ {
		board := boards[i]
		winner := board.Mark(number)
		if winner != nil {
			return winner
		}
	}

	return nil
}

func (boards Boards) SumBoardRemaining(number int) (*int, error) {
	board, ok := boards[number]
	if !ok {
		return nil, fmt.Errorf("board #'%d' not found", number)
	}

	remainingSum := board.SumRemaining()
	return &remainingSum, nil
}
