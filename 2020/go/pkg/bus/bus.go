package bus

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

type Bus struct {
	id     int
	offset int
}

type Buses struct {
	list []Bus
	min  int
	max  int
}

func (self Buses) Max() int {
	return self.max
}

func (self Buses) Min() int {
	return self.min
}

func (self Buses) List() pie.Ints {
	ints := pie.Ints{}
	for _, bus := range self.list {
		ints = ints.Append(bus.id)
	}
	return ints
}

type Schedule struct {
	earliestDeparture int
	buses             *Buses
	nextDeparture     *int
	nextDepartureID   *int
}

func (self *Schedule) EarliestDeparture() int {
	return self.earliestDeparture
}

func (self *Schedule) NextDeparture() int {
	if self.nextDeparture != nil {
		return *self.nextDeparture
	}
	return *self.calcNextDeparture().nextDeparture
}

func (self *Schedule) calcNextDeparture() *Schedule {
	next := self.earliestDeparture
	for {
		for _, bus := range self.buses.list {
			if next%bus.id == 0 {
				self.nextDeparture = &next
				self.nextDepartureID = &bus.id
				return self
			}
		}
		next++
	}
}

func (self *Schedule) NextDepartureID() int {
	if self.nextDepartureID != nil {
		return *self.nextDepartureID
	}
	return *self.calcNextDeparture().nextDepartureID
}

// EarliestTimestampOffsetAlignment returns the earliest timestamp
// which the offset aligns for the given Schedule
func (self *Schedule) EarliestTimestampOffsetAlignment() int {
	step := self.buses.list[0].id
	timestamp := 0

	for _, bus := range self.buses.list {
		if bus.id == step {
			continue
		}

		for (timestamp+bus.offset)%bus.id != 0 {
			timestamp += step
		}

		step = helpers.LCM(step, bus.id)
	}

	return timestamp
}

func ReadSchedule(ss pie.Strings) (*Schedule, error) {
	earliest, buses, err := parseSchedule(ss)
	if err != nil {
		return nil, err
	}
	return &Schedule{
		earliestDeparture: earliest,
		buses:             buses,
	}, nil
}

func parseSchedule(ss pie.Strings) (int, *Buses, error) {
	if len := ss.Len(); len > 2 {
		return 0, nil, fmt.Errorf("expected two line entry, found: %d", len)
	}

	timestamp, err := strconv.Atoi(ss[0])
	if err != nil {
		return 0, nil, fmt.Errorf("unable to parse earliest bus timestamp '%s'", ss[0])
	}
	buses, err := parseBuses(ss[1])
	if err != nil {
		return 0, nil, err
	}

	return timestamp, buses, nil
}

func parseBuses(s string) (*Buses, error) {
	buses := []Bus{}
	min := 0
	max := 0
	offset := 0
	for i, b := range strings.Split(s, ",") {
		if b == "x" {
			offset++
			continue
		}
		bus, err := strconv.Atoi(b)
		if err != nil {
			return nil, fmt.Errorf("unable to parse bus id '%s' at entry #%d", b, i)
		}
		buses = append(buses, Bus{id: bus, offset: offset})

		if i == 0 {
			min = bus
			max = bus
		} else {
			if bus > max {
				max = bus
			}
			if bus < min {
				min = bus
			}
		}
		offset++
	}
	return &Buses{list: buses, min: min, max: max}, nil
}

// EarliestTimestampOffsetAlignment takes in a comma separated bus schedule and returns the
// earliest timestamp which the offset aligns
func EarliestTimestampOffsetAlignment(s string) int {
	buses := pie.Ints{}
	// "7,13,x,x,59,x,31,19" --> [7,13,1,1,59,1,31,19]
	pie.Strings(strings.Split(s, ",")).Each(func(s string) {
		i, err := strconv.Atoi(s)
		if err != nil {
			buses = buses.Append(1)
		} else {
			buses = buses.Append(i)
		}
	})

	step := buses[0]
	timestamp := 0

	for i, bus := range buses {
		if bus == 1 || bus == step {
			continue
		}

		for (timestamp+i)%bus != 0 {
			timestamp += step
		}

		step = helpers.LCM(step, bus)
	}

	return timestamp
}
