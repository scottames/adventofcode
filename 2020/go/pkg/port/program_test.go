package port

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testExample = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
)

var (
	testExampleMemory = Memory{
		8: 64,
		7: 101,
	}
)

func TestPassport_Example(t *testing.T) {
	expected := testExampleMemory
	actual, err := InitializeProgram(strings.Split(testExample, "\n"))
	if err != nil {
		assert.Error(t, err)
	}
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestPassport_ExampleSum(t *testing.T) {
	expected := 165
	prog, err := InitializeProgram(strings.Split(testExample, "\n"))
	if err != nil {
		assert.Error(t, err)
	}
	actual := prog.Sum()
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
