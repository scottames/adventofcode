package trajectory

import (
	"github.com/elliotchance/pie/pie"
)

const tree = "#"

// NewSlope returns a pointer to a new Splope
// from the given hill and steps (Coordinates)
func NewSlope(hill *pie.Strings, steps *Coordinates) *Slope {
	if steps == nil {
		steps = NewCoordinates(0, 0)
	}
	return &Slope{
		trajectory: steps,
		position: Coordinates{
			x: 0,
			y: 0,
		},
		rowCount:  hill.Len(),
		cols:      len(hill.First()),
		rows:      hill,
		treeCount: 0,
	}
}

// Slope represents a sled's path along a given map
type Slope struct {
	trajectory *Coordinates
	position   Coordinates
	rowCount   int
	cols       int
	rows       *pie.Strings
	treeCount  int
}

// IsJourneyComplete returns true if the sled's
// journey is complete along it's trajectory
func (s *Slope) IsJourneyComplete() bool {
	return s.position.y >= s.rowCount
}

// Run steps through the slope's trajectory
// until it is complete
// returning a pointer to the slope
func (s *Slope) Run() *Slope {
	if !s.IsJourneyComplete() {
		for s.nextStep(); !s.IsJourneyComplete(); s.nextStep() {
		}
	}
	return s
}

// TreesEncountered returns the number (int) of trees
// encountered by the Slope in any state
func (s *Slope) TreesEncountered() int {
	return s.treeCount
}

// currentChar returns the current character at the given
// position of the Slope's trajectory
func (s *Slope) currentChar() string {
	if s.IsJourneyComplete() {
		return ""
	}
	return string([]rune((*s.rows)[s.position.y])[s.position.x])
}

// nextCoordinates returns the next Coordinates
// on the Slope's trajectory
func (s *Slope) nextCoordinates() Coordinates {
	return Coordinates{
		s.position.x + s.trajectory.x,
		s.position.y + s.trajectory.y,
	}
}

// nextStep takes the next step in the trajectory
// of the given Slope - modifying the Slope itself
func (s *Slope) nextStep() {
	next := s.nextCoordinates()
	if next.y > s.rowCount {
		s.position.y = s.rowCount
		s.position.x = s.cols
		return
	}
	if next.x >= s.cols {
		s.position.x = next.x - s.cols
	} else {
		s.position.x = next.x
	}
	s.position.y = next.y
	if s.currentChar() == tree {
		s.treeCount++
	}
}
