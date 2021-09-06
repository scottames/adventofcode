package port

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	clear = iota
	set
)

func newMemory() Memory {
	return make(Memory)
}

type Memory map[int]value

func (self Memory) Sum() int {
	sum := 0
	for _, v := range self {
		sum += v.int()
	}

	return sum
}

func (self Memory) setAddressString(s string, v value) error {
	i, err := extractMemAddr(s)
	if err != nil {
		return err
	}
	self.setAddress(i, v)

	return nil
}

func (self Memory) setAddress(i int, v value) {
	self[i] = v
}

func extractMemAddr(s string) (int, error) {
	re := regexp.MustCompile(`\[(.*?)\]`)
	match := re.FindString(s)
	return strconv.Atoi(strings.Trim(strings.Trim(match, "["), "]"))
}

func newMask(s string) mask {
	result := make(mask)
	for i, c := range s {
		n, err := strconv.Atoi(string(c))
		ri := len(s) - 1 // reverse index because we want to read the mask right to left
		if err != nil {
			continue
		}
		result.set(uint64(ri-i), uint64(n))
	}
	return result
}

type mask map[uint64]uint64 // position: set || clear

func (self mask) set(position uint64, value uint64) {
	self[position] = value
}

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

type value uint64

func (self value) int() int {
	return int(self)
}

// TODO -> document these to be able to explain!

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
