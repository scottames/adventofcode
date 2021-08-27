package ferry

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/stretchr/testify/assert"
)

var (
	exampleStr = pie.Strings{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	testWaypoint1 = Waypoint{
		action:   "F",
		units:    10,
		waypoint: waypoint{10, 1},
		ship:     point{0, 0},
	}
	testWaypoint2 = Waypoint{
		action:   "N",
		units:    3,
		waypoint: waypoint{10, 4},
		ship:     point{100, 10},
		previous: &testWaypoint1,
	}
	// R90 rotates the waypoint around the ship clockwise 90 degrees, moving it to 4 units east
	// and 10 units south of the ship. The ship remains at east 170, north 38.
	// Result of ^
	testWaypoint3 = Waypoint{
		action:   "F",
		units:    7,
		waypoint: waypoint{4, -10},
		ship:     point{170, 38},
		previous: &testWaypoint2,
	}
	// F11 moves the ship to the waypoint 11 times (a total of 44 units east and 110 units south),
	// leaving the ship at east 214, south 72. The waypoint stays 4 units east and 10 units south
	// of the ship.
	// Result of ^
	testWaypoint4 = Waypoint{
		action:   "R",
		units:    90,
		waypoint: waypoint{44, -110},
		ship:     point{170, 38},
		previous: &testWaypoint3,
	}
	testWaypoint5 = Waypoint{
		action:   "F",
		units:    11,
		waypoint: waypoint{44, -110},
		ship:     point{214, -72},
		previous: &testWaypoint4,
	}
	testWaypoints = Waypoints{
		head: &testWaypoint1,
		tail: &testWaypoint5,
		len:  5,
	}
)

func Test_WaypointTurn(t *testing.T) {
	testMap := [][]waypoint{
		{waypoint{-1, 10}, testWaypoint1.waypoint.turn("L", 90)},
		{waypoint{-10, -1}, testWaypoint1.waypoint.turn("L", 180)},
		{waypoint{1, -10}, testWaypoint1.waypoint.turn("L", 270)},

		{waypoint{1, -10}, testWaypoint1.waypoint.turn("R", 90)},
		{waypoint{-10, -1}, testWaypoint1.waypoint.turn("R", 180)},
		{waypoint{-1, 10}, testWaypoint1.waypoint.turn("R", 270)},
	}
	for _, test := range testMap {
		expected := test[0]
		actual := test[1]
		msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
		assert.Equal(t, expected, actual, msg)
	}
}

func Test_WaypointMove(t *testing.T) {
	n, err := testWaypoint1.waypoint.move("N", 3)
	if err != nil {
		assert.Error(t, err)
	}
	s, err := testWaypoint1.waypoint.move("S", 4)
	if err != nil {
		assert.Error(t, err)
	}
	e, err := testWaypoint1.waypoint.move("E", 6)
	if err != nil {
		assert.Error(t, err)
	}
	w, err := testWaypoint1.waypoint.move("W", 7)
	if err != nil {
		assert.Error(t, err)
	}
	testMap := [][]waypoint{
		{waypoint{10, 4}, *n},
		{waypoint{10, -3}, *s},
		{waypoint{16, 1}, *e},
		{waypoint{3, 1}, *w},
	}
	for _, test := range testMap {
		expected := test[0]
		actual := test[1]
		msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
		assert.Equal(t, expected, actual, msg)
	}
}

func Test_MoveToWaypoint1(t *testing.T) {
	expected := point{100, 10}
	actual := testWaypoint1.MoveToWaypoint(10)
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_MoveToWaypoint2(t *testing.T) {
	expected := point{170, 38}
	actual := testWaypoint2.MoveToWaypoint(7)
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_MoveToWaypoint3(t *testing.T) {
	expected := point{44, -110}
	actual := testWaypoint3.MoveToWaypoint(10)
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Waypoint4Result(t *testing.T) {
	actualShip, actualWaypoint, err := testWaypoint4.Result()
	if err != nil {
		assert.Error(t, err)
	}

	expectedShip := point{170, 38}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expectedShip, actualShip)
	assert.Equal(t, expectedShip, actualShip, msg)

	expectedWaypoint := waypoint{44, -110}
	msg = fmt.Sprintf("Expected %#v. Got %#v.", expectedWaypoint, actualWaypoint)
	assert.Equal(t, expectedWaypoint, actualWaypoint, msg)
}

func Test_Waypoint4ManhattanDistance(t *testing.T) {
	expected := 286
	ship, _, err := testWaypoint4.Result()
	if err != nil {
		assert.Error(t, err)
	}
	actual := ship.manhattanDistance()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewWayPointsManhattanDistance(t *testing.T) {
	expected := 286
	wps, err := NewWaypoints(exampleStr)
	if err != nil {
		assert.Error(t, err)
	}
	actual, err := wps.Last().ManhattanDistance()
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewWaypoint1(t *testing.T) {
	expected := testWaypoint1
	actual, err := NewWaypoint("F10", nil)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewWaypoint2(t *testing.T) {
	expected := testWaypoint2
	actual, err := NewWaypoint("N3", &testWaypoint1)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewWaypoint3(t *testing.T) {
	expected := testWaypoint3
	actual, err := NewWaypoint("F7", &testWaypoint2)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewWaypoint4(t *testing.T) {
	expected := testWaypoint4
	actual, err := NewWaypoint("R90", &testWaypoint3)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewWaypoint5(t *testing.T) {
	expected := testWaypoint5
	actual, err := NewWaypoint("F11", &testWaypoint4)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewWaypoints(t *testing.T) {
	expected := testWaypoints
	actual, err := NewWaypoints(exampleStr)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
