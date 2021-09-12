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
	testExample2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
)

var (
	testExampleMemory = memory{
		8: 64,
		7: 101,
	}
	testExampleProgram = Program{
		raw:    strings.Split(testExample, "\n"),
		memory: testExampleMemory,
	}
)

func TestProgram_ExampleProgramV1memory(t *testing.T) {
	expected := &testExampleProgram.memory
	prog, err := NewProgram(strings.Split(testExample, "\n")).V1().Run()
	if err != nil {
		assert.Error(t, err)
	}
	actual := &prog.memory
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestProgram_ExampleSum(t *testing.T) {
	expected := 165
	prog := NewProgram(strings.Split(testExample, "\n"))
	v1, err := prog.V1().Run()
	if err != nil {
		assert.Error(t, err)
	}
	actual := v1.Sum()
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}

func TestProgram_Example2Sum(t *testing.T) {
	expected := 208
	prog := NewProgram(strings.Split(testExample2, "\n"))
	v1, err := prog.V2().Run()
	if err != nil {
		assert.Error(t, err)
	}
	actual := v1.Sum()
	msg := fmt.Sprintf("Expected %v. Got %v.", expected, actual)
	assert.Equal(t, expected, actual, msg)
}
