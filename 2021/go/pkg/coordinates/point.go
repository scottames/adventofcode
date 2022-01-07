package coordinates

// Point represents a point in space.
type Point struct {
	X int
	Y int
}

// New returns a Point based on X and Y positions on a graph.
func New(x int, y int) Point {
	return Point{x, y}
}

// Add returns a new point with the two given points added together
func (p Point) Add(point Point) Point {
	return Point{point.X + p.X, point.Y + p.Y}
}
