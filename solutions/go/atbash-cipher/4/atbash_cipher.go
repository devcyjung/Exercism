package atbash

import (
    "strings"
)

func Atbash(s string) string {
    var b strings.Builder
    for i, atbash := range strings.Map(atbashMapping, s) {
        if i > 0 && i % 5 == 0 {
            b.WriteRune(' ')
        }
        b.WriteRune(atbash)
    }
    return b.String()
}

func atbashMapping(r rune) rune {
    switch {
    case 'a' <= r && r <= 'z':
        return 'a' + 'z' - r
    case 'A' <= r && r <= 'Z':
        return 'a' + 'Z' - r
    case '0' <= r && r <= '9':
        return r
    default:
        return -1
    }
}