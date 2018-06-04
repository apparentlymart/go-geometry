package svgpath

import (
	"github.com/apparentlymart/go-geometry/geom"
)

// Path represents an SVG path, consisting of a sequence of drawing commands.
type Path []Command

// Equal returns true if the reciever is identical to the given other path.
func (p Path) Equal(o Path) bool {
	if len(p) != len(o) {
		return false
	}
	for i := range p {
		if !p[i].Equal(o[i]) {
			return false
		}
	}
	return true
}

// MakeAbsolute rewrites any relative steps in the path to be absolute, in place.
//
// This method tracks the effect of each command on the sub-path start point
// and the previous end point to calculate the effective absolute coordinates
// of any relative commands.
func (p Path) MakeAbsolute() {
	start := geom.Origin
	prev := geom.Origin

	for i, cmd := range p {
		newCmd, newPrev := cmd.ToAbsolute(start, prev)
		prev = newPrev
		if newCmd.Inst == MoveTo {
			start = newPrev
		}
		p[i] = newCmd
	}
}

// Subpaths separates all of the sub-paths of the receiver into separate paths.
//
// The resulting paths are not necessarily self-contained, since a sub-path
// may begin relative to its predecessor. Call MakeAbsolute first (before
// calling Subpaths) to ensure that the resulting sub-paths are self-contained.
func (p Path) Subpaths() []Path {
	if len(p) == 0 {
		return nil
	}

	// Count how many sub-paths we will have.
	pc := 0
	prevInst := ClosePath
	for _, c := range p {
		switch c.Inst.ToAbsolute() {
		case ClosePath, MoveTo:
			if prevInst != ClosePath {
				pc++
			}
		}
	}

	if pc == 0 {
		pc = 1 // always at least one, even if MoveTo and Close are omitted
	}

	ret := make([]Path, 0, pc)
	s := 0
	var i int
	for i = 0; i < len(p); i++ {
		switch p[i].Inst.ToAbsolute() {
		case ClosePath:
			ret = append(ret, p[s:i+1])
			s = i + 1
		case MoveTo:
			if s != i {
				ret = append(ret, p[s:i])
				s = i
			}
		}
	}
	if i > s {
		ret = append(ret, p[s:i])
	}
	return ret
}
