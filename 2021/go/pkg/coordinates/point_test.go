package coordinates_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/scottames/adventofcode/2021/go/pkg/coordinates"
)

func Test_directions_forward(t *testing.T) {
	expected := coordinates.Point{X: 1, Y: 0}
	actual := coordinates.New(0, 0).Add(coordinates.Point{X: 1, Y: 0})

	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_directions_back(t *testing.T) {
	expected := coordinates.Point{X: -1, Y: 0}
	actual := coordinates.New(0, 0).Add(coordinates.Point{X: -1, Y: 0})

	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_directions_down(t *testing.T) {
	expected := coordinates.Point{X: 0, Y: -1}
	actual := coordinates.New(0, 0).Add(coordinates.Point{X: 0, Y: -1})

	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_directions_up(t *testing.T) {
	expected := coordinates.Point{X: 0, Y: 1}
	actual := coordinates.New(0, 0).Add(coordinates.Point{X: 0, Y: 1})

	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
