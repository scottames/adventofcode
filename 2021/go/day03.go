package aoc2021

import (
	"fmt"
	"strconv"

	"github.com/elliotchance/pie/pie"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/scottames/adventofcode/pkg/helpers"
	"github.com/scottames/adventofcode/pkg/style"
)

type Day03 struct {
	input      []byte
	report     pie.Strings
	colLen     int
	gamma      pie.Strings
	gammaInt   int64
	epsilon    pie.Strings
	epsilonInt int64
	o2rating   int64
	co2rating  int64
}

func init() {
	Days[3] = &Day03{}
}

func (d *Day03) String() string {
	p1 := NewTableRounded()
	p1.AppendHeader(table.Row{"Part 1"})
	p1.AppendRows([]table.Row{
		{style.Blue("gamma"), d.gamma.Join("")},
		{style.Magenta("epsilon"), d.epsilon.Join("")},
	})
	p1.AppendFooter(table.Row{
		style.Red("power consumption"), d.gammaInt * d.epsilonInt,
	})

	p2 := NewTableRounded()
	p2.AppendHeader(table.Row{"Part 2"})
	p2.AppendRows([]table.Row{
		{style.Blue("o2 rating"), d.o2rating},
		{style.Magenta("co2 rating"), d.co2rating},
	})
	p2.AppendFooter(table.Row{
		style.Red("CO2 scrubber rating"), d.o2rating * d.co2rating,
	})

	return p1.Render() + "\n\n" + p2.Render()
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

	d.gammaInt, err = strconv.ParseInt(d.gamma.Join(""), 2, 64)
	if err != nil {
		return err
	}

	d.calcEpsilon()
	d.epsilonInt, err = strconv.ParseInt(d.epsilon.Join(""), 2, 64)
	if err != nil {
		return err
	}

	return nil
}

func (d *Day03) Part2() error {
	var err error
	d.o2rating, err = strconv.ParseInt(d.rating(commons{"1", "0"}), 2, 64)
	if err != nil {
		return err
	}

	d.co2rating, err = strconv.ParseInt(d.rating(commons{"0", "1"}), 2, 64)
	if err != nil {
		return err
	}

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
