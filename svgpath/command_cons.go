package svgpath

import (
	"github.com/apparentlymart/go-geometry/geom"
)

var Close = Command{
	Inst: ClosePath,
}

var CloseRel = Command{
	Inst: ClosePathRel,
}

func Move(new geom.Point) Command {
	return Command{
		Inst: MoveTo,
		Args: pointCoords(new),
	}
}

func MoveRel(new geom.Point) Command {
	return Command{
		Inst: MoveToRel,
		Args: pointCoords(new),
	}
}

func Line(new geom.Point) Command {
	return Command{
		Inst: LineTo,
		Args: pointCoords(new),
	}
}

func LineRel(new geom.Point) Command {
	return Command{
		Inst: LineToRel,
		Args: pointCoords(new),
	}
}

func HorizLine(x float64) Command {
	return Command{
		Inst: HorizLineTo,
		Args: []float64{x},
	}
}

func HorizLineRel(length float64) Command {
	return Command{
		Inst: HorizLineToRel,
		Args: []float64{length},
	}
}

func VertLine(y float64) Command {
	return Command{
		Inst: VertLineTo,
		Args: []float64{y},
	}
}

func VertLineRel(length float64) Command {
	return Command{
		Inst: VertLineToRel,
		Args: []float64{length},
	}
}

func Curve(c1, c2, end geom.Point) Command {
	return Command{
		Inst: CurveTo,
		Args: []float64{
			c1.X, c1.Y,
			c2.X, c2.Y,
			end.X, end.Y,
		},
	}
}

func CurveRel(c1, c2, end geom.Point) Command {
	return Command{
		Inst: CurveToRel,
		Args: []float64{
			c1.X, c1.Y,
			c2.X, c2.Y,
			end.X, end.Y,
		},
	}
}

func SmoothCurve(c2, end geom.Point) Command {
	return Command{
		Inst: SmoothCurveTo,
		Args: []float64{
			c2.X, c2.Y,
			end.X, end.Y,
		},
	}
}

func SmoothCurveRel(c2, end geom.Point) Command {
	return Command{
		Inst: SmoothCurveToRel,
		Args: []float64{
			c2.X, c2.Y,
			end.X, end.Y,
		},
	}
}

func QuadCurve(c, end geom.Point) Command {
	return Command{
		Inst: QuadCurveTo,
		Args: []float64{
			c.X, c.Y,
			end.X, end.Y,
		},
	}
}

func QuadCurveRel(c, end geom.Point) Command {
	return Command{
		Inst: QuadCurveToRel,
		Args: []float64{
			c.X, c.Y,
			end.X, end.Y,
		},
	}
}

func SmoothQuadCurve(end geom.Point) Command {
	return Command{
		Inst: SmoothQuadCurveTo,
		Args: []float64{
			end.X, end.Y,
		},
	}
}

func SmoothQuadCurveRel(end geom.Point) Command {
	return Command{
		Inst: SmoothQuadCurveToRel,
		Args: []float64{
			end.X, end.Y,
		},
	}
}

func Arc(radii geom.Point, xRot float64, largeArc, sweep bool, end geom.Point) Command {
	largeArcArg, sweepArg := 0.0, 0.0
	if largeArc {
		largeArcArg = 1.0
	}
	if sweep {
		sweepArg = 1.0
	}
	return Command{
		Inst: ArcTo,
		Args: []float64{
			radii.X, radii.Y,
			xRot,
			largeArcArg, sweepArg,
			end.X, end.Y,
		},
	}
}

func ArcRel(radii geom.Point, xRot float64, largeArc, sweep bool, end geom.Point) Command {
	largeArcArg, sweepArg := 0.0, 0.0
	if largeArc {
		largeArcArg = 1.0
	}
	if sweep {
		sweepArg = 1.0
	}
	return Command{
		Inst: ArcToRel,
		Args: []float64{
			radii.X, radii.Y,
			xRot,
			largeArcArg, sweepArg,
			end.X, end.Y,
		},
	}
}
