package svgpath

import (
	"github.com/apparentlymart/go-geometry/geom"
)

type Path []Command

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
