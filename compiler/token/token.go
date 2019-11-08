package token

import "strconv"

// 词法记号类型
type Token int

const (
	TokenError        Token = iota // error occurred; value is text of error
	TokenBool                      // boolean constant
	TokenChar                      // printable ASCII character; grab bag for comma etc.
	TokenCharConstant              // character constant
	TokenComplex                   // complex constant (1+2i); imaginary is just a number
	TokenAssign                    // equals ('=') introducing an assignment
	TokenDeclare                   // colon-equals (':=') introducing a declaration
	TokenEOF
	TokenField      // alphanumeric identifier starting with '.'
	TokenIdentifier // alphanumeric identifier not starting with '.'
	TokenLeftDelim  // left action delimiter
	TokenLeftParen  // '(' inside action
	TokenNumber     // simple number, including imaginary
	TokenPipe       // pipe symbol
	TokenRawString  // raw quoted string (includes quotes)
	TokenRightDelim // right action delimiter
	TokenRightParen // ')' inside action
	TokenSpace      // run of spaces separating arguments
	TokenString     // quoted string (includes quotes)
	TokenText       // plain text
	TokenVariable   // variable starting with '$', such as '$' or  '$1' or '$hello'

	// Keywords appear after all the rest.
	TokenKeyword_begin // used only to delimit the keywords
	TokenBlock         // block keyword
	TokenDot           // the cursor, spelled '.'
	TokenDefine        // define keyword
	TokenElse          // else keyword
	TokenEnd           // end keyword
	TokenIf            // if keyword
	TokenNil           // the untyped nil constant, easiest to treat as a keyword
	TokenRange         // range keyword
	TokenTemplate      // template keyword
	TokenWith          // with keyword
	TokenKeyword_end   // used only to delimit the keywords
)

var tokens = [...]string{
	TokenError:        "TokenError",
	TokenBool:         "TokenBool",
	TokenChar:         "TokenChar",
	TokenCharConstant: "TokenCharConstant",
	TokenComplex:      "TokenComplex",
	TokenAssign:       "TokenAssign",
	TokenDeclare:      "TokenDeclare",
	TokenEOF:          "TokenEOF",
	TokenField:        "TokenField",
	TokenIdentifier:   "TokenIdentifier",
	TokenLeftDelim:    "TokenLeftDelim",
	TokenLeftParen:    "TokenLeftParen",
	TokenNumber:       "TokenNumber",
	TokenPipe:         "TokenPipe",
	TokenRawString:    "TokenRawString",
	TokenRightDelim:   "TokenRightDelim",
	TokenRightParen:   "TokenRightParen",
	TokenSpace:        "TokenSpace",
	TokenString:       "TokenString",
	TokenText:         "TokenText",
	TokenVariable:     "TokenVariable",
	TokenBlock:        "TokenBlock",
	TokenDot:          "TokenDot",
	TokenDefine:       "TokenDefine",
	TokenElse:         "TokenElse",
	TokenEnd:          "TokenEnd",
	TokenIf:           "TokenIf",
	TokenNil:          "TokenNil",
	TokenRange:        "TokenRange",
	TokenTemplate:     "TokenTemplate",
	TokenWith:         "TokenWith",
}

var keywords = map[string]Token{
	".":        TokenDot,
	"block":    TokenBlock,
	"define":   TokenDefine,
	"else":     TokenElse,
	"end":      TokenEnd,
	"if":       TokenIf,
	"range":    TokenRange,
	"nil":      TokenNil,
	"template": TokenTemplate,
	"with":     TokenWith,
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

func (tok Token) IsKeyword() bool {
	return tok > TokenKeyword_begin && tok < TokenKeyword_end
}

func IsKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

// 查找关键字
func Lookup(ident string) Token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return TokenIdentifier
}
