package isogram

import "unicode"

func IsIsogram(word string) bool {
	m := make(map[rune]bool)
    for _, letter := range word {
        if unicode.IsSpace(letter) || letter == '-' {
            continue
        }
        lower := unicode.ToLower(letter)
        if m[lower] {
            return false
        }
        m[lower] = true
    }
    return true
}
