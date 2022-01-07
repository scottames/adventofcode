package aoc2021

import (
	"fmt"
	"testing"

	"github.com/scottames/adventofcode/2021/go/pkg/coordinates"
	"github.com/scottames/adventofcode/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

const (
	day02Example = `forward 5
down 5
forward 8
up 3
down 8
forward 2
`
)

func Test_navigate_directionWithoutAim(t *testing.T) {
	expected := coordinates.New(15, -10)

	moves, err := newMoves(helpers.StringSplitNewlinesStrings([]byte(day02Example)))
	if err != nil {
		t.Error(err)
	}

	actual, err := moves.navigate(directionWithoutAim)
	if err != nil {
		t.Errorf(err.Error())
	}

	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func Test_navigate_directionWithAim(t *testing.T) {
	expected := coordinates.New(15, -60)

	moves, err := newMoves(helpers.StringSplitNewlinesStrings([]byte(day02Example)))
	if err != nil {
		t.Error(err)
	}

	actual, err := moves.navigate(directionWithAim)
	if err != nil {
		t.Error(err)
	}

	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
