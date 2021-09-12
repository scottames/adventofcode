package port

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elliotchance/pie/pie"
)

// NewProgram returns a new Program initialized with the given slice of strings
func NewProgram(ss pie.Strings) *Program {
	return &Program{raw: ss}
}

// Program represents the port program
type Program struct {
	memory  memory
	raw     []string
	version func(mem memory, m mask, addr int, val int)
}

// Run the given Program at the given version (latest used if not set with .VX() method)
func (self *Program) Run() (*Program, error) {
	self.initializememory()
	mask := newMask("")

	if self.version == nil {
		self.V2()
	}

	for i, s := range self.Raw() {
		k, v, err := splitKV(s)
		if err != nil {
			return nil, fmt.Errorf("%w at index %d: %s", err, i, s)
		}
		if strings.Contains(k, "mask") {
			mask = newMask(strings.TrimSpace(v))
			continue
		}
		addr, err := extractMemAddr(k)
		if err != nil {
			return nil, fmt.Errorf("%w at index %d: %s", err, i, s)
		}

		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("%w at index %d: %s", err, i, s)
		}
		self.version(self.memory, mask, addr, val)
	}
	return self, nil
}

// Raw returns the raw input to the Program
func (self *Program) Raw() []string {
	return self.raw
}

// Sum returns the sum of the Program memory
func (self *Program) Sum() int {
	return self.memory.sum()
}

// V1 sets the Program to run with version V1
func (self *Program) V1() *Program {
	self.version = func(mem memory, m mask, addr int, val int) {
		mem.setValue(value(addr), m.apply(value(val)))
	}
	return self
}

// V2 sets the Program to run with version V2
func (self *Program) V2() *Program {
	self.version = func(mem memory, m mask, addr int, val int) {
		addrs := m.applyFloating(value(addr))
		for _, a := range addrs {
			mem.setValue(a, value(val))
		}
	}
	return self
}

func (self *Program) initializememory() {
	if self.memory == nil {
		self.memory = make(memory)
	}
}

func splitKV(s string) (key string, val string, err error) {
	kv := strings.Split(s, "=")
	if l := len(kv); l != 2 {
		return "", "", fmt.Errorf("expected two values, got '%d'", l)
	}
	key = kv[0]
	val = strings.TrimSpace(kv[1])
	return
}
