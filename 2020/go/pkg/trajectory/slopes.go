package trajectory

import "github.com/elliotchance/pie/pie"

// NewSlopes returns a pointer to Slopes from the given
// hill (pie.Strings) an one or more Coordinates
func NewSlopes(hill *pie.Strings, cs ...*Coordinates) *Slopes {
	var ss = &Slopes{}
	if cs == nil {
		return ss
	}
	for _, c := range cs {
		ss.Append(NewSlope(hill, c))
	}
	return ss
}

// Slopes represents a slice of Slope
type Slopes []*Slope

// Append one or more Slope to the given Slopes
func (ss *Slopes) Append(s ...*Slope) *Slopes {
	if s == nil || ss == nil {
		return ss
	}
	*ss = append(*ss, s...)
	return ss
}

// Run all Slopes' trajectories
func (ss *Slopes) Run() *Slopes {
	if ss == nil {
		return ss
	}
	for _, s := range *ss {
		s.Run()
	}
	return ss
}

// TreesEncountered returns a slice of integers (pie.Ints)
// of the number of trees encountered for the given Slopes
func (ss *Slopes) TreesEncountered() *pie.Ints {
	if ss == nil {
		return nil
	}
	trees := pie.Ints{}
	for _, s := range *ss {
		trees = trees.Append(s.Run().TreesEncountered())
	}
	return &trees
}
