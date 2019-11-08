package sacnner

type Scanner struct {
	s string
}

func New(s string) *Scanner {
	return &Scanner{s: s}
}

func (s *Scanner) Scan() Item {
	panic("TODO")
}
