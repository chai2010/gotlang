// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package sacnner

import (
	"fmt"
	"testing"
)

// Make the types prettyprint.
var itemName = map[TokType]string{
	ItemError:        "error",
	ItemBool:         "bool",
	ItemChar:         "char",
	ItemCharConstant: "charconst",
	ItemComplex:      "complex",
	ItemDeclare:      ":=",
	ItemEOF:          "EOF",
	ItemField:        "field",
	ItemIdentifier:   "identifier",
	ItemLeftDelim:    "left delim",
	ItemLeftParen:    "(",
	ItemNumber:       "number",
	ItemPipe:         "pipe",
	ItemRawString:    "raw string",
	ItemRightDelim:   "right delim",
	ItemRightParen:   ")",
	ItemSpace:        "space",
	ItemString:       "string",
	ItemVariable:     "variable",

	// keywords
	ItemDot:      ".",
	ItemBlock:    "block",
	ItemDefine:   "define",
	ItemElse:     "else",
	ItemIf:       "if",
	ItemEnd:      "end",
	ItemNil:      "nil",
	ItemRange:    "range",
	ItemTemplate: "template",
	ItemWith:     "with",
}

func (i TokType) String() string {
	s := itemName[i]
	if s == "" {
		return fmt.Sprintf("item%d", int(i))
	}
	return s
}

type lexTest struct {
	name  string
	input string
	items []Token
}

func mkItem(typ TokType, text string) Token {
	return Token{
		Typ: typ,
		Val: text,
	}
}

var (
	tDot        = mkItem(ItemDot, ".")
	tBlock      = mkItem(ItemBlock, "block")
	tEOF        = mkItem(ItemEOF, "")
	tFor        = mkItem(ItemIdentifier, "for")
	tLeft       = mkItem(ItemLeftDelim, "{{")
	tLpar       = mkItem(ItemLeftParen, "(")
	tPipe       = mkItem(ItemPipe, "|")
	tQuote      = mkItem(ItemString, `"abc \n\t\" "`)
	tRange      = mkItem(ItemRange, "range")
	tRight      = mkItem(ItemRightDelim, "}}")
	tRpar       = mkItem(ItemRightParen, ")")
	tSpace      = mkItem(ItemSpace, " ")
	raw         = "`" + `abc\n\t\" ` + "`"
	rawNL       = "`now is{{\n}}the time`" // Contains newline inside raw quote.
	tRawQuote   = mkItem(ItemRawString, raw)
	tRawQuoteNL = mkItem(ItemRawString, rawNL)
)

var lexTests = []lexTest{
	{"empty", "", []Token{tEOF}},
	{"spaces", " \t\n", []Token{mkItem(ItemText, " \t\n"), tEOF}},
	{"text", `now is the time`, []Token{mkItem(ItemText, "now is the time"), tEOF}},
	{"text with comment", "hello-{{/* this is a comment */}}-world", []Token{
		mkItem(ItemText, "hello-"),
		mkItem(ItemText, "-world"),
		tEOF,
	}},
	{"punctuation", "{{,@% }}", []Token{
		tLeft,
		mkItem(ItemChar, ","),
		mkItem(ItemChar, "@"),
		mkItem(ItemChar, "%"),
		tSpace,
		tRight,
		tEOF,
	}},
	{"parens", "{{((3))}}", []Token{
		tLeft,
		tLpar,
		tLpar,
		mkItem(ItemNumber, "3"),
		tRpar,
		tRpar,
		tRight,
		tEOF,
	}},
	{"empty action", `{{}}`, []Token{tLeft, tRight, tEOF}},
	{"for", `{{for}}`, []Token{tLeft, tFor, tRight, tEOF}},
	{"block", `{{block "foo" .}}`, []Token{
		tLeft, tBlock, tSpace, mkItem(ItemString, `"foo"`), tSpace, tDot, tRight, tEOF,
	}},
	{"quote", `{{"abc \n\t\" "}}`, []Token{tLeft, tQuote, tRight, tEOF}},
	{"raw quote", "{{" + raw + "}}", []Token{tLeft, tRawQuote, tRight, tEOF}},
	{"raw quote with newline", "{{" + rawNL + "}}", []Token{tLeft, tRawQuoteNL, tRight, tEOF}},
	{"numbers", "{{1 02 0x14 0X14 -7.2i 1e3 1E3 +1.2e-4 4.2i 1+2i 1_2 0x1.e_fp4 0X1.E_FP4}}", []Token{
		tLeft,
		mkItem(ItemNumber, "1"),
		tSpace,
		mkItem(ItemNumber, "02"),
		tSpace,
		mkItem(ItemNumber, "0x14"),
		tSpace,
		mkItem(ItemNumber, "0X14"),
		tSpace,
		mkItem(ItemNumber, "-7.2i"),
		tSpace,
		mkItem(ItemNumber, "1e3"),
		tSpace,
		mkItem(ItemNumber, "1E3"),
		tSpace,
		mkItem(ItemNumber, "+1.2e-4"),
		tSpace,
		mkItem(ItemNumber, "4.2i"),
		tSpace,
		mkItem(ItemComplex, "1+2i"),
		tSpace,
		mkItem(ItemNumber, "1_2"),
		tSpace,
		mkItem(ItemNumber, "0x1.e_fp4"),
		tSpace,
		mkItem(ItemNumber, "0X1.E_FP4"),
		tRight,
		tEOF,
	}},
	{"characters", `{{'a' '\n' '\'' '\\' '\u00FF' '\xFF' '本'}}`, []Token{
		tLeft,
		mkItem(ItemCharConstant, `'a'`),
		tSpace,
		mkItem(ItemCharConstant, `'\n'`),
		tSpace,
		mkItem(ItemCharConstant, `'\''`),
		tSpace,
		mkItem(ItemCharConstant, `'\\'`),
		tSpace,
		mkItem(ItemCharConstant, `'\u00FF'`),
		tSpace,
		mkItem(ItemCharConstant, `'\xFF'`),
		tSpace,
		mkItem(ItemCharConstant, `'本'`),
		tRight,
		tEOF,
	}},
	{"bools", "{{true false}}", []Token{
		tLeft,
		mkItem(ItemBool, "true"),
		tSpace,
		mkItem(ItemBool, "false"),
		tRight,
		tEOF,
	}},
	{"dot", "{{.}}", []Token{
		tLeft,
		tDot,
		tRight,
		tEOF,
	}},
	{"nil", "{{nil}}", []Token{
		tLeft,
		mkItem(ItemNil, "nil"),
		tRight,
		tEOF,
	}},
	{"dots", "{{.x . .2 .x.y.z}}", []Token{
		tLeft,
		mkItem(ItemField, ".x"),
		tSpace,
		tDot,
		tSpace,
		mkItem(ItemNumber, ".2"),
		tSpace,
		mkItem(ItemField, ".x"),
		mkItem(ItemField, ".y"),
		mkItem(ItemField, ".z"),
		tRight,
		tEOF,
	}},
	{"keywords", "{{range if else end with}}", []Token{
		tLeft,
		mkItem(ItemRange, "range"),
		tSpace,
		mkItem(ItemIf, "if"),
		tSpace,
		mkItem(ItemElse, "else"),
		tSpace,
		mkItem(ItemEnd, "end"),
		tSpace,
		mkItem(ItemWith, "with"),
		tRight,
		tEOF,
	}},
	{"variables", "{{$c := printf $ $hello $23 $ $var.Field .Method}}", []Token{
		tLeft,
		mkItem(ItemVariable, "$c"),
		tSpace,
		mkItem(ItemDeclare, ":="),
		tSpace,
		mkItem(ItemIdentifier, "printf"),
		tSpace,
		mkItem(ItemVariable, "$"),
		tSpace,
		mkItem(ItemVariable, "$hello"),
		tSpace,
		mkItem(ItemVariable, "$23"),
		tSpace,
		mkItem(ItemVariable, "$"),
		tSpace,
		mkItem(ItemVariable, "$var"),
		mkItem(ItemField, ".Field"),
		tSpace,
		mkItem(ItemField, ".Method"),
		tRight,
		tEOF,
	}},
	{"variable invocation", "{{$x 23}}", []Token{
		tLeft,
		mkItem(ItemVariable, "$x"),
		tSpace,
		mkItem(ItemNumber, "23"),
		tRight,
		tEOF,
	}},
	{"pipeline", `intro {{echo hi 1.2 |noargs|args 1 "hi"}} outro`, []Token{
		mkItem(ItemText, "intro "),
		tLeft,
		mkItem(ItemIdentifier, "echo"),
		tSpace,
		mkItem(ItemIdentifier, "hi"),
		tSpace,
		mkItem(ItemNumber, "1.2"),
		tSpace,
		tPipe,
		mkItem(ItemIdentifier, "noargs"),
		tPipe,
		mkItem(ItemIdentifier, "args"),
		tSpace,
		mkItem(ItemNumber, "1"),
		tSpace,
		mkItem(ItemString, `"hi"`),
		tRight,
		mkItem(ItemText, " outro"),
		tEOF,
	}},
	{"declaration", "{{$v := 3}}", []Token{
		tLeft,
		mkItem(ItemVariable, "$v"),
		tSpace,
		mkItem(ItemDeclare, ":="),
		tSpace,
		mkItem(ItemNumber, "3"),
		tRight,
		tEOF,
	}},
	{"2 declarations", "{{$v , $w := 3}}", []Token{
		tLeft,
		mkItem(ItemVariable, "$v"),
		tSpace,
		mkItem(ItemChar, ","),
		tSpace,
		mkItem(ItemVariable, "$w"),
		tSpace,
		mkItem(ItemDeclare, ":="),
		tSpace,
		mkItem(ItemNumber, "3"),
		tRight,
		tEOF,
	}},
	{"field of parenthesized expression", "{{(.X).Y}}", []Token{
		tLeft,
		tLpar,
		mkItem(ItemField, ".X"),
		tRpar,
		mkItem(ItemField, ".Y"),
		tRight,
		tEOF,
	}},
	{"trimming spaces before and after", "hello- {{- 3 -}} -world", []Token{
		mkItem(ItemText, "hello-"),
		tLeft,
		mkItem(ItemNumber, "3"),
		tRight,
		mkItem(ItemText, "-world"),
		tEOF,
	}},
	{"trimming spaces before and after comment", "hello- {{- /* hello */ -}} -world", []Token{
		mkItem(ItemText, "hello-"),
		mkItem(ItemText, "-world"),
		tEOF,
	}},
	// errors
	{"badchar", "#{{\x01}}", []Token{
		mkItem(ItemText, "#"),
		tLeft,
		mkItem(ItemError, "unrecognized character in action: U+0001"),
	}},
	{"unclosed action", "{{\n}}", []Token{
		tLeft,
		mkItem(ItemError, "unclosed action"),
	}},
	{"EOF in action", "{{range", []Token{
		tLeft,
		tRange,
		mkItem(ItemError, "unclosed action"),
	}},
	{"unclosed quote", "{{\"\n\"}}", []Token{
		tLeft,
		mkItem(ItemError, "unterminated quoted string"),
	}},
	{"unclosed raw quote", "{{`xx}}", []Token{
		tLeft,
		mkItem(ItemError, "unterminated raw quoted string"),
	}},
	{"unclosed char constant", "{{'\n}}", []Token{
		tLeft,
		mkItem(ItemError, "unterminated character constant"),
	}},
	{"bad number", "{{3k}}", []Token{
		tLeft,
		mkItem(ItemError, `bad number syntax: "3k"`),
	}},
	{"unclosed paren", "{{(3}}", []Token{
		tLeft,
		tLpar,
		mkItem(ItemNumber, "3"),
		mkItem(ItemError, `unclosed left paren`),
	}},
	{"extra right paren", "{{3)}}", []Token{
		tLeft,
		mkItem(ItemNumber, "3"),
		tRpar,
		mkItem(ItemError, `unexpected right paren U+0029 ')'`),
	}},

	// Fixed bugs
	// Many elements in an action blew the lookahead until
	// we made lexInsideAction not loop.
	{"long pipeline deadlock", "{{|||||}}", []Token{
		tLeft,
		tPipe,
		tPipe,
		tPipe,
		tPipe,
		tPipe,
		tRight,
		tEOF,
	}},
	{"text with bad comment", "hello-{{/*/}}-world", []Token{
		mkItem(ItemText, "hello-"),
		mkItem(ItemError, `unclosed comment`),
	}},
	{"text with comment close separated from delim", "hello-{{/* */ }}-world", []Token{
		mkItem(ItemText, "hello-"),
		mkItem(ItemError, `comment ends before closing delimiter`),
	}},
	// This one is an error that we can't catch because it breaks templates with
	// minimized JavaScript. Should have fixed it before Go 1.1.
	{"unmatched right delimiter", "hello-{.}}-world", []Token{
		mkItem(ItemText, "hello-{.}}-world"),
		tEOF,
	}},
}

// collect gathers the emitted items into a slice.
func collect(t *lexTest, left, right string) (items []Token) {
	l := lex(t.name, t.input, left, right)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.Typ == ItemEOF || item.Typ == ItemError {
			break
		}
	}
	return
}

func equal(i1, i2 []Token, checkPos bool) bool {
	if len(i1) != len(i2) {
		return false
	}
	for k := range i1 {
		if i1[k].Typ != i2[k].Typ {
			return false
		}
		if i1[k].Val != i2[k].Val {
			return false
		}
		if checkPos && i1[k].Pos != i2[k].Pos {
			return false
		}
		if checkPos && i1[k].Line != i2[k].Line {
			return false
		}
	}
	return true
}

func TestLex(t *testing.T) {
	for _, test := range lexTests {
		items := collect(&test, "", "")
		if !equal(items, test.items, false) {
			t.Errorf("%s: got\n\t%+v\nexpected\n\t%v", test.name, items, test.items)
		}
	}
}

// Some easy cases from above, but with delimiters $$ and @@
var lexDelimTests = []lexTest{
	{"punctuation", "$$,@%{{}}@@", []Token{
		tLeftDelim,
		mkItem(ItemChar, ","),
		mkItem(ItemChar, "@"),
		mkItem(ItemChar, "%"),
		mkItem(ItemChar, "{"),
		mkItem(ItemChar, "{"),
		mkItem(ItemChar, "}"),
		mkItem(ItemChar, "}"),
		tRightDelim,
		tEOF,
	}},
	{"empty action", `$$@@`, []Token{tLeftDelim, tRightDelim, tEOF}},
	{"for", `$$for@@`, []Token{tLeftDelim, tFor, tRightDelim, tEOF}},
	{"quote", `$$"abc \n\t\" "@@`, []Token{tLeftDelim, tQuote, tRightDelim, tEOF}},
	{"raw quote", "$$" + raw + "@@", []Token{tLeftDelim, tRawQuote, tRightDelim, tEOF}},
}

var (
	tLeftDelim  = mkItem(ItemLeftDelim, "$$")
	tRightDelim = mkItem(ItemRightDelim, "@@")
)

func TestDelims(t *testing.T) {
	for _, test := range lexDelimTests {
		items := collect(&test, "$$", "@@")
		if !equal(items, test.items, false) {
			t.Errorf("%s: got\n\t%v\nexpected\n\t%v", test.name, items, test.items)
		}
	}
}

var lexPosTests = []lexTest{
	{"empty", "", []Token{{ItemEOF, 0, "", 1}}},
	{"punctuation", "{{,@%#}}", []Token{
		{ItemLeftDelim, 0, "{{", 1},
		{ItemChar, 2, ",", 1},
		{ItemChar, 3, "@", 1},
		{ItemChar, 4, "%", 1},
		{ItemChar, 5, "#", 1},
		{ItemRightDelim, 6, "}}", 1},
		{ItemEOF, 8, "", 1},
	}},
	{"sample", "0123{{hello}}xyz", []Token{
		{ItemText, 0, "0123", 1},
		{ItemLeftDelim, 4, "{{", 1},
		{ItemIdentifier, 6, "hello", 1},
		{ItemRightDelim, 11, "}}", 1},
		{ItemText, 13, "xyz", 1},
		{ItemEOF, 16, "", 1},
	}},
	{"trimafter", "{{x -}}\n{{y}}", []Token{
		{ItemLeftDelim, 0, "{{", 1},
		{ItemIdentifier, 2, "x", 1},
		{ItemRightDelim, 5, "}}", 1},
		{ItemLeftDelim, 8, "{{", 2},
		{ItemIdentifier, 10, "y", 2},
		{ItemRightDelim, 11, "}}", 2},
		{ItemEOF, 13, "", 2},
	}},
	{"trimbefore", "{{x}}\n{{- y}}", []Token{
		{ItemLeftDelim, 0, "{{", 1},
		{ItemIdentifier, 2, "x", 1},
		{ItemRightDelim, 3, "}}", 1},
		{ItemLeftDelim, 6, "{{", 2},
		{ItemIdentifier, 10, "y", 2},
		{ItemRightDelim, 11, "}}", 2},
		{ItemEOF, 13, "", 2},
	}},
}

// The other tests don't check position, to make the test cases easier to construct.
// This one does.
func TestPos(t *testing.T) {
	for _, test := range lexPosTests {
		items := collect(&test, "", "")
		if !equal(items, test.items, true) {
			t.Errorf("%s: got\n\t%v\nexpected\n\t%v", test.name, items, test.items)
			if len(items) == len(test.items) {
				// Detailed print; avoid item.String() to expose the position value.
				for i := range items {
					if !equal(items[i:i+1], test.items[i:i+1], true) {
						i1 := items[i]
						i2 := test.items[i]
						t.Errorf("\t#%d: got {%v %d %q %d} expected {%v %d %q %d}",
							i, i1.Typ, i1.Pos, i1.Val, i1.Line, i2.Typ, i2.Pos, i2.Val, i2.Line)
					}
				}
			}
		}
	}
}
