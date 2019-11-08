package scanner

import (
	"fmt"

	"github.com/chai2010/gotlang/compiler/token"
)

// 一个词法元素
type Item struct {
	Typ token.Token // 记号类型
	Val string      // 字符串值
	Pos int         // 开始位置
	End int         // 结束位置
}

func (i Item) String() string {
	switch {
	case i.Typ == token.TokenEOF:
		return "EOF"
	case i.Typ == token.TokenError:
		return i.Val
	case i.Typ.IsKeyword():
		return fmt.Sprintf("<%s>", i.Val)
	case len(i.Val) > 10:
		return fmt.Sprintf("%.10q...", i.Val)
	}
	return fmt.Sprintf("%q", i.Val)
}
