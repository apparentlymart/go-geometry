package svgpath

type Command struct {
	Inst Instruction
	Args []float64
}

type Instruction byte

const (
	Move                 Instruction = 'M'
	MoveRel              Instruction = 'm'
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

// Absolute converts the receving instruction to its absolute equivalent,
// if it isn't already absolute.
func (i Instruction) Absolute() Instruction {
	return i & 0xdf
}

// Relative converts the receving instruction to its relative equivalent,
// if it isn't already relative.
func (i Instruction) Relative() Instruction {
	return i | 0x20
}
