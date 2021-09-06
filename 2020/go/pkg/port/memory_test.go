package port

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testMask1Str = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
)

var (
	testMask1 = mask{1: 0, 6: 1}
)

func TestPassport_newMask(t *testing.T) {
	expected := testMask1
	actual := newMask(testMask1Str)
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_applyMask11(t *testing.T) {
	expected := value(73)
	actual := testMask1.apply(11)
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_applyMask101(t *testing.T) {
	expected := value(101)
	actual := testMask1.apply(101)
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_applyMask0(t *testing.T) {
	expected := value(64)
	actual := testMask1.apply(0)
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
