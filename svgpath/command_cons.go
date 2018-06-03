package svgpath

import (
	"github.com/apparentlymart/go-geometry/geom"
)

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

var Close = Command{
	Inst: ClosePath,
}

var CloseRel = Command{
	Inst: ClosePathRel,
}
