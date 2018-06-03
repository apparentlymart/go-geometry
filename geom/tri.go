package geom

// Tri represents a triangle using the coordinates of its three vertices.
type Tri [3]Point

// Poly returns the equivalent polygon to the recieving triangle.
func (t Tri) Poly() Poly {
	copy := t // Ensure that poly updates can't affect our tri
	return Poly(copy[:])
}

// Facing returns either 1 or -1 depending on the ordering of the points.
// Assuming an X axis that increases to the right and a Y axis that increases
// upward, facing is 1 if the points are clockwise, and -1 for anti-clockwise.
//
// If the polygon is self-intersecting then the result is not meaningful.
func (t Tri) Facing() int {
	return Poly(t[:]).Facing()
}

// Area returns the area of the triangle.
func (t Tri) Area() float64 {
	return Poly(t[:]).Area()
}
