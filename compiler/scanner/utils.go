package scanner

import "unicode"

// 空白字符(不含换行符号)
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// 行尾符号
func isEOL(r rune) bool {
	return r == '\r' || r == '\n'
}

// 是否为字面或数字(包含下划线)
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
