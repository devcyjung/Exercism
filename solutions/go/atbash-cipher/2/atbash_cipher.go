package atbash

import "strings"

func Atbash(s string) string {
	var b strings.Builder
    var size int
	for _, v := range s {
        if !('a' <= v && v <= 'z') && !('A' <= v && v <= 'Z') && !('0' <= v && v<= '9') {
            continue
        }
        if size > 0 && size % 5 == 0 {
            b.WriteRune(' ')
        }
        size++
        if 'a' <= v && v <= 'z' {
            b.WriteRune('z' - v + 'a')
        } else if 'A' <= v && v <= 'Z' {
            b.WriteRune('Z' - v + 'a')
        } else {
            b.WriteRune(v)
        }
    }
    return b.String()
}
