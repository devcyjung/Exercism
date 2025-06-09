package acronym

func Abbreviate(s string) string {
    var ignore bool
    var runes []rune
    for i, ch := range s {
        if isAlnum(ch) {
            if ignore {
                continue
            }
            runes = append(runes, toCap(ch))
            ignore = true
        } else {
            if ch == '\'' {
                t := []rune(s)
                if i != len(t)-1 && isAlnum(t[i-1]) && isAlnum(t[i+1]) {
                    ignore = true
                    continue
                }
            }
            ignore = false
        }
    }
	return string(runes)
}

func isAlnum(ch rune) bool {
    return 0 <= ch && ch <= 9 || 'A' <= ch && ch <= 'Z' || 'a' <= ch && ch <= 'z' 
}

func toCap(ch rune) rune {
    if 'a' <= ch && ch <= 'z' {
        return ch - 'a' + 'A'
    }
    return ch
}
