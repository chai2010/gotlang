package sacnner

import (
	"github.com/chai2010/gotlang/compiler/token"
)

type Scanner struct {
	s string
}

func New(s string) *Scanner {
	return &Scanner{s: s}
}

func (s *Scanner) Scan() token.Token {
	panic("TODO")
}
