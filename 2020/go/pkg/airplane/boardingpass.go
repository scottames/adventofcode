package airplane

import (
	"github.com/elliotchance/pie/pie"
)

type BoardingPass struct {
	Column int
	Row    int
	Pass   string
	Seat   int
}

// NewBoardingPasses returns a slice of BoardingPass and their respective seat IDs
// from a given slice of strings
//
// Each string expected to be in the format like 'FBFBBFFRLR'
// where F means "front", B means "back", L means "left", and R means "right".
func NewBoardingPasses(ss pie.Strings) (passes []BoardingPass, seats pie.Ints) {
	passes = []BoardingPass{}
	seats = pie.Ints{}

	ss.Map(func(s string) string {
		if s == "" {
			return s
		}
		pass := NewBoardingPass(s)
		passes = append(passes, pass)
		if pass.Seat > 0 {
			seats = seats.Append(pass.Seat)
		}
		return s
	})

	return passes, seats.Sort()
}

// NewBoardingPass returns a BoardingPass from a given string
//
// String expected to be in the format like 'FBFBBFFRLR'
// where F means "front", B means "back", L means "left", and R means "right".
func NewBoardingPass(s string) BoardingPass {
	bp := BoardingPass{
		Pass: s,
	}
	bp.Row = bp.calculateRow()
	bp.Column = bp.calculateColumn()
	bp.Seat = bp.calculateSeat()
	return bp
}

// Columns returns the column string from the BoardingPass string
func (bp BoardingPass) Columns() string {
	if len(bp.Pass) != 10 {
		return ""
	}
	return string([]rune(bp.Pass)[7:])
}

// Rows returns the column string from the BoardingPass string
func (bp BoardingPass) Rows() string {
	if len(bp.Pass) != 10 {
		return ""
	}
	return string([]rune(bp.Pass)[:7])
}

// calculateColumn calculates the Column ID for the given BoardingPass
func (bp BoardingPass) calculateColumn() int {
	return halving(bp.Columns(), 7)
}

// calculateRow calculates the Row ID for the given BoardingPass
func (bp BoardingPass) calculateRow() int {
	return halving(bp.Rows(), 127)
}

// calculateSeat calculates the Seat for the given BoardingPass
func (bp BoardingPass) calculateSeat() int {
	return bp.calculateRow()*8 + bp.calculateColumn()
}

func halving(s string, num int) int {
	top := num
	bottom := 0
	for i, c := range s {
		middle := (top + bottom) / 2
		switch string(c) {
		case "B", "R":
			bottom = middle + 1
			if i == len(s)-1 {
				return top
			}
		case "F", "L":
			top = middle
			if i == len(s)-1 {
				return bottom
			}
		}
	}
	return 0
}
