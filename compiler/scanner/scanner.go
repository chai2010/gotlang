package sacnner

type Token int

type Scanner struct {
}

func New(s string) *Scanner {
	return &Scanner{}
}

func (s *Scanner) Scan() (pos int, tok Token, lit string) {
	panic("TODO")
}
