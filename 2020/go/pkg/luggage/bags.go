package luggage

import (
	"strconv"
	"strings"

	"github.com/elliotchance/pie/pie"
)

// NewBag returns a new Bag from the given bag description
func NewBag(s string) Bag {
	z := strings.Split(s, "bags contain")

	description, rulesString := z[0], z[1]
	bag := Bag{
		Description: strings.TrimSpace(description),
	}
	holds := bag.holds(rulesString)
	bag.HoldsBags = holds

	return bag
}

// NewBags returns a new Bags from a given
// rule set of bags (problem input)
func NewBags(ss []string) Bags {
	bags := Bags{}
	for _, s := range ss {
		bag := NewBag(s)
		bags[bag.Description] = bag
	}
	return bags
}

// Bags represents a mapping (list/set) of Bag
type Bags map[string]Bag

// Bag represents a bag and the subsequent
// number and name of bags it can hold
type Bag struct {
	Description string
	HoldsBags   map[string]int
}

// NumBagsCanContainBag - returns the number of bags
// that a given bag can contain (Part 1)
func (self Bags) NumBagsCanContainBag(bag string) int {
	list := pie.Strings{bag}
	for n := list; n != nil; n = self.BagsFitIn(n) {
		if n != nil {
			list = list.Append(n...)
		}
	}
	return list.Unique().Filter(func(s string) bool {
		return s != bag
	}).Len()
}

// BagsFitIn - given a list of bag names returns a list of
// bags that each of the bags fit in
func (self Bags) BagsFitIn(bags pie.Strings) pie.Strings {
	var list pie.Strings
	for _, bag := range self {
		for b := range bag.HoldsBags {
			if bags.Contains(b) {
				list = list.Append(bag.Description)
			}
		}
	}
	return list
}

// NumBagHolds - returns the number of bags a
// given bag holds (Part 2)
func (self Bags) NumBagHolds(bag string) int {
	var thisBag Bag
	var ok bool
	if thisBag, ok = self[bag]; !ok {
		return 0
	}
	count := 0
	for b, n := range thisBag.HoldsBags {
		count += n + (n * self.NumBagHolds(b))
	}
	return count
}

// holds returns a map of name:int from a given
// bag rule description
func (self Bag) holds(s string) map[string]int {
	if strings.Contains(s, "no other bags") {
		return map[string]int{}
	}

	rs := strings.Split(s, ",")
	holds := map[string]int{}
	for _, r := range rs {
		ss := strings.Split(strings.TrimSpace(r), " ")
		num, err := strconv.Atoi(ss[0])

		// TODO: return error... in the real world this should be handled
		// for time's sake, we'll assume no error will occur with our data
		if err != nil {
			return map[string]int{}
		}

		desc := ss[1] + " " + ss[2]
		holds[desc] = num
	}
	return holds
}
