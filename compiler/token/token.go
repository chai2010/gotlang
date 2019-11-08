package token

import "strconv"

// TokType identifies the type of lex items.
type TokType int

const (
	ItemError        TokType = iota // error occurred; value is text of error
	ItemBool                        // boolean constant
	ItemChar                        // printable ASCII character; grab bag for comma etc.
	ItemCharConstant                // character constant
	ItemComplex                     // complex constant (1+2i); imaginary is just a number
	ItemAssign                      // equals ('=') introducing an assignment
	ItemDeclare                     // colon-equals (':=') introducing a declaration
	ItemEOF
	ItemField      // alphanumeric identifier starting with '.'
	ItemIdentifier // alphanumeric identifier not starting with '.'
	ItemLeftDelim  // left action delimiter
	ItemLeftParen  // '(' inside action
	ItemNumber     // simple number, including imaginary
	ItemPipe       // pipe symbol
	ItemRawString  // raw quoted string (includes quotes)
	ItemRightDelim // right action delimiter
	ItemRightParen // ')' inside action
	ItemSpace      // run of spaces separating arguments
	ItemString     // quoted string (includes quotes)
	ItemText       // plain text
	ItemVariable   // variable starting with '$', such as '$' or  '$1' or '$hello'

	// Keywords appear after all the rest.
	itemKeyword // used only to delimit the keywords

	ItemBlock    // block keyword
	ItemDot      // the cursor, spelled '.'
	ItemDefine   // define keyword
	ItemElse     // else keyword
	ItemEnd      // end keyword
	ItemIf       // if keyword
	ItemNil      // the untyped nil constant, easiest to treat as a keyword
	ItemRange    // range keyword
	ItemTemplate // template keyword
	ItemWith     // with keyword
)

// TODO
var tokens = [...]string{}

var key = map[string]TokType{
	".":        ItemDot,
	"block":    ItemBlock,
	"define":   ItemDefine,
	"else":     ItemElse,
	"end":      ItemEnd,
	"if":       ItemIf,
	"range":    ItemRange,
	"nil":      ItemNil,
	"template": ItemTemplate,
	"with":     ItemWith,
}

func (tok TokType) String() string {
	s := ""
	if 0 <= tok && tok < TokType(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

func (tok TokType) IsKeyword() bool {
	return tok > itemKeyword
}

func IsKeyword(name string) bool {
	return false
}

func Lookup(ident string) TokType {
	return ItemEOF // todo
}
