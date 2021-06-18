package aoc2019

import (
	"fmt"
	"github.com/elliotchance/pie/pie"
	"math"
)

// Day04 - https://adventofcode.com/2019/day/4
type Day04 struct {
	Beginning int
	End       int
}

func (self Day04) Part1() int {
	return self.calculate()
}

func (self Day04) calculate() int {
	//for i := self.Beginning; i <= self.End; i++ {
	//	double := false
	list := make(pie.Ints, 5, 5)
	for j := 0; j < 5; j++ {
		list[j] = digit(self.Beginning, j)
	}
	fmt.Printf("%#v", list)
	//}
	return 0
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
