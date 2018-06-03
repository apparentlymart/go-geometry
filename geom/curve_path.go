package geom

// CubicCurveSeq represents a sequence of connected cubic bezier curves.
//
// The internal representation is a slice of points: a start point, followed
// by three more points for each curve, where each curve is implied to start
// at the endpoint before it.
type CubicCurveSeq []Point

// BeginCubicCurveSeq creates an invalid CubicCurveSeq that has only a start
// point, ready to be grown with method Append.
//
// The given cap value should be the number of expected curve segments in
// the final sequence, which will be used to reserve capacity in the sequence's
// underlying backing array.
func BeginCubicCurveSeq(start Point, cap int) CubicCurveSeq {
	rCap := 1 + cap*3
	ret := make(CubicCurveSeq, 1, rCap)
	ret[0] = start
	return ret
}

// Append adds a line segment to the receiver, returning the updated sequence.
// The result may share a backing array with the receiver, if sufficient
// capacity is available in that array.
func (s CubicCurveSeq) Append(c1, c2, end Point) CubicCurveSeq {
	return append(s, c1, c2, end)
}

// Iterator returns an interator over the curve segments in the receiving
// sequence.
func (s CubicCurveSeq) Iterator() CubicCurveIterator {
	return &cubicCurveSeqIter{
		seq: s,
		pos: -1,
	}
}

// CubicCurveIterator is a stateful iterator over a sequence of cubic curve
// segments.
type CubicCurveIterator interface {
	Next() bool
	CubicCurve() CubicCurve
}

type cubicCurveSeqIter struct {
	seq CubicCurveSeq
	pos int
}

func (i *cubicCurveSeqIter) Next() bool {
	i.pos += 3
	return i.pos <= (len(i.seq) - 1)
}

func (i *cubicCurveSeqIter) CubicCurve() CubicCurve {
	s := i.seq[i.pos:]
	var ret CubicCurve
	copy(ret[:], s)
	return ret
}
