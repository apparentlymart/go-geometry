package svgpath

import (
	"fmt"
	"strconv"

	"github.com/apparentlymart/go-geometry/geom"
)

type token struct {
	Type  tokenType
	Chars string
}

type tokenType rune

const (
	tokenNone  tokenType = 0
	tokenInst  tokenType = 'I'
	tokenArg   tokenType = 'A'
	tokenComma tokenType = ','
	tokenEnd   tokenType = '␄'
	tokenBad   tokenType = '�'
)

type scanner struct {
	Remain string
	Peeked token
}

func (s *scanner) peek() token {
	if s.Peeked.Type == tokenNone {
		s.Peeked = s.next()
	}
	return s.Peeked
}

func (s *scanner) read() token {
	ret := s.peek()
	if ret.Type != tokenEnd {
		s.Peeked = token{}
	}
	return ret
}

func (s *scanner) nextIsArg() bool {
	p := s.peek().Type
	return p == tokenArg || p == tokenComma
}

func (s *scanner) skipComma() {
	if s.peek().Type == tokenComma {
		s.read()
	}
}

func (s *scanner) hasMoreTokens() bool {
	return s.peek().Type != tokenEnd
}

func (s *scanner) reqNumber() (float64, error) {
	tok := s.peek()
	if tok.Type != tokenArg {
		return 0, fmt.Errorf("expecting number but got %s", tok.Type)
	}
	s.read()
	ret, err := strconv.ParseFloat(tok.Chars, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number")
	}
	return ret, nil
}

func (s *scanner) reqPoint() (geom.Point, error) {
	x, err := s.reqNumber()
	if err != nil {
		return geom.Origin, err
	}
	if s.peek().Type == tokenComma {
		s.read()
	}
	y, err := s.reqNumber()
	if err != nil {
		return geom.Origin, err
	}
	return geom.Point{x, y}, nil
}

func (s *scanner) reqInst() (Instruction, error) {
	tok := s.peek()
	if tok.Type != tokenInst {
		return 0, fmt.Errorf("expecting instruction but got %s", tok.Type)
	}
	s.read()
	return Instruction(tok.Chars[0]), nil
}

func (s *scanner) reqFlag() (float64, error) {
	tok := s.peek()
	if tok.Type != tokenArg {
		return 0, fmt.Errorf("expecting flag but got %s", tok.Type)
	}
	s.read()
	switch tok.Chars {
	case "0":
		return 0, nil
	case "1":
		return 1, nil
	default:
		return 0, fmt.Errorf("flag argument must be either '0' or '1'")
	}
}

func (s *scanner) next() token {
	// Trim off any leading whitespace first
	for {
		if len(s.Remain) == 0 {
			return token{tokenEnd, ""}
		}

		next := s.Remain[0]
		if next == 0x20 || next == 0x09 || next == 0x0d || next == 0x0a {
			s.Remain = s.Remain[1:]
			continue
		}
		break
	}

	next := s.Remain[0]
	switch {
	case instructionSyms[next] != 0:
		inst := s.Remain[0:1]
		s.Remain = s.Remain[1:]
		return token{tokenInst, inst}

	case next == ',':
		chars := s.Remain[0:1]
		s.Remain = s.Remain[1:]
		return token{tokenComma, chars}

	case canStartArg(next):
		ret := s.nextArg()
		return ret

	default:
		bad := s.Remain
		s.Remain = ""
		return token{tokenBad, bad}
	}
}

func (s *scanner) nextArg() token {
	first := s.Remain[0]
	var l int

	seenDot := false
	seenExp := false
	inExp := false
	seenDigit := true
	if first == '-' || first == '+' {
		seenDigit = false
	}

Chars:
	for l = 1; l < len(s.Remain); l++ {
		next := s.Remain[l]
		switch next {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			seenDigit = true
			inExp = false
		case '.':
			if seenDot {
				break Chars
			}
			seenDot = true
			seenDigit = false // must have at least one more digit
			inExp = false
		case 'e', 'E':
			if seenExp || !seenDigit {
				break Chars
			}
			seenExp = true
			inExp = true
			seenDot = true    // no dots allowed after exp
			seenDigit = false // must have at least one more digit
		case '-', '+':
			if !inExp {
				break Chars
			}
			inExp = false
		default:
			break Chars
		}
	}

	chars := s.Remain[:l]
	s.Remain = s.Remain[len(chars):]
	if !seenDigit {
		return token{tokenBad, chars}
	}
	return token{tokenArg, chars}
}

func canStartArg(c byte) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	case '+', '-', '.':
		return true
	default:
		return false
	}
}
