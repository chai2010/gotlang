package token

import "fmt"

// Item represents a token or text string returned from the scanner.
type Item struct {
	typ  ItemType // The type of this item.
	pos  Pos      // The starting position, in bytes, of this item in the input string.
	val  string   // The value of this item.
	line int      // The line number at the start of this item.
}

func (i Item) String() string {
	switch {
	case i.typ == ItemEOF:
		return "EOF"
	case i.typ == ItemError:
		return i.val
	case i.typ > itemKeyword:
		return fmt.Sprintf("<%s>", i.val)
	case len(i.val) > 10:
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}
