package geom

// Point represents a single 2D point.
type Point struct {
	X, Y float64
}

// Origin is the zero point, with both coordinates set to zero.
var Origin Point

// Add returns the sum of the receiver and the given other point.
func (p Point) Add(o Point) Point {
	return Point{p.X + o.X, p.Y + o.Y}
}

// Sub returns the difference between the receiver and the given other point.
func (p Point) Sub(o Point) Point {
	return Point{p.X - o.X, p.Y - o.Y}
}

// Scale multiplies the point by the given scale factor.
func (p Point) Scale(m float64) Point {
	return Point{p.X * m, p.Y * m}
}

// Mul multiplies each dimension in the receiver by the corresponding dimension
// in the given point, creating a non-uniform scale.
func (p Point) Mul(o Point) Point {
	return Point{p.X * o.X, p.Y * o.Y}
}
