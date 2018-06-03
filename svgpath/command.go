package svgpath

import (
	"fmt"

	"github.com/apparentlymart/go-geometry/geom"
)

// Command represents a single command in a path.
type Command struct {
	// Instruction identifies the action taken by the command.
	Inst Instruction

	// Args are the arguments to the command. The length of this slice depends
	// on the instruction, and methods on a command may panic if Args is not
	// the correct length for the given instruction.
	Args []float64
}

// Equal returns true if the receiver is identical to the given other command.
func (c Command) Equal(o Command) bool {
	if c.Inst != o.Inst {
		return false
	}
	if len(c.Args) != len(o.Args) {
		return false
	}
	for i := range c.Args {
		if c.Args[i] != o.Args[i] {
			return false
		}
	}
	return true
}

// ToAbsolute converts a relative command into its absolute equivalent, using
// the given points as the path start and previous endpoint respectively.
//
// The results are a new command that is absolute and the new point that might
// be used as the start point for the subsequent command in a path.
//
// If the receiver is already absolute, an identical copy will be returned
// along with the endpoint it refers to.
func (c Command) ToAbsolute(start, prev geom.Point) (Command, geom.Point) {
	retInst := c.Inst.ToAbsolute()
	switch c.Inst {
	case MoveTo, LineTo:
		return c, c.endpoint()
	case MoveToRel, LineToRel:
		end := start.Add(c.endpoint())
		return Command{retInst, pointCoords(end)}, end
	case ClosePath, ClosePathRel:
		return Command{retInst, nil}, start
	default:
		panic(fmt.Sprintf("AsAbsolute with invalid instruction %s", c.Inst))
	}
}

// endpoint returns the endpoint for a command where one can be determined.
// For the horizontal and vertical line instructions, the specified component
// will be zero. For the "close path" instructions, the result is the origin.
// This method is private because of these surprising behaviors.
func (c Command) endpoint() geom.Point {
	a := c.Args

	switch c.Inst {
	case MoveTo, MoveToRel, LineTo, LineToRel:
		return geom.Point{a[0], a[1]}
	case ClosePath, ClosePathRel:
		return geom.Origin
	default:
		panic(fmt.Sprintf("endpoint for invalid instruction %s", c.Inst))
	}
}

func (c Command) GoString() string {
	switch c.Inst {
	case MoveTo:
		return fmt.Sprintf("svgpath.Move(%#v)", c.Args)
	case MoveToRel:
		return fmt.Sprintf("svgpath.MoveRel(%#v)", c.Args)
	case LineTo:
		return fmt.Sprintf("svgpath.Line(%#v)", c.Args)
	case LineToRel:
		return fmt.Sprintf("svgpath.LineRel(%#v)", c.Args)
	case ClosePath:
		return fmt.Sprintf("svgpath.Close")
	case ClosePathRel:
		return fmt.Sprintf("svgpath.CloseRel")
	default:
		return fmt.Sprintf("svgpath.Command{svgpath.%s, %#v}", c.Inst, c.Args)
	}
}

func pointCoords(p geom.Point) []float64 {
	return []float64{p.X, p.Y}
}

type Instruction byte

const (
	MoveTo               Instruction = 'M'
	MoveToRel            Instruction = 'm'
	ClosePath            Instruction = 'Z'
	ClosePathRel         Instruction = 'z'
	LineTo               Instruction = 'L'
	LineToRel            Instruction = 'l'
	HorizLineTo          Instruction = 'H'
	HorizLineToRel       Instruction = 'h'
	VertLineTo           Instruction = 'V'
	VertLineToRel        Instruction = 'v'
	CurveTo              Instruction = 'C'
	CurveToRel           Instruction = 'c'
	SmoothCurveTo        Instruction = 'S'
	SmoothCurveToRel     Instruction = 's'
	QuadCurveTo          Instruction = 'Q'
	QuadCurveToRel       Instruction = 'q'
	SmoothQuadCurveTo    Instruction = 'T'
	SmoothQuadCurveToRel Instruction = 't'
	ArcTo                Instruction = 'A'
	ArcToRel             Instruction = 'a'
)

//go:generate stringer -type Instruction

func (i Instruction) Absolute() bool {
	return (i & 0x20) == 0
}

func (i Instruction) Relative() bool {
	return (i & 0x20) != 0
}

// ToAbsolute converts the receving instruction to its absolute equivalent,
// if it isn't already absolute.
func (i Instruction) ToAbsolute() Instruction {
	return i & 0xdf
}

// ToRelative converts the receving instruction to its relative equivalent,
// if it isn't already relative.
func (i Instruction) ToRelative() Instruction {
	return i | 0x20
}

var instructionSyms = map[byte]Instruction{
	'M': MoveTo,
	'm': MoveToRel,
	'Z': ClosePath,
	'z': ClosePathRel,
	'L': LineTo,
	'l': LineToRel,
	'H': HorizLineTo,
	'h': HorizLineToRel,
	'V': VertLineTo,
	'v': VertLineToRel,
	'C': CurveTo,
	'c': CurveToRel,
	'S': SmoothCurveTo,
	's': SmoothCurveToRel,
	'Q': QuadCurveTo,
	'q': QuadCurveToRel,
	'T': SmoothQuadCurveTo,
	't': SmoothCurveToRel,
	'A': ArcTo,
	'a': ArcToRel,
}
