package aoc2021

type Dayer interface {
	Parse([]byte) error
	Part1() error
	Part2() error
}

var Days map[int]Dayer

func init() {
	Days = make(map[int]Dayer)
}
