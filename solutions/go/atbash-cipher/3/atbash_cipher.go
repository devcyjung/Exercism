package atbash

import (
    "strings"
    "unicode"
)

func Atbash(s string) string {
	var b strings.Builder
    var nonSpaceLen int
    var isLower, isUpper, isDigit bool
	for _, v := range s {
        isLower, isUpper, isDigit = unicode.IsLower(v), unicode.IsUpper(v), unicode.IsDigit(v)
        if !isLower && !isUpper && !isDigit {
            continue
        }
        if nonSpaceLen > 0 && nonSpaceLen % 5 == 0 {
            b.WriteRune(' ')
        }
        nonSpaceLen++
        switch {
        case isLower:
            b.WriteRune('z' - v + 'a')
        case isUpper:
            b.WriteRune('Z' - v + 'a')
        case isDigit:
            b.WriteRune(v)
        }
    }
    return b.String()
}