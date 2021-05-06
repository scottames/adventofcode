package aoc2019

import (
	"fmt"
	"image"
	"strconv"

	"github.com/scottames/adventofcode/pkg/helpers"
)

const (
	right = "R"
	left  = "L"
	up    = "U"
	down  = "D"
)

// Day03 - 2019 Day 3 Part 1 & 2
func Day03() error {
	input, err := helpers.ReadInput(2019, 3)
	if err != nil {
		return err
	}

	fmt.Println("--- Part 1 ---")
	fmt.Printf("\nThe Manhattan distance from the central port to the closest intersection:\n\t %d\n",
		breakInputToWires(input).findPoint(closestIntersectToStart),
	)

	fmt.Println("\n--- Part 2 ---")
	fmt.Printf("\nThe fewest combined steps the wires must take to reach an intersection:\n\t %d\n",
		breakInputToWires(input).findPoint(shortestDistanceToIntersect),
	)
	return nil
}

type wires struct {
	wires      []*wire
	intersects map[point]int // where int is the distance traveled
}

func newWires(strings []string) *wires {
	ws := wires{}
	for i, s := range strings {
		w := newWire(s, i, &ws)
		ws.wires = append(ws.wires, w)
		ws.intersects = make(map[point]int)
	}
	return &ws
}

func (ws *wires) findPoint(fn func(*wires) int) int {
	ws.build()
	return fn(ws)
}

func (ws *wires) build() *wires {
	for _, w := range ws.wires {
		w.toInstructions().toPoints()
	}
	return ws
}

type wire struct {
	num          int
	input        string
	instructions []instruction
	points       map[point]int // where int is the distance traveled
	wires        *wires
}

type instruction struct {
	direction string
	distance  int
}

type point image.Point

func newWire(s string, num int, wires *wires) *wire {
	return &wire{
		num:    num,
		input:  s,
		points: make(map[point]int, 0),
		wires:  wires,
	}
}

func (w *wire) toInstructions() *wire {
	if len(w.input) == 0 {
		helpers.ExitOnError(fmt.Errorf("cannot calculate instructions: wire input is of zero length"))
	}
	split := helpers.StringSplitCommaToStrings([]byte(w.input))
	for _, strInst := range split {
		instr, err := convertStringToInstruction(strInst)
		helpers.ExitOnError(err)
		w.instructions = append(w.instructions, instr)
	}
	return w
}

func (w *wire) toPoints() *wire {
	pointer := point{X: 0, Y: 0}
	steps := 0
	for _, instr := range w.instructions {
		switch instr.direction {
		case right:
			w.instrIter(&steps, &pointer, &pointer.X, instr.distance, helpers.AddInt)
		case left:
			w.instrIter(&steps, &pointer, &pointer.X, instr.distance, helpers.SubtractInt)
		case up:
			w.instrIter(&steps, &pointer, &pointer.Y, instr.distance, helpers.AddInt)
		case down:
			w.instrIter(&steps, &pointer, &pointer.Y, instr.distance, helpers.SubtractInt)
		}
	}
	return w
}

func (w *wire) instrIter(steps *int, point *point, axis *int, distance int, fn func(a, b int) int) {
	for i := 0; i < distance; i++ {
		*steps++
		*axis = fn(*axis, 1)
		w.points[*point] = *steps
		if w.num == 1 {
			if wireOneSteps, ok := w.wires.wires[0].points[*point]; ok {
				w.wires.intersects[*point] += *steps + wireOneSteps
			}
		}
	}
}

func convertStringToInstruction(i string) (instruction, error) {
	bytes := []byte(i)
	distance, err := strconv.Atoi(string(bytes[1:]))
	if err != nil {
		return instruction{}, err
	}
	return instruction{
		direction: string(bytes[0]),
		distance:  distance,
	}, nil
}

func breakInputToWires(b []byte) *wires {
	wireStrings := helpers.CovertBytesToStrings(b, "\n")
	if lenWires := len(wireStrings); lenWires != 2 {
		helpers.ExitOnError(fmt.Errorf("%d wires found. Expected 2", lenWires))
	}

	wires := newWires(wireStrings)
	if wires == nil {
		helpers.ExitOnError(fmt.Errorf("nil wires found"))
	}
	return wires
}

func closestIntersectToStart(ws *wires) int {
	if len(ws.intersects) == 0 {
		return 0
	}
	var closest int
	for p := range ws.intersects {
		if closest == 0 {
			closest = helpers.Absolute(p.X) + helpers.Absolute(p.Y)
		} else {
			x := helpers.Absolute(p.X)
			y := helpers.Absolute(p.Y)
			if sum := x + y; sum < closest {
				closest = sum
			}
		}
	}

	return closest
}

func shortestDistanceToIntersect(ws *wires) int {
	if len(ws.intersects) == 0 {
		return 0
	}
	var shortest int
	for _, dist := range ws.intersects {
		if shortest == 0 {
			shortest = dist
		} else {
			if dist < shortest {
				shortest = dist
			}
		}
	}
	return shortest
}
