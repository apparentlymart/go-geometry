package svgpath

import (
	"fmt"
)

// Parse interprets the given string as SVG Path data, returning the result
// as a Path.
//
// The commands in the string are passed through as-is, without further
// interpretation. However, if a particular command has multiple sets of
// arguments these are normalized to separate commands, ensuring that each
// element of the resulting path represents only one drawing command.
//
// If the string is not valid per the grammar, an error is returned along with
// any commands that were successfully parsed so far.
func Parse(s string) (Path, error) {
	var path Path

	sc := &scanner{
		Remain: s,
	}

	// First we'll pass over the string to count how many commands and
	// arguments it seems to have, so we can pre-allocate our arrays for these.
	cmdC := 0
	argC := 0
	for sc.hasMoreTokens() {
		tok := sc.read()
		switch tok.Type {
		case tokenArg:
			argC++
		case tokenInst:
			cmdC++
		}
	}

	// reset
	*sc = scanner{
		Remain: s,
	}

	// We put all of our args together in a single buffer to avoid lots of
	// small heap allocations as we parse.
	var args []float64
	if argC != 0 {
		args = make([]float64, 0, argC)
	}

	// We also try to pre-allocate our path, but it's not an exact science
	// because an instruction can actually introduce multiple commands.
	if cmdC != 0 {
		path = make([]Command, 0, cmdC)
	}

	for sc.hasMoreTokens() {
		inst, err := sc.reqInst()
		if err != nil {
			return path, err
		}

		switch inst.ToAbsolute() {
		case MoveTo, LineTo, SmoothQuadCurveTo:
			for {
				pt, err := sc.reqPoint()
				if err != nil {
					return path, err
				}
				args = append(args, pt.X, pt.Y)
				path = append(path, Command{inst, args})
				args = args[2:]
				if !sc.nextIsArg() {
					break
				}
				sc.skipComma()
				switch inst {
				case MoveTo:
					inst = LineTo
				case MoveToRel:
					inst = LineToRel
				}
			}
		case HorizLineTo, VertLineTo:
			for {
				v, err := sc.reqNumber()
				if err != nil {
					return path, err
				}
				args = append(args, v)
				path = append(path, Command{inst, args})
				args = args[1:]
				if !sc.nextIsArg() {
					break
				}
				sc.skipComma()
			}
		case CurveTo:
			for {
				cpt1, err := sc.reqPoint()
				if err != nil {
					return path, err
				}
				sc.skipComma()
				cpt2, err := sc.reqPoint()
				if err != nil {
					return path, err
				}
				sc.skipComma()
				ept, err := sc.reqPoint()
				if err != nil {
					return path, err
				}
				args = append(args, cpt1.X, cpt1.Y)
				args = append(args, cpt2.X, cpt2.Y)
				args = append(args, ept.X, ept.Y)
				path = append(path, Command{inst, args})
				args = args[6:]
				if !sc.nextIsArg() {
					break
				}
				sc.skipComma()
			}
		case SmoothCurveTo, QuadCurveTo:
			for {
				cpt, err := sc.reqPoint()
				if err != nil {
					return path, err
				}
				sc.skipComma()
				ept, err := sc.reqPoint()
				if err != nil {
					return path, err
				}
				args = append(args, cpt.X, cpt.Y)
				args = append(args, ept.X, ept.Y)
				path = append(path, Command{inst, args})
				args = args[4:]
				if !sc.nextIsArg() {
					break
				}
				sc.skipComma()
			}
		case ArcTo:
			for {
				r, err := sc.reqPoint()
				if err != nil {
					return path, err
				}
				sc.skipComma()
				xRot, err := sc.reqNumber()
				if err != nil {
					return path, err
				}
				sc.skipComma()
				largeArc, err := sc.reqFlag()
				if err != nil {
					return path, err
				}
				sc.skipComma()
				sweep, err := sc.reqFlag()
				if err != nil {
					return path, err
				}
				sc.skipComma()
				ept, err := sc.reqPoint()
				if err != nil {
					return path, err
				}
				args = append(args, r.X, r.Y)
				args = append(args, xRot, largeArc, sweep)
				args = append(args, ept.X, ept.Y)
				path = append(path, Command{inst, args})
				args = args[7:]
				if !sc.nextIsArg() {
					break
				}
				sc.skipComma()
			}
		case ClosePath:
			path = append(path, Command{inst, nil})
		default:
			// Should never happen, because the above should be exhaustive
			// for everything in instructoinSyms.
			panic(fmt.Sprintf("Parse can't handle instruction %s", inst))
		}
	}

	return path, nil
}
