package geom

// Poly represents a closed polygon as a sequence of vertices. Each vertex
// is connected to the next by an edge, and the final vertex is implicitly
// connected to the first.
//
// A polygon with self-intersecting edges is not well-formed and will produce
// meaningless results from most operations.
//
// The zero value of Poly is an invalid polygon.
type Poly []Point

// Poly returns a copy of the receiver. This method is present only to
// implement Polyer.
func (p Poly) Poly() Poly {
	r := make(Poly, len(p))
	copy(r, p)
	return r
}

// Area returns the area enclosed by the polygon.
func (p Poly) Area() float64 {
	a := p.dirArea()
	if a < 0.0 {
		return a * -1.0
	}
	return a
}

// Facing returns either 1 or -1 depending on the ordering of the points.
// Assuming an X axis that increases to the right and a Y axis that increases
// upward, facing is 1 if the points are clockwise, and -1 for anti-clockwise.
func (p Poly) Facing() int {
	a := p.dirArea()
	if a < 0.0 {
		return -1
	}
	return 1
}

func (p Poly) dirArea() float64 {
	l := len(p)
	var s float64
	for i := 0; i < l; i++ {
		j := (i + 1) % l
		s += (p[j].X - p[i].X) * (p[j].Y + p[i].Y)
	}
	return s
}

// A Polyer is a shape that can express itself as a polygon.
type Polyer interface {
	// Poly returns a polygon representation of the receiver.
	//
	// It always returns a copy of the original shape, so that the polygon
	// can be mutated independently.
	Poly() Poly
}
