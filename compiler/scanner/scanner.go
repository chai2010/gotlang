// 模板词法扫描器.
//
// 扫描结果没有处理文本的开头和结尾的空白, 也没有检查小括弧是否匹配.
// 更严格的检查或者是预处理在语法解析部分完成.
package scanner

import (
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
		// 解析文本
		if tok, err := LexText(text, leftDelim, pos); err == nil {
			allTokens = append(allTokens, tok)
			pos = tok.End
		} else {
			return nil, err
		}

		// 解析一个Action, 包含一组记号
		if toks, err := LexAction(text, leftDelim, rightDelim, pos); err == nil {
			allTokens = append(allTokens, toks...)
			pos = toks[len(toks)-1].End
		} else {
			return nil, err
		}
	}

	// 处理收尾空白字符截断(可在AST部分处理)
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

	// TODO: 检查 () 是否匹配
	// 可在语法解析部分检查

	// OK
	return allTokens, nil
}

// 读取开始的文本, 直到遇到左分隔符'{{'
func LexText(text, leftDelim string, pos int) (tok Item, err error) {
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
func LexAction(text, leftDelim, rightDelim string, pos int) (toks []Item, err error) {
	if leftDelim == "" {
		leftDelim = "{{"
	}
	if rightDelim == "" {
		rightDelim = "}}"
	}

	// 1. 读取左分隔符
	// 2. a:注释(是否吃掉右分隔符?); b:内部记号
	// 3. 读取右分隔符

	// 到达文件末尾
	// 内部函数不用检查末尾表姐
	if atEOF(text, pos) {
		toks = []Item{{Typ: token.TokenEOF, Pos: pos}}
		return
	}

	// 读取一个记号
	// 尝试每个记号读取
	// 最后还没有读取，则是错误

	// {{}}内部和外部的语法是不太一样的, 因此需要简单的上下文支持

	panic("todo")
}

// 是否到底文件末尾
func atEOF(text string, pos int) bool {
	return pos >= len(text)
}

// 读取Action内部的一个元素
// 遇到右分隔符/文件末尾/错误等返回
func lexInsideActionItem(text, leftDelim, rightDelim string, pos int) (tok Item, err error) {
	panic("todo")

	// 分类处理
	// 不认识则错误

	// EOF
	// 空白
	// =
	// :=
	// |
	// "
	// `
	// '
	// .
	// + - [0-9]
	// isalnum
	// (
	// )
	// isprint
	// error
}

// lexLeftDelim
// lexRightDelim
// lexSpace
// lexIdentifier
// lexField
// lexChar
// lexNumber
// lexQuote
// lexRawQuote

func lexCommnet(text, leftDelim, rightDelim string, pos int) (tok Item, err error) {
	panic("todo")
}
