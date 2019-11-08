// 模板词法扫描器.
//
// 扫描结果没有处理文本的开头和结尾的空白, 也没有检查小括弧是否匹配.
// 更严格的检查或者是预处理在语法解析部分完成.
package scanner

import (
	"fmt"
	"strings"

	"github.com/chai2010/gotlang/compiler/token"
)

const (
	spaceChars      = " \t\r\n" // 空白字符
	leftTrimMarker  = "- "      // 左分隔符, 删除前面记号的尾部空白字符
	rightTrimMarker = " -"      // 右分隔符, 删除后面记号的开头空白字符
)

// 解析全部记号
func LexAll(text, leftDelim, rightDelim string) ([]Item, error) {
	var allTokens []Item
	for pos := 0; pos < len(text); {
		if strings.HasPrefix(text[pos:], leftDelim) {
			// 解析一个Action, 包含一组记号
			if toks, err := lexAction(text, leftDelim, rightDelim, pos); err == nil {
				allTokens = append(allTokens, toks...)
				pos = toks[len(toks)-1].End
			} else {
				return nil, err
			}
		} else {
			// 解析文本
			if tok, err := lexText(text, leftDelim, pos); err == nil {
				allTokens = append(allTokens, tok)
				pos = tok.End
			} else {
				return nil, err
			}
		}
	}

	// 检查 () 是否匹配
	for i := 0; i < len(allTokens); i++ {
		if allTokens[i].Typ == token.TokenLeftDelim {
			var parenDepth = 0
			for j := i; j < len(allTokens); j++ {
				if allTokens[j].Typ == token.TokenLeftParen {
					parenDepth++
				}
				if allTokens[j].Typ == token.TokenRightParen {
					parenDepth--
					if parenDepth < 0 {
						return allTokens, fmt.Errorf("unexpected right paren, at %v", allTokens[j].Pos)
					}
				}
				if allTokens[j].Typ == token.TokenRightDelim {
					i = j
					break
				}
			}
			if parenDepth != 0 {
				return allTokens, fmt.Errorf("unclosed left paren, at %v", allTokens[i].Pos)
			}
		}
	}

	// 处理收尾空白字符截断
	// {{23 -}} < {{- 45}} => 23<45
	for i, tok := range allTokens {
		if tok.Typ == token.TokenLeftDelim {
			if i > 0 && strings.HasSuffix(tok.Val, leftTrimMarker) {
				allTokens[i-1].Val = strings.TrimRight(allTokens[i-1].Val, spaceChars)
			}
		}
		if tok.Typ == token.TokenRightDelim {
			if i < len(allTokens)-1 && strings.HasPrefix(tok.Val, rightTrimMarker) {
				allTokens[i+1].Val = strings.TrimLeft(allTokens[i+1].Val, spaceChars)
			}
		}
	}

	// OK
	return allTokens, nil
}

// 读取开始的文本, 直到遇到左分隔符'{{'
func lexText(text, leftDelim string, pos int) (tok Item, err error) {
	// 查找左分隔符
	if x := strings.Index(text[pos:], leftDelim); x >= 0 {
		tok = Item{Typ: token.TokenText, Val: text[pos:x], Pos: pos, End: x}
		return
	}

	// 没有找到左分隔符, 说明已经到末尾文本
	tok = Item{Typ: token.TokenText, Val: text[pos:], Pos: pos, End: len(text)}
	return
}

// 从指定位置开始解析一个动作的一组记号
// 返回记号列表, 每个记号和消耗的字节数, 或者是错误
// 因为依赖上下文, 这里不支持处理收尾空白字符截断, 需要在外层手工处理
func lexAction(text, leftDelim, rightDelim string, pos int) (toks []Item, err error) {
	if leftDelim == "" {
		leftDelim = "{{"
	}
	if rightDelim == "" {
		rightDelim = "}}"
	}

	// 1. 读取左分隔符
	if tok, err := lexLeftDelim(text, leftDelim, rightDelim, pos); err == nil {
		toks = append(toks, tok)
		pos = tok.End
	} else {
		return toks, err
	}

	// 左右小括弧深度
	var parenDepth = 0

	// 2. 循环处理Action内部记号, 直到右分隔符/文件结束或错误
Loop:
	for pos < len(text) {
		switch r := text[pos]; true {
		case strings.HasPrefix(text[pos:], "}}"):
			// 右分隔符, 结束循环
			tok := Item{Typ: token.TokenRightDelim, Pos: pos, End: pos + 1}
			toks = append(toks, tok)
			pos = tok.End
			break Loop

		case strings.HasPrefix(text[pos:], "/*"):
			// 跳过注释
			if i := strings.Index(text[pos:], "*/"); i >= 0 {
				pos = i
			} else {
				return nil, fmt.Errorf("unclosed comment, at %d", pos)
			}

		case isSpace(text[pos]):
			// 跳过空白字符
			for i := pos; true; i++ {
				if i >= len(text) || !isSpace(text[i]) {
					pos = i
					break
				}
			}

		case r == '=':
			tok := Item{Typ: token.TokenAssign, Pos: pos, End: pos + 1}
			toks = append(toks, tok)
			pos = tok.End

		case strings.HasPrefix(text[pos:], ":="):
			// return l.errorf("expected :=")

			tok := Item{Typ: token.TokenDeclare, Pos: pos, End: pos + 2}
			toks = append(toks, tok)
			pos = tok.End

		case r == '|':
			tok := Item{Typ: token.TokenPipe, Pos: pos, End: pos + 1}
			toks = append(toks, tok)
			pos = tok.End

		case r == '"':
		case r == '`':
		case r == '\'':
		case r == '.':

		case r == '+' || r == '-' || ('0' <= r && r <= '9'):
		case isAlphaNumer(r):
		case r == '(':
			parenDepth++
		case r == ')':
			parenDepth--
			if parenDepth < 0 {
				return nil, fmt.Errorf("unexpected right paren, at %d", pos)
			}
		case r == '$':
		case r == ',':
			// {{$i, $v := range m}}
			// itemChar 只有 ',' 有意义, 可以简化
		default:
			return nil, fmt.Errorf("unrecognized character in action: at %v", pos)
		}
	}

	// 检查小括弧匹配
	if parenDepth != 0 {
		return nil, fmt.Errorf("unclosed left paren, at %d", pos)
	}

	// 检查右分隔符
	if toks[len(toks)-1].Typ != token.TokenRightDelim {
		return nil, fmt.Errorf("expected %q, at %v", rightDelim, pos)
	}

	// OK
	return toks, nil
}

func lexLeftDelim(text, leftDelim, rightDelim string, pos int) (Item, error) {
	panic("todo")
}
func lexRightDelim(text, leftDelim, rightDelim string, pos int) (Item, error) {
	panic("todo")
}

// lexInsideAction_Space
// lexInsideAction_Identifier
// lexInsideAction_Field
// lexInsideAction_Char
// lexInsideAction_Number
// lexInsideAction_Quote
// lexInsideAction_RawQuote

func lexComment(text, leftDelim, rightDelim string, pos int) (tok Item, err error) {
	panic("todo")
}
