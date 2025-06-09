package acronym

import (
    "strings"
    "unicode"
)

func Abbreviate(s string) string {
    var b strings.Builder
    for _, field := range strings.Fields(strings.Map(sanitizer, s)) {
        b.WriteString(field[:1])
    }
    return b.String()
}

func sanitizer(r rune) rune {
    switch {
    case 'A' <= r && r <= 'Z':
        return r
    case 'a' <= r && r <= 'z':
        return 'A' + r - 'a'
    case '0' <= r && r <= '9':
        return r
    case r == '-' || unicode.IsSpace(r):
        return ' '
    default:
        return -1
    }
}