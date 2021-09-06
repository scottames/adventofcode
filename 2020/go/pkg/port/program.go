package port

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elliotchance/pie/pie"
)

func InitializeProgram(ss pie.Strings) (Memory, error) {
	mem := newMemory()
	mask := newMask("")

	// TODO: functionalize
	for i, s := range ss {
		kv := strings.Split(s, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("unexpected value at index %d: %s", i, s)
		}
		k := kv[0]
		v := strings.TrimSpace(kv[1])
		if strings.Contains(k, "mask") {
			mask = newMask(strings.TrimSpace(kv[1]))
		} else {
			input, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("unable to process value at index %d: %s", i, s)
			}
			err = mem.setAddressString(k, mask.apply(value(input)))
			if err != nil {
				return nil, fmt.Errorf("unable to determine memory address at index %d: %s", i, s)
			}
		}
	}
	return mem, nil
}
