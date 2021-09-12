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
	testMask1 = mask{
		0:  2,
		1:  0,
		2:  2,
		3:  2,
		4:  2,
		5:  2,
		6:  1,
		7:  2,
		8:  2,
		9:  2,
		10: 2,
		11: 2,
		12: 2,
		13: 2,
		14: 2,
		15: 2,
		16: 2,
		17: 2,
		18: 2,
		19: 2,
		20: 2,
		21: 2,
		22: 2,
		23: 2,
		24: 2,
		25: 2,
		26: 2,
		27: 2,
		28: 2,
		29: 2,
		30: 2,
		31: 2,
		32: 2,
		33: 2,
		34: 2,
		35: 2,
	}
)

func TestMemory_newMask(t *testing.T) {
	expected := testMask1
	actual := newMask(testMask1Str)
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestMemory_applyMask11(t *testing.T) {
	expected := value(73)
	actual := testMask1.apply(11)
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestMemory_applyMask101(t *testing.T) {
	expected := value(101)
	actual := testMask1.apply(101)
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestMemory_applyMask0(t *testing.T) {
	expected := value(64)
	actual := testMask1.apply(0)
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
