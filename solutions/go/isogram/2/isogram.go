package isogram

import "unicode"

func IsIsogram(word string) bool {
	m := make(map[rune]struct{})
    for _, letter := range word {
        if unicode.IsSpace(letter) || letter == '-' {
            continue
        }
        lower := unicode.ToLower(letter)
        if _, ok := m[lower]; ok {
            return false
        }
        m[lower] = struct{}{}
    }
    return true
}