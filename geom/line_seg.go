package geom

// LineSeg is a line segment represented by its two endpoints.
type LineSeg [2]Point

// LineSeg returns a copy of the receiver. This method exists only to implement
// LineSegger.
func (s LineSeg) LineSeg() LineSeg {
	return s
}

// CubicCurve returns a cubic bezier curve equivalent to the receiving line
// segment.
func (c LineSeg) CubicCurve() CubicCurve {
	return CubicCurve{
		c[0],
		c[0], c[1],
		c[1],
	}
}

// A LineSegger can convert itself into a line segment.
type LineSegger interface {
	LineSeg() LineSeg
}
