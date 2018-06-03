package geom

const twoThirds = 2 / 3

// CubicCurve represents a cubic bezier curve.
type CubicCurve [4]Point

// CubicCurve returns a copy of the reciever. This method is present only to
// implement CubicCurver.
func (c CubicCurve) CubicCurve() CubicCurve {
	return c
}

// QuadraticCurve represents a quadratic bezier curve.
type QuadraticCurve [3]Point

// CubicCurve returns a CubicCurve equivalent to the receiver.
func (c QuadraticCurve) CubicCurve() CubicCurve {
	return CubicCurve{
		c[0],
		c[0].Add(c[1].Sub(c[0]).Scale(twoThirds)),
		c[0].Add(c[1].Sub(c[2]).Scale(twoThirds)),
		c[2],
	}
}

// A CubicCurver can convert itself into a CubicCurve.
type CubicCurver interface {
	CubicCurve() CubicCurve
}
