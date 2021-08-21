package seats

import (
	"reflect"

	"github.com/elliotchance/pie/pie"
)

const (
	floor = iota
	empty
	occupied
)

const (
	floorChar    = '.'
	emptyChar    = 'L'
	occupiedChar = '#'
)

type Arrangement struct {
	previous *Arrangement
	set      []pie.Ints
	occupied int
}

func NewArrangement(ss pie.Strings) *Arrangement {
	set, occupied := parseStringToSet(ss)
	return &Arrangement{
		set:      set,
		occupied: occupied,
	}
}

func parseStringToSet(ss pie.Strings) ([]pie.Ints, int) {
	set := newEmptySet(len(ss), len(ss.First()))
	occupied := 0

	for r, rc := range ss {
		for s, sc := range rc {
			switch sc {
			case emptyChar:
				set[r][s] = empty
			case floorChar:
				set[r][s] = floor
			case occupiedChar:
				set[r][s] = occupied
				occupied += 1
			}
		}
	}
	return set, occupied
}

func (self *Arrangement) rowLen() int {
	if self == nil {
		return 0
	}
	return len(self.set)
}

func (self *Arrangement) columnLen() int {
	if self == nil || self.rowLen() == 0 {
		return 0
	}
	return self.set[0].Len()
}

func (self *Arrangement) Next(fn func(*Arrangement) ([]pie.Ints, int)) *Arrangement {
	nextSet, setOccupancy := fn(self)
	return &Arrangement{
		previous: self,
		set:      nextSet,
		occupied: setOccupancy,
	}
}

func (self *Arrangement) NextUntilMatchingPart1() *Arrangement {
	return self.NextUntilMatching(part1SeatingLogic)
}

func (self *Arrangement) NextUntilMatchingPart2() *Arrangement {
	return self.NextUntilMatching(part2SeatingLogic)
}

func (self *Arrangement) NextUntilMatching(fn func(*Arrangement) ([]pie.Ints, int)) *Arrangement {
	arrangement := self
	for ; ; arrangement = arrangement.Next(fn) {
		if arrangement.EqualTo(arrangement.Previous()) {
			return arrangement
		}
	}
}

func (self *Arrangement) EqualTo(arr *Arrangement) bool {
	if arr == nil && self != nil {
		return false
	}
	return reflect.DeepEqual(self.set, arr.set)
}

func (self *Arrangement) Current() *Arrangement {
	return self
}

func (self *Arrangement) Previous() *Arrangement {
	return self.previous
}

func (self *Arrangement) OccupiedSeats() int {
	return self.occupied
}

func newEmptySet(rows, cols int) []pie.Ints {
	r := make([]pie.Ints, rows)
	for i := range r {
		r[i] = make(pie.Ints, cols)
	}
	return r
}

func part1SeatingLogic(self *Arrangement) ([]pie.Ints, int) {
	set := newEmptySet(self.rowLen(), self.columnLen())
	totalOccupancy := 0
	for ri, seats := range self.set {

		// range over rows arround the given seat
		rr := [3]int{ri - 1, ri, ri + 1}
		for si, seat := range seats {

			if seat == floor {
				continue
			}

			// range over seats around the given seat
			sr := [3]int{si - 1, si, si + 1}
			occupiedCount := 0
			// count number of surrounding occupied seats by ranging over the nearby rows + seats
			for _, r := range rr {
				for _, s := range sr {
					if r == ri && s == si ||
						r < 0 ||
						s < 0 ||
						r >= self.rowLen() ||
						s >= self.columnLen() {
						continue
					}
					// is the surrounding seat occupied or empty?
					if self.set[r][s] == occupied {
						occupiedCount++
					}
				}
			}
			switch seat {
			case empty:
				if occupiedCount == 0 {
					set[ri][si] = occupied
					totalOccupancy += 1
				} else {
					set[ri][si] = empty
				}
			case occupied:
				if occupiedCount >= 4 {
					set[ri][si] = empty
				} else {
					set[ri][si] = occupied
					totalOccupancy += 1
				}
			}
		}
	}
	return set, totalOccupancy
}

func part2SeatingLogic(self *Arrangement) ([]pie.Ints, int) {
	set := newEmptySet(self.rowLen(), self.columnLen())
	totalOccupancy := 0

	for ri, seats := range self.set {
		for si, seat := range seats {

			if seat == floor {
				continue
			}

			occupiedCount := 0
			for _, fn := range []func(int, int) bool{
				self.rightOccupied,
				self.leftOccupied,
				self.upOccupied,
				self.downOccupied,
				self.upLeftOccupied,
				self.upRightOccupied,
				self.downLeftOccupied,
				self.downRightOccupied,
			} {
				if fn(ri, si) {
					occupiedCount++
				}
			}

			switch seat {
			case empty:
				if occupiedCount == 0 {
					set[ri][si] = occupied
					totalOccupancy += 1
				} else {
					set[ri][si] = empty
				}
			case occupied:
				if occupiedCount >= 5 {
					set[ri][si] = empty
				} else {
					set[ri][si] = occupied
					totalOccupancy += 1
				}
			}
		}
	}
	return set, totalOccupancy
}

func (self *Arrangement) rightOccupied(ri, si int) bool {
	for s := si + 1; s < self.columnLen(); s++ {
		switch self.set[ri][s] {
		case floor:
			continue
		case empty:
			return false
		case occupied:
			return true
		}
	}
	return false
}

func (self *Arrangement) leftOccupied(ri, si int) bool {
	for s := si - 1; s >= 0; s-- {
		switch self.set[ri][s] {
		case floor:
			continue
		case empty:
			return false
		case occupied:
			return true
		}
	}
	return false
}

func (self *Arrangement) upOccupied(ri, si int) bool {
	for r := ri - 1; r >= 0; r-- {
		switch self.set[r][si] {
		case floor:
			continue
		case empty:
			return false
		case occupied:
			return true
		}
	}
	return false
}

func (self *Arrangement) downOccupied(ri, si int) bool {
	for r := ri + 1; r < self.rowLen(); r++ {
		switch self.set[r][si] {
		case floor:
			continue
		case empty:
			return false
		case occupied:
			return true
		}
	}
	return false
}

func (self *Arrangement) upRightOccupied(ri, si int) bool {
	for r, s := ri+1, si+1; r < self.rowLen() && s < self.columnLen(); r, s = r+1, s+1 {
		switch self.set[r][s] {
		case floor:
			continue
		case empty:
			return false
		case occupied:
			return true
		}
	}
	return false
}

func (self *Arrangement) downRightOccupied(ri, si int) bool {
	for r, s := ri-1, si+1; r >= 0 && s < self.columnLen(); r, s = r-1, s+1 {
		switch self.set[r][s] {
		case floor:
			continue
		case empty:
			return false
		case occupied:
			return true
		}
	}
	return false
}

func (self *Arrangement) upLeftOccupied(ri, si int) bool {
	for r, s := ri-1, si-1; s >= 0 && r >= 0; r, s = r-1, s-1 {
		switch self.set[r][s] {
		case floor:
			continue
		case empty:
			return false
		case occupied:
			return true
		}
	}
	return false
}

func (self *Arrangement) downLeftOccupied(ri, si int) bool {
	for r, s := ri+1, si-1; r < self.rowLen() && s >= 0; r, s = r+1, s-1 {
		switch self.set[r][s] {
		case floor:
			continue
		case empty:
			return false
		case occupied:
			return true
		}
	}
	return false
}
