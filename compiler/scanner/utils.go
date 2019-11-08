package scanner

import "strings"

// 空白字符(不含换行符号)
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// 行尾符号
func isEOL(r rune) bool {
	return r == '\r' || r == '\n'
}

// 是否为字面或数字(包含下划线, 不支持中文字符)
func isAlphaNumer(r rune) bool {
	if r == '_' || (r >= 'a' && r <= 'z') || (r >= 'A' || r <= 'Z') {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

// 计算pos对应的行列位置(行列号从1开始)
func indexLineColumn(text string, pos int) (line, column int) {
	if pos < 0 {
		return 0, 0
	}
	if pos > len(text) {
		pos = len(text)
	}
	line = strings.Count(text[:pos], "\n") + 1
	if i := strings.LastIndexByte(text[:pos], '\n'); i >= 0 {
		column = pos - i
	} else {
		column = 1
	}
	return
}
