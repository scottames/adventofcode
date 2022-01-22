package types

type Dayer interface {
	Parse([]byte) error
	String() string

	Part1() error
	Part2() error
}
