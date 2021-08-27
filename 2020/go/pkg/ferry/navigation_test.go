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
	testAction1 = Action{heading: north}
	testAction2 = Action{
		heading:     east,
		action:      "F",
		units:       10,
		coordinates: point{0, 0},
	}
	testAction3 = Action{heading: south}
	testAction4 = Action{heading: west}
	testAction5 = Action{
		heading:     east,
		action:      "N",
		units:       3,
		coordinates: point{10, 0},
		previous:    &testAction2,
	}
	testAction6 = Action{
		heading:     east,
		action:      "F",
		units:       7,
		coordinates: point{10, 3},
		previous:    &testAction5,
	}
	testAction7 = Action{
		heading:     east,
		action:      "R",
		units:       90,
		coordinates: point{17, 3},
		previous:    &testAction6,
	}
	testAction8 = Action{
		heading:     south,
		action:      "F",
		units:       11,
		coordinates: point{17, 3},
		previous:    &testAction7,
	}
	testAction9 = Action{
		heading:     south,
		action:      "F",
		units:       100,
		coordinates: point{17, -8},
		previous:    &testAction8,
	}
	testActions = Actions{
		head: &testAction2,
		tail: &testAction9,
		len:  5,
	}
)

func Test_ActionTurn(t *testing.T) {
	testMap := [][]heading{
		// North - Turn L 90
		{west, testAction1.turn("L", 90)},
		// North - Turn L 180
		{south, testAction1.turn("L", 180)},
		// North - Turn L 270
		{east, testAction1.turn("L", 270)},
		// North - Turn L 360
		{north, testAction1.turn("L", 360)},
		// North - Turn R 90
		{east, testAction1.turn("R", 90)},
		// North - Turn R 180
		{south, testAction1.turn("R", 180)},
		// North - Turn R 270
		{west, testAction1.turn("R", 270)},
		// North - Turn R 360
		{north, testAction1.turn("R", 360)},
		// East - Turn L 90
		{north, testAction2.turn("L", 90)},
		// East - Turn L 180
		{west, testAction2.turn("L", 180)},
		// East - Turn L 270
		{south, testAction2.turn("L", 270)},
		// East - Turn L 360
		{east, testAction2.turn("L", 360)},
		// East - Turn R 90
		{south, testAction2.turn("R", 90)},
		// East - Turn R 180
		{west, testAction2.turn("R", 180)},
		// East - Turn R 270
		{north, testAction2.turn("R", 270)},
		// East - Turn R 360
		{east, testAction2.turn("R", 360)},
	}
	for _, test := range testMap {
		expected := test[0]
		actual := test[1]
		msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
		assert.Equal(t, expected, actual, msg)
	}
}

func Test_NewActionFromStart(t *testing.T) {
	expected := &testAction2
	actual, err := NewAction("F10", nil)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionFromPrevious1(t *testing.T) {
	expected := &testAction5
	actual, err := NewAction("N3", &testAction2)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionFromPrevious2(t *testing.T) {
	expected := &testAction6
	actual, err := NewAction("F7", &testAction5)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionFromPrevious3(t *testing.T) {
	expected := &testAction7
	actual, err := NewAction("R90", &testAction6)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionFromPrevious4(t *testing.T) {
	expected := &testAction8
	actual, err := NewAction("F11", &testAction7)
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_NewActionsManhattanDistance(t *testing.T) {
	expected := 25
	actions, err := NewActions(example)
	if err != nil {
		assert.Error(t, err)
	}
	actual := actions.ManhattanDistance()
	msg := fmt.Sprintf("Expected %#v. Got %#v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
