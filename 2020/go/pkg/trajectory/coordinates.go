package trajectory

// Coordinates represent the x & y position on a plane
type Coordinates struct {
	x int
	y int
}

// NewCoordinates returns a pointer to a single Coordinates
// from two integers, x & y
func NewCoordinates(x, y int) *Coordinates {
	return &Coordinates{x, y}
}

// NewCoordinateses returns a slice of Coordinates pointers
// from one or more slices of integers (integer pairs) - example:
//  [][]int{{1,2},{3,4}} --> []*Coordinates{{1,2},{3,4}}
func NewCoordinateses(is ...[]int) []*Coordinates {
	var cs = []*Coordinates{}
	for _, i := range is {
		if len(i) == 2 {
			cs = append(cs, &Coordinates{
				x: i[0],
				y: i[1],
			})
		}
	}
	return cs
}
