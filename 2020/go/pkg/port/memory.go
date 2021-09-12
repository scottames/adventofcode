package port

import (
	"fmt"
	"strconv"
)

const (
	clear uint64 = iota
	set
	floating
)

type memory map[value]value

func (self memory) sum() int {
	sum := 0
	for _, v := range self {
		sum += v.int()
	}

	return sum
}

func (self memory) setValue(index value, val value) {
	self[index] = val
}

func extractMemAddr(s string) (int, error) {
	var i int
	_, err := fmt.Sscanf(s, "mem[%d]", &i)
	if err != nil {
		return 0, fmt.Errorf("unable to process mem address")
	}
	return i, nil
}

func newMask(s string) mask {
	result := make(mask)
	for i, c := range s {
		ri := len(s) - 1 // reverse index because we want to read the mask right to left
		n, err := strconv.ParseUint(string(c), 10, 64)
		if err != nil {
			result.set(uint64(ri-i), floating)
		} else {
			result.set(uint64(ri-i), n)
		}
	}
	return result
}

type mask map[uint64]uint64 // position: set || clear || floating

func (self mask) set(position uint64, value uint64) {
	self[position] = value
}

// TODO: merge these into an interface
func (self mask) apply(i value) value {
	result := i
	for pos, action := range self {
		if action == clear {
			result = result.clearBit(pos)
		} else if action == set {
			result = result.setBit(pos)
		}
	}
	return result
}

func (self mask) applyFloating(val value) []value {
	// The entire 36-bit address space still begins initialized to the value 0 at every address
	result := []value{0}

	for pos, action := range self {
		switch action {
		// If the bitmask bit is 0, the corresponding memory address bit is unchanged.
		case clear:
			for i := range result {
				if val.getBit(pos) == 1 {
					result[i] = result[i].setBit(pos)
				}
			}
		// If the bitmask bit is 1, the corresponding memory address bit is overwritten with 1.
		case set:
			for i := range result {
				result[i] = result[i].setBit(pos)
			}
		// If the bitmask bit is X, the corresponding memory address bit is floating.
		case floating:
			for i := range result {
				result = append(result, result[i])
				result[i] = result[i].setBit(pos)
			}
		}
	}
	return result
}

type value uint64

func (self value) int() int {
	return int(self)
}

func (self value) setBit(pos uint64) value {
	return self.setBits(1 << pos)
}

func (self value) setBits(bits value) value {
	return self | bits
}

func (self value) clearBit(pos uint64) value {
	return self.clearBits(1 << pos)
}

func (self value) clearBits(bits value) value {
	return self & ^bits
}

func (self value) getBit(pos uint64) value {
	mask := value(1) << pos
	return (self & mask) >> pos
}
