package jolts

import (
	"fmt"

	"github.com/elliotchance/pie/pie"
)

// Chain represents the number of diffs in a given chain of adapters
type Chain struct {
	adapters     pie.Ints
	arrangements pie.Ints
	ordered      pie.Ints
	diffs        map[int]int // [differences]count
}

// NewChain returns a new Chain from a given slice of ints
func NewChain(adapters pie.Ints) *Chain {
	if adapters.Len() == 0 {
		return &Chain{}
	}

	ordered := orderAdapters(adapters)
	chain := Chain{
		adapters:     adapters,
		arrangements: initArrangements(len(ordered)),
		diffs:        make(map[int]int),
		ordered:      ordered,
	}

	return chain.connectBottomUp()
}

// connectBottomUp connects all adapters and calculates the differences between them
// and the number of distinct arrangements for each adapter
func (self *Chain) connectBottomUp() *Chain {
	for i := 1; i < len(self.ordered); i++ {
		self.diffs[self.diffFromLast(i)]++
		self.arrangements[i] = self.arrangementSum(i)
	}
	return self
}

// diffFromLast returns the difference in "jolts" from the last adapter in the chain
func (self *Chain) diffFromLast(index int) int {
	return self.ordered[index] - self.ordered[index-1]
}

// arrangementSum calculates the sum of the previous valid (difference of 1-3)
// adapters' arrangements using a bottom-up approach - or in other words,
// the number of paths of the previous valid adapters that can be connected
// Example:
//
//  0  1  4  5  6   7   10  11  12  15  16  19  22
//  1  1  1  1  2   4   4   4   8   8   8   8   8
//              |   |           |
//             [4] [4]         [10]
//              +   +           +
//             [5] [5]         [11]
//                  +
//                 [6]
//
func (self *Chain) arrangementSum(index int) int {
	sum := 0
	for n := 1; n <= 3; n++ {
		if index-n < 0 {
			continue
		} else if self.ordered[index]-self.ordered[index-n] <= 3 {
			sum += self.arrangements[index-n]
		}
	}
	return sum
}

// Multiples returns the product of 1 & 3 differences found in the given chain
func (self Chain) Multiples() (int, error) {
	ones, ok := self.diffs[1]
	if !ok {
		return 0, fmt.Errorf("no instances of 1 found in chain diffs")
	}
	threes, ok := self.diffs[3]
	if !ok {
		return 0, fmt.Errorf("no instances of 3 found in chain diffs")
	}
	return ones * threes, nil
}

func (self Chain) DistinctArrangements() int {
	return self.arrangements.Last()
}

func (self Chain) Diffs() map[int]int {
	return self.diffs
}

// initArrangements returns a slice of ints with the desired length
// and the first index initialized to 1
func initArrangements(length int) pie.Ints {
	a := make(pie.Ints, length)
	a[0] = 1
	return a
}

// orderAdapters returns a sorted version of the given slice of ints
// with the first and last connectors prepended (0) & appended (len(slice)+3)
func orderAdapters(is pie.Ints) pie.Ints {
	ordered := is.Sort().Insert(0, 0)
	return ordered.Append(ordered[len(is)] + 3)
}
