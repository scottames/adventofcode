package jolts

import (
	"fmt"

	"github.com/elliotchance/pie/pie"
)

// ChainDIffs represents the number of diffs in a given chain of adapters
type ChainDiffs map[int]int // [differences]count

// NewChainDiffs returns a new ChainDiffs from a given slice of ints
func NewChainDiffs(adapters pie.Ints) ChainDiffs {
	sorted := adapters.Sort()
	chainDiffs := make(ChainDiffs)
	last := 0
	for _, adapter := range sorted {
		diff := adapter - last
		chainDiffs[diff]++
		last = adapter
	}
	// Finally, your device's built-in adapter is always 3 higher than the highest
	//  adapter, so its rating is 22 jolts (always a difference of 3).
	chainDiffs[3]++
	return chainDiffs
}

// Multiples returns the product of 1 & 3 differences found in the given chainDiffs
func (self ChainDiffs) Multiples() (int, error) {
	ones, ok := self[1]
	if !ok {
		return 0, fmt.Errorf("no instances of 1 found in chain diffs")
	}
	threes, ok := self[3]
	if !ok {
		return 0, fmt.Errorf("no instances of 3 found in chain diffs")
	}
	return ones * threes, nil
}
