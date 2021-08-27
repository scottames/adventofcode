package ferry

import (
	"fmt"
	"strconv"

	"github.com/elliotchance/pie/pie"
	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	North   string = "N"
	East    string = "E"
	South   string = "S"
	West    string = "W"
	Left    string = "L"
	Right   string = "R"
	Forward string = "F"
)

const (
	north = heading(0)
	east  = heading(90)
	south = heading(180)
	west  = heading(270)
)

// heading is an integer representation of the heading on a grid
// assumed to be divisible by 90 and within 360 degrees
type heading int

// str returns the string representation of the heading
func (self heading) str() string {
	hm := map[heading]string{
		north: North,
		east:  East,
		south: South,
		west:  West,
	}
	return hm[self]
}

// NewActions returns a new set of Actions from a given slice of strings
func NewActions(ss pie.Strings) (*Actions, error) {
	// TODO - use https://pkg.go.dev/container/list
	actions := Actions{}
	for i, s := range ss {
		action, err := NewAction(s, actions.tail)
		if err != nil {
			return nil, fmt.Errorf("line %d - %w", i, err)
		}
		actions.Append(action)
	}

	return &actions, nil
}

// Actions represents a linked list of Actions
type Actions struct {
	head *Action
	tail *Action
	len  int
}

// Append appends a new Action to the end of the Actions
func (self *Actions) Append(a *Action) {
	if self.head == nil {
		self.head = a
	}
	self.tail = a
	self.len++
}

// Last returns the last Action from the given Actions
func (self Actions) Last() *Action {
	return self.tail
}

// Len returns the length of the given Actions
func (self Actions) Len() int {
	return self.len
}

// ManhattanDistance returns the ManhattanDistance from the last Action in the Actions list
func (self Actions) ManhattanDistance() int {
	return self.Last().ManhattanDistance()
}

// NewAction returns a new Action from the given string
// linked to the previous Action, if provided
func NewAction(s string, previous *Action) (*Action, error) {
	action, units, err := parseAction(s)
	if err != nil {
		return nil, err
	} else if action == nil {
		return nil, fmt.Errorf("invalid action found")
	} else if units == nil {
		return nil, fmt.Errorf("invalid units found")
	}

	var coordinates point
	var heading heading
	if previous != nil {
		coordinates, heading = previous.Result()
	} else {
		coordinates = point{0, 0}
		heading = east
	}

	return &Action{
		previous:    previous,
		action:      *action,
		heading:     heading,
		units:       *units,
		coordinates: coordinates,
	}, nil
}

// Action represents a single action in the given string of Actions that represents the Navigation
type Action struct {
	previous *Action

	// TODO: make action a type
	action      string
	heading     heading
	units       int
	coordinates point
}

// Result returns a new set of coordinates and the heading as a result of the given Action
func (self *Action) Result() (point, heading) {
	action := self.action

	switch action {
	case Left, Right:
		return self.coordinates, self.turn(self.action, self.units)
	case Forward:
		action = self.heading.str()
	}

	// move in the given direction
	return self.coordinates.addInDirection(action, self.units), self.heading
}

// turn returns a new heading given the direction
func (self *Action) turn(direction string, degrees int) heading {
	d := heading((degrees % 360))
	if direction == Left {
		return (self.heading + 360 - d) % 360
	}
	return (self.heading + d) % 360
}

// X returns the X coordinates for the given Action
func (self *Action) X() int {
	if self == nil {
		return 0
	}

	return self.coordinates.x
}

// Y returns the Y coordinates for the given Action
func (self *Action) Y() int {
	if self == nil {
		return 0
	}

	return self.coordinates.y
}

// ManhattanDistance returns the sum of the absolute values from the starting point
// All Actions are assumed to have started at point{0,0} thus the ManhattanDistance is the
// sum of the x & y coordinates
func (self *Action) ManhattanDistance() int {
	if self == nil {
		return 0
	}

	resultingPoint, _ := self.Result()

	return resultingPoint.manhattanDistance()
}

// point represents a point on a grid - by the x & y axis
type point struct {
	x int
	y int
}

// add returns the result of adding a new point to the given point
func (self point) add(p point) point {
	return point{self.x + p.x, self.y + p.y}
}

func (self point) addInDirection(dir string, i int) point {
	actionMap := map[string]point{
		North: {0, i},
		South: {0, -i},
		East:  {i, 0},
		West:  {-i, 0},
	}
	return self.add(actionMap[dir])
}

func (self point) manhattanDistance() int {
	return helpers.Absolute(self.x) + helpers.Absolute(self.y)
}

// parseAction returns a set of action & units and a possible error
// from the given string
// Example:
//   F11 --> (F, 11, nil)
func parseAction(s string) (*string, *int, error) {
	if s == "" {
		return nil, nil, fmt.Errorf("empty action string found")
	} else if sLen := len(s); sLen < 2 {
		return nil, nil, fmt.Errorf("invalid action string found: %s (with len(%d))", s, sLen)
	}

	action := s[:1]
	unitsStr := s[1:]
	units, err := strconv.Atoi(unitsStr)
	if err != nil {
		return nil,
			nil,
			fmt.Errorf("invalid action - unable to convert '%s' to unit (int): %w", unitsStr, err)
	}

	return &action, &units, nil
}
