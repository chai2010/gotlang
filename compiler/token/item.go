package token

import "fmt"

// Token represents a token or text string returned from the scanner.
type Token struct {
	Typ  TokType // The type of this item.
	Pos  int     // The starting position, in bytes, of this item in the input string.
	Val  string  // The value of this item.
	Line int     // The line number at the start of this item.
}

func (i Token) String() string {
	switch {
	case i.Typ == ItemEOF:
		return "EOF"
	case i.Typ == ItemError:
		return i.Val
	case i.Typ > itemKeyword:
		return fmt.Sprintf("<%s>", i.Val)
	case len(i.Val) > 10:
		return fmt.Sprintf("%.10q...", i.Val)
	}
	return fmt.Sprintf("%q", i.Val)
}
