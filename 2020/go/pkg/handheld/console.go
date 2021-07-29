package handheld

import (
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	argument  int
	position  int
}

func NewInstruction(s string, position int) Instruction {
	ss := strings.Split(s, " ")
	if len(ss) != 2 {
		return Instruction{}
	}
	arg, err := strconv.Atoi(strings.TrimSpace(ss[1]))
	if err != nil {
		return Instruction{}
	}
	return Instruction{
		operation: strings.TrimSpace(ss[0]),
		argument:  arg,
		position:  position,
	}
}

func (self Instruction) Argument() int {
	return self.argument
}

func (self Instruction) Operation() string {
	return self.operation
}

func (self Instruction) Position() int {
	return self.position
}

func (self Instruction) Next() int {
	switch self.operation {
	case "jmp":
		return self.position + self.argument
	case "nop":
		return self.position + 1
	case "acc":
		return self.position + 1

	}
	return self.position + 1
}

type Instructions []Instruction

func NewInstructions(s string) Instructions {
	instructions := Instructions{}
	for i, is := range strings.Split(s, "\n") {
		instructions = instructions.Append(NewInstruction(is, i))
	}
	return instructions
}

func (self Instructions) executeUntilBroken() (result int, finalNode int) {
	positions := make(map[int]int) // index: count
	instruction := self[0]
	accumulator := 0
	for i := 0; i < len(self); i++ {
		position := instruction.Position()
		next := instruction.Next()
		positions[position]++
		if instruction.Operation() == "acc" {
			accumulator += instruction.Argument()
		}
		if next < 0 || next >= len(self) || positions[next] >= 1 {
			return accumulator, position
		}
		instruction = self[next]
	}
	return accumulator, 0
}

func (self Instructions) Append(i ...Instruction) Instructions {
	result := append([]Instruction{}, self...)
	result = append(result, i...)
	return result
}

func (self Instructions) Execute() int {
	result, _ := self.executeUntilBroken()
	return result
}

func (self Instructions) Fix() int {
	for index, instruction := range self {
		newSelf := self.Copy()
		switch instruction.Operation() {
		case "jmp":
			newSelf[index].operation = "nop"
		case "nop":
			newSelf[index].operation = "jmp"
		}
		result, final := newSelf.executeUntilBroken()
		if final+1 == len(self) {
			return result
		}
	}
	return 0
}

func (self Instructions) Copy() Instructions {
	new := make(Instructions, len(self))
	copy(new, self)
	return new
}
