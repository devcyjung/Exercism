package pangram

func IsPangram(input string) (b bool) {
	m := make(map[rune]bool)
    for _, v := range input {
        if 'A' <= v && v <= 'Z' || 'a' <= v && v <= 'z' {
            m[v] = true
        }
    }
    for offset := rune(0); offset < 26; offset++ {
        if !m['A' + offset] && !m['a' + offset] {
            return
        }
    }
    b = true
    return
}
