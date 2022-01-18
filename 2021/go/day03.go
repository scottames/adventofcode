package aoc2021

import (
	"fmt"
	"strconv"

	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

type Day03 struct {
	input     []byte
	report    pie.Strings
	colLen    int
	gamma     pie.Strings
	epsilon   pie.Strings
	o2rating  int
	co2rating int
}

func init() {
	Days[3] = &Day03{}
}

func (d *Day03) Parse(input []byte) error {
	d.input = input
	d.report = helpers.StringSplitNewlinesStrings(d.input)
	if len(d.report) == 0 {
		return fmt.Errorf("input of zero length found")
	}

	d.colLen = len(d.report[0])

	return nil
}

func (d *Day03) Part1() error {
	err := d.calcGamma()
	if err != nil {
		return err
	}

	gamma, err := d.binaryStringToInt(d.gamma.Join(""))
	if err != nil {
		return err
	}

	d.calcEpsilon()
	epsilon, err := d.binaryStringToInt(d.epsilon.Join(""))
	if err != nil {
		return err
	}

	fmt.Println(d.gamma)
	fmt.Println(d.epsilon)

	fmt.Println("  gamma:   ", gamma)
	fmt.Println("  epsilon: ", epsilon)
	fmt.Println("  ---")
	fmt.Println("  power consumption: ", gamma*epsilon)

	return nil
}

func (d *Day03) Part2() error {
	var err error
	d.o2rating, err = d.binaryStringToInt(d.rating(commons{"1", "0"}))
	if err != nil {
		return err
	}

	d.co2rating, err = d.binaryStringToInt(d.rating(commons{"0", "1"}))
	if err != nil {
		return err
	}

	fmt.Println("  o2 rating:  ", d.o2rating)
	fmt.Println("  co2 rating: ", d.co2rating)
	fmt.Println("  ---")
	fmt.Println("  CO2 scrubber rating: ", d.o2rating*d.co2rating)

	return nil
}

func colTotals(ss pie.Strings) (map[int]int, error) {
	result := make(map[int]int)
	valid := pie.Strings{"0", "1"}

	for lineNum, s := range ss {
		for colNum, col := range s {
			if !valid.Contains(string(col)) {
				return nil, fmt.Errorf("invalid bit ('%s') at %d:%d (line:col)", string(col), lineNum, colNum)
			}
			i, _ := strconv.Atoi(string(col))
			result[colNum] += i
		}
	}

	return result, nil
}

func (d *Day03) calcGamma() error {
	d.gamma = make(pie.Strings, d.colLen)

	ct, err := colTotals(d.report)
	if err != nil {
		return err
	}

	for col, sum := range ct {
		if sum > d.report.Len()/2 {
			d.gamma[col] = "1"
		} else {
			d.gamma[col] = "0"
		}
	}

	return nil
}

func (d *Day03) calcEpsilon() {
	d.epsilon = make(pie.Strings, d.colLen)

	for col, num := range d.gamma {
		var result string
		if num == "1" {
			result = "0"
		} else {
			result = "1"
		}
		d.epsilon[col] = result
	}
}

func (d *Day03) binaryStringToInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse string slice '%s': %w", s, err)
	}
	return int(i), nil
}

type commons struct {
	common   string
	uncommon string
}

func (d *Day03) rating(c commons) string {
	result := make(pie.Strings, len(d.report))
	copy(result, d.report)

	for i := 0; i < d.colLen; i++ {
		if len(result) == 1 {
			break
		}

		bitsAtCol := result.Map(func(s string) string {
			return string(s[i])
		})
		var commonBit string
		lr := len(result)
		if float64(bitsAtCol.Ints().Sum()) >= float64(lr)/2 {
			commonBit = c.common
		} else {
			commonBit = c.uncommon
		}

		result = result.Filter(func(s string) bool {
			return string(s[i]) == commonBit
		})
	}

	return result[0]
}
