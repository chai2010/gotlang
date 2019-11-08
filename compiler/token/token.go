package token

import "strconv"

// 词法记号类型
type Token int

const (
	// 特殊标记
	TokenError Token = iota // 零值, 错误类型
	TokenEOF                // 文件结尾

	// 面值常量
	TokenBool         // boolean constant
	TokenChar         // printable ASCII character; grab bag for comma etc.
	TokenCharConstant // character constant
	TokenComplex      // complex constant (1+2i); imaginary is just a number
	TokenNumber       // simple number, including imaginary
	TokenString       // quoted string (includes quotes)
	TokenRawString    // raw quoted string (includes quotes)
	TokenText         // plain text

	// 运算符
	TokenAssign     // =
	TokenDeclare    // :=
	TokenLeftParen  // (
	TokenRightParen // )
	TokenPipe       // |
	TokenLeftDelim  // {{
	TokenRightDelim // }}
	TokenSpace      // 空白分隔符, 类似其它语言的分号';'

	// 标识符
	TokenIdentifier // 标识符, 一般是内置的函数
	TokenField      // 成员, 例如: .Field, .Field1.Field2
	TokenVariable   // 变量, $开头, 例如: $, $x, $hello

	// 关键字
	TokenKeyword_begin // 关键字开始
	TokenDot           // .
	TokenBlock         // block
	TokenDefine        // define
	TokenElse          // else
	TokenEnd           // end
	TokenIf            // if
	TokenNil           // nil
	TokenRange         // range
	TokenTemplate      // template
	TokenWith          // with
	TokenKeyword_end   // 关键字结束
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
