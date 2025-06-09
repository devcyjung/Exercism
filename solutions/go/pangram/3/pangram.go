package pangram

import "strings"

func IsPangram(input string) bool {
    if len(input) < 26 {
        return false
    }
	m := make(map[rune]bool)
    for _, v := range strings.ToLower(input) {
        if 'a' <= v && v <= 'z' {
            m[v] = true
        }
    }
    var ok bool
    for r := 'a'; r <= 'z'; r++ {
        if ok = m[r]; !ok {
            return ok
        }
    }
    return ok
}
