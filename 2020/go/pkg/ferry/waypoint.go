package ferry

import (
	"fmt"

	"github.com/elliotchance/pie/pie"
)

// TODO: rename to something other than Waypoint?

func NewWaypoints(ss pie.Strings) (*Waypoints, error) {
	// TODO - use https://pkg.go.dev/container/list
	wps := Waypoints{}
	for i, s := range ss {
		wp, err := NewWaypoint(s, wps.tail)
		if err != nil {
			return nil, fmt.Errorf("line %d - %w", i, err)
		}
		wps.Append(wp)
	}

	return &wps, nil
}

type Waypoints struct {
	head *Waypoint
	tail *Waypoint
	len  int
}

func (self *Waypoints) Append(a *Waypoint) {
	if self.head == nil {
		self.head = a
	}
	self.tail = a
	self.len++
}

func (self Waypoints) Last() *Waypoint {
	return self.tail
}

func (self Waypoints) Len() int {
	return self.len
}

func (self Waypoints) ManhattanDistance() (int, error) {
	return self.Last().ManhattanDistance()
}

func NewWaypoint(s string, previous *Waypoint) (*Waypoint, error) {
	action, units, err := parseAction(s)
	if err != nil {
		return nil, err
	} else if action == nil {
		return nil, fmt.Errorf("invalid action found")
	} else if units == nil {
		return nil, fmt.Errorf("invalid units found")
	}

	var ship point
	var wp waypoint
	if previous != nil {
		p, pp, err := previous.Result()
		if err != nil {
			return nil, err
		}
		ship, wp = *p, *pp
	} else {
		ship = point{0, 0}
		wp = waypoint{10, 1}
	}

	return &Waypoint{
		previous: previous,
		action:   *action,
		units:    *units,
		ship:     ship,
		waypoint: wp,
	}, nil
}

type Waypoint struct {
	previous *Waypoint
	action   string
	units    int
	ship     point
	waypoint waypoint
}

func (self *Waypoint) Result() (*point, *waypoint, error) {
	action := self.action

	switch action {
	case Left, Right:
		p, wp := self.ship, self.waypoint.turn(self.action, self.units)
		return &p, &wp, nil
	case Forward:
		p, wp := self.MoveToWaypoint(self.units), self.waypoint
		return &p, &wp, nil
	}

	p := self.ship
	wp, err := self.waypoint.move(self.action, self.units)

	return &p, wp, err
}

func (self *Waypoint) MoveToWaypoint(num int) point {
	p := self.ship
	for i := 0; i < num; i++ {
		p = p.add(point(self.waypoint))
	}
	return p
}

func (self *Waypoint) ManhattanDistance() (int, error) {
	if self == nil {
		return 0, fmt.Errorf("nil Waypoint")
	}

	p, _, err := self.Result()

	return p.manhattanDistance(), err
}

type waypoint point

func (self waypoint) move(action string, num int) (*waypoint, error) {
	switch action {
	case North:
		return &waypoint{self.x, self.y + num}, nil
	case South:
		return &waypoint{self.x, self.y - num}, nil
	case East:
		return &waypoint{self.x + num, self.y}, nil
	case West:
		return &waypoint{self.x - num, self.y}, nil
	}
	return nil, fmt.Errorf("invalid action found when attempting to move waypoint: %s", action)
}

func (self waypoint) turn(action string, num int) waypoint {
	p := self
	x, y := 1, 1
	switch action {
	case Left:
		y = -1
	case Right:
		x = -1
	}
	for i := 0; i < (num / 90); i++ { // assuming num will always be one of: [90, 180, 270]
		p = waypoint{(p.y * y), (p.x * x)}
	}
	return p
}
