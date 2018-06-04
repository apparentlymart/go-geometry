package svgpath

func (t tokenType) String() string {
	switch t {
	case tokenInst:
		return "instruction"
	case tokenArg:
		return "number"
	case tokenComma:
		return "comma"
	case tokenEnd:
		return "end of string"
	default:
		return "invalid token"
	}
}
