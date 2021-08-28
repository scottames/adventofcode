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

// NewActions returns a new set of Actions from a given slice of strings
func NewActions(ss pie.Strings, start *Point, moveWaypoint bool) (*Actions, error) {
	wps := Actions{}
	for i, s := range ss {
		wp, err := NewAction(s, wps.tail, start, moveWaypoint)
		if err != nil {
			return nil, fmt.Errorf("line %d - %w", i, err)
		}
		wps.Append(wp)
	}

	return &wps, nil
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
func (self Actions) ManhattanDistance() (int, error) {
	return self.Last().ManhattanDistance()
}

// NewAction returns a new Action from the given string
// linked to the previous Action, if provided
func NewAction(s string, previous *Action, start *Point, moveWaypoint bool) (*Action, error) {
	action, units, err := parseAction(s)
	if err != nil {
		return nil, err
	} else if action == nil {
		return nil, fmt.Errorf("invalid action found")
	} else if units == nil {
		return nil, fmt.Errorf("invalid units found")
	}
	if start == nil {
		start = &Point{1, 0}
	}

	var ship Point
	var wp Point
	if previous != nil {
		p, pp, err := previous.Result(moveWaypoint)
		if err != nil {
			return nil, err
		}
		ship, wp = *p, *pp
	} else {
		ship = Point{0, 0}
		wp = *start
	}

	return &Action{
		previous:     previous,
		action:       *action,
		units:        *units,
		ship:         ship,
		waypoint:     wp,
		moveWaypoint: moveWaypoint,
	}, nil
}

type Action struct {
	previous     *Action
	action       string
	units        int
	ship         Point
	waypoint     Point
	moveWaypoint bool
}

func (self *Action) Result(moveWaypoint bool) (ship *Point, waypoint *Point, err error) {
	action := self.action

	switch action {
	case Left, Right:
		p, wp := self.ship, self.waypoint.turn(self.action, self.units)
		return &p, &wp, nil
	case Forward:
		p, wp := self.MoveToWaypoint(self.units), self.waypoint
		return &p, &wp, nil
	}

	if moveWaypoint {
		waypoint, err = self.waypoint.move(self.action, self.units)
		ship = &self.ship
	} else {
		ship, err = self.ship.move(self.action, self.units)
		waypoint = &self.waypoint
	}

	return
}

func (self *Action) MoveToWaypoint(num int) Point {
	p := self.ship
	for i := 0; i < num; i++ {
		p = p.add(Point(self.waypoint))
	}
	return p
}

func (self *Action) ManhattanDistance() (int, error) {
	if self == nil {
		return 0, fmt.Errorf("nil Waypoint")
	}

	p, _, err := self.Result(self.moveWaypoint)

	return p.manhattanDistance(), err
}

func (self Point) move(action string, num int) (*Point, error) {
	switch action {
	case North:
		return &Point{self.X, self.Y + num}, nil
	case South:
		return &Point{self.X, self.Y - num}, nil
	case East:
		return &Point{self.X + num, self.Y}, nil
	case West:
		return &Point{self.X - num, self.Y}, nil
	}
	return nil, fmt.Errorf("invalid action found when attempting to move waypoint: %s", action)
}

func (self Point) turn(action string, num int) Point {
	p := self
	x, y := 1, 1
	switch action {
	case Left:
		y = -1
	case Right:
		x = -1
	}
	for i := 0; i < (num / 90); i++ { // assuming num will always be one of: [90, 180, 270]
		p = Point{(p.Y * y), (p.X * x)}
	}
	return p
}

// point represents a point on a grid - by the x & y axis
type Point struct {
	X int
	Y int
}

// add returns the result of adding a new point to the given point
func (self Point) add(p Point) Point {
	return Point{self.X + p.X, self.Y + p.Y}
}

func (self Point) addInDirection(dir string, i int) Point {
	actionMap := map[string]Point{
		North: {0, i},
		South: {0, -i},
		East:  {i, 0},
		West:  {-i, 0},
	}
	return self.add(actionMap[dir])
}

func (self Point) manhattanDistance() int {
	return helpers.Absolute(self.X) + helpers.Absolute(self.Y)
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
