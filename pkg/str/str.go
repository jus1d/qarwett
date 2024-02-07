package str

import "unicode"

// Capitalize capitalizes first char of string.
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
