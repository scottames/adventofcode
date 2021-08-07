package xmas

import (
	"github.com/elliotchance/pie/pie"
)

func FindFirstBreak(input pie.Ints, preamble int) int {
	if input.Len() <= preamble {
		return 0
	}

	for set := input.Top(preamble); input.Len() > preamble; set = input.Top(preamble) {
		val := input[preamble]

		if filter := set.Filter(func(i int) bool {
			for _, z := range set {
				if i+z == val {
					return true
				}
			}
			return false
		}); filter.Len() < 2 {
			return val
		}

		input.Pop()
	}

	return 0
}

func FindContiguousSetWithSum(is pie.Ints, num int) pie.Ints {
	for i := range is {
		for j := i + 1; j < len(is); j++ {
			try := is.SubSlice(i, j) // is[i:j]
			switch sum := try.Sum(); {
			case sum == num:
				return try
			case sum > num:
				break
			}
		}
	}
	return pie.Ints{}
}

func SumSmallestAndLargestInts(is pie.Ints) int {
	sorted := is.Sort()
	return sorted.First() + sorted.Last()
}
