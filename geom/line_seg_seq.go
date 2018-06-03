package geom

// LineSegSeq represents a sequence of connected line segments.
//
// The internal representation is a slice of points.
type LineSegSeq []Point

// BeginLineSegSeq creates an invalid LineSegSeq that has only a start
// point, ready to be grown with method Append.
//
// The given cap value should be the number of expected line segments in
// the final sequence, which will be used to reserve capacity in the sequence's
// underlying backing array.
func BeginLineSegSeq(start Point, cap int) LineSegSeq {
	rCap := 1 + cap
	ret := make(LineSegSeq, 1, rCap)
	ret[0] = start
	return ret
}

// Append adds a line segment to the receiver, returning the updated sequence.
// The result may share a backing array with the receiver, if sufficient
// capacity is available in that array.
func (s LineSegSeq) Append(next Point) LineSegSeq {
	return append(s, next)
}

// Iterator returns an interator over the line segments in the receiving
// sequence.
func (s LineSegSeq) Iterator() LineSegIterator {
	return &lineSegSeqIter{
		seq:    s,
		pos:    -1,
		closed: false,
	}
}

// LineSegIterator is a stateful iterator over a sequence of line segments.
type LineSegIterator interface {
	Next() bool
	LineSeg() LineSeg
}

type lineSegSeqIter struct {
	seq    LineSegSeq
	pos    int
	closed bool
}

func (i *lineSegSeqIter) Next() bool {
	i.pos++
	if i.closed {
		// one more iteration, then
		return i.pos <= len(i.seq)
	}
	return i.pos <= (len(i.seq) - 1)
}

func (i *lineSegSeqIter) LineSeg() LineSeg {
	si := i.pos
	ei := (si + 1) % len(i.seq)
	return LineSeg{i.seq[si], i.seq[ei]}
}
