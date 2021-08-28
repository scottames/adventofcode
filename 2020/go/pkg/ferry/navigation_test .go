package ferry

import (
	"fmt"
	"testing"

	"github.com/elliotchance/pie/pie"
	"github.com/stretchr/testify/assert"
)

var (
	example = pie.Strings{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	// Part 1
	start1            = Point{1, 0}
	moveWaypoint1     = false
	testWaypointEast  = start1
	testWaypointSouth = Point{0, -1}
	testAction2       = Action{
		waypoint: testWaypointEast,
		action:   "F",
		units:    10,
		ship:     Point{0, 0},
	}
	testAction5 = Action{
		waypoint: testWaypointEast,
		action:   "N",
		units:    3,
		ship:     Point{10, 0},
		previous: &testAction2,
	}
	testAction6 = Action{
		waypoint: testWaypointEast,
		action:   "F",
		units:    7,
		ship:     Point{10, 3},
		previous: &testAction5,
	}
	testAction7 = Action{
		waypoint: testWaypointEast,
		action:   "R",
		units:    90,
		ship:     Point{17, 3},
		previous: &testAction6,
	}
	testAction8 = Action{
		waypoint: testWaypointSouth,
		action:   "F",
		units:    11,
		ship:     Point{17, 3},
		previous: &testAction7,
	}
	testAction9 = Action{
		waypoint: testWaypointEast,
		action:   "F",
		units:    100,
		ship:     Point{17, -8},
		previous: &testAction8,
	}
	testActions1 = Actions{
		head: &testAction2,
		tail: &testAction9,
		len:  5,
	}

	// Part 2
	start2        = Point{10, 1}
	moveWaypoint2 = true
	testAction10  = Action{
		action:       "F",
		units:        10,
		waypoint:     Point{10, 1},
		ship:         Point{0, 0},
		moveWaypoint: moveWaypoint2,
	}
	testAction11 = Action{
		action:       "N",
		units:        3,
		waypoint:     Point{10, 4},
		ship:         Point{100, 10},
		previous:     &testAction10,
		moveWaypoint: moveWaypoint2,
	}
	// R90 rotates the waypoint around the ship clockwise 90 degrees, moving it to 4 units east
	// and 10 units south of the ship. The ship remains at east 170, north 38.
	// Result of ^
	testAction12 = Action{
		action:       "F",
		units:        7,
		waypoint:     Point{4, -10},
		ship:         Point{170, 38},
		previous:     &testAction11,
		moveWaypoint: moveWaypoint2,
	}
	// F11 moves the ship to the waypoint 11 times (a total of 44 units east and 110 units south),
	// leaving the ship at east 214, south 72. The waypoint stays 4 units east and 10 units south
	// of the ship.
	// Result of ^
	testAction13 = Action{
		action:       "R",
		units:        90,
		waypoint:     Point{44, -110},
		ship:         Point{170, 38},
		previous:     &testAction12,
		moveWaypoint: moveWaypoint2,
	}
	testAction14 = Action{
		action:       "F",
		units:        11,
		waypoint:     Point{44, -110},
		ship:         Point{214, -72},
		previous:     &testAction13,
		moveWaypoint: moveWaypoint2,
	}
	testActions2 = Actions{
		head: &testAction10,
		tail: &testAction14,
		len:  5,
	}
)

// Part 1
func Test_NewActionFromStart(t *testing.T) {
	expected := &testAction2
	actual, err := NewAction("F10", nil, &start1, moveWaypoint1)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionFromPrevious1(t *testing.T) {
	expected := &testAction5
	actual, err := NewAction("N3", &testAction2, &start1, moveWaypoint1)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionFromPrevious2(t *testing.T) {
	expected := &testAction6
	actual, err := NewAction("F7", &testAction5, &start1, moveWaypoint1)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionFromPrevious3(t *testing.T) {
	expected := &testAction7
	actual, err := NewAction("R90", &testAction6, &start1, moveWaypoint1)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionFromPrevious4(t *testing.T) {
	expected := &testAction8
	actual, err := NewAction("F11", &testAction7, &start1, moveWaypoint1)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionsManhattanDistance(t *testing.T) {
	expected := 25
	actions, err := NewActions(example, &start1, moveWaypoint1)
	if err != nil {
		assert.Error(t, err)
	}
	actual, err := actions.ManhattanDistance()
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

// Part 2
func Test_WaypointTurn(t *testing.T) {
	testMap := [][]Point{
		{Point{-1, 10}, testAction10.waypoint.turn("L", 90)},
		{Point{-10, -1}, testAction10.waypoint.turn("L", 180)},
		{Point{1, -10}, testAction10.waypoint.turn("L", 270)},

		{Point{1, -10}, testAction10.waypoint.turn("R", 90)},
		{Point{-10, -1}, testAction10.waypoint.turn("R", 180)},
		{Point{-1, 10}, testAction10.waypoint.turn("R", 270)},
	}
	for _, test := range testMap {
		expected := test[0]
		actual := test[1]
		msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
		assert.Equal(t, expected, actual, msg)
	}
}

func Test_WaypointMove(t *testing.T) {
	n, err := testAction10.waypoint.move("N", 3)
	if err != nil {
		assert.Error(t, err)
	}
	s, err := testAction10.waypoint.move("S", 4)
	if err != nil {
		assert.Error(t, err)
	}
	e, err := testAction10.waypoint.move("E", 6)
	if err != nil {
		assert.Error(t, err)
	}
	w, err := testAction10.waypoint.move("W", 7)
	if err != nil {
		assert.Error(t, err)
	}
	testMap := [][]Point{
		{Point{10, 4}, *n},
		{Point{10, -3}, *s},
		{Point{16, 1}, *e},
		{Point{3, 1}, *w},
	}
	for _, test := range testMap {
		expected := test[0]
		actual := test[1]
		msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
		assert.Equal(t, expected, actual, msg)
	}
}

func Test_MoveToWaypoint1(t *testing.T) {
	expected := Point{100, 10}
	actual := testAction10.MoveToWaypoint(10)
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_MoveToWaypoint2(t *testing.T) {
	expected := Point{170, 38}
	actual := testAction11.MoveToWaypoint(7)
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_MoveToWaypoint3(t *testing.T) {
	expected := Point{44, -110}
	actual := testAction12.MoveToWaypoint(10)
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_Waypoint4Result(t *testing.T) {
	actualShip, actualWaypoint, err := testAction13.Result(moveWaypoint2)
	if err != nil {
		assert.Error(t, err)
	}

	expectedShip := Point{170, 38}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expectedShip, actualShip)
	assert.Equal(t, expectedShip, actualShip, msg)

	expectedWaypoint := Point{44, -110}
	msg = fmt.Sprintf("Expected %#v. Got %#v.", expectedWaypoint, actualWaypoint)
	assert.Equal(t, expectedWaypoint, actualWaypoint, msg)
}

func Test_Waypoint4ManhattanDistance(t *testing.T) {
	expected := 286
	ship, _, err := testAction13.Result(moveWaypoint2)
	if err != nil {
		assert.Error(t, err)
	}
	actual := ship.manhattanDistance()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewWayPointsManhattanDistance(t *testing.T) {
	expected := 286
	wps, err := NewActions(example, &start2, moveWaypoint2)
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

func Test_NewAction1(t *testing.T) {
	expected := testAction10
	actual, err := NewAction("F10", nil, &start2, moveWaypoint2)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewAction2(t *testing.T) {
	expected := testAction11
	actual, err := NewAction("N3", &testAction10, &start2, moveWaypoint2)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewAction3(t *testing.T) {
	expected := testAction12
	actual, err := NewAction("F7", &testAction11, &start2, moveWaypoint2)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewAction4(t *testing.T) {
	expected := testAction13
	actual, err := NewAction("R90", &testAction12, &start2, moveWaypoint2)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewAction5(t *testing.T) {
	expected := testAction14
	actual, err := NewAction("F11", &testAction13, &start2, moveWaypoint2)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActions(t *testing.T) {
	expected := testActions2
	actual, err := NewActions(example, &start2, moveWaypoint2)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
