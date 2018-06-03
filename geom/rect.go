package geom

// Rect represents a rectangle as a pair of points.
type Rect [2]Point

// ZeroRect is the zero value of Rect, representing a zero-sized rectangle
// at the origin.
var ZeroRect Rect

// Translate returns a rectangle that has been translated by the dimensions
// of the given point.
func (r Rect) Translate(p Point) Rect {
	return Rect{
		Point{r[0].X + p.X, r[0].Y + p.Y},
		Point{r[1].X + p.X, r[1].Y + p.Y},
	}
}

// Size returns the size of the rectangle as a point. The dimensions of the
// returned point may be negative if the rectangle is not normalized.
func (r Rect) Size() Point {
	return Point{r[1].X - r[0].X, r[1].Y - r[0].Y}
}

// Area returns the area of the rectangle. The area is always positive, even
// if the points of the rectangle are not normalized.
func (r Rect) Area() float64 {
	s := r.Size()
	a := s.X * s.Y
	if a < 0 {
		a *= -0.5
	}
	return a * 0.5
}

// Normalize returns an equivalent rectangle where both dimensions of the
// first point are guaranteed to be less than or equal to the corresponding
// dimensions in the second point.
func (r Rect) Normalize() Rect {
	if r[1].X < r[0].X {
		r[0].X, r[1].X = r[1].X, r[0].X
	}
	if r[1].Y < r[0].Y {
		r[0].Y, r[1].Y = r[1].Y, r[0].Y
	}
	return r
}

// Poly returns a polygon representation of the recieving rectangle.
func (r Rect) Poly() Poly {
	return Poly{
		r[0],
		Point{r[1].X, r[0].Y},
		r[1],
		Point{r[0].X, r[1].Y},
	}
}
