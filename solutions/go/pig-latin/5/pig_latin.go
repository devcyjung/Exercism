package piglatin

import (
    "fmt"
    str "strings"
    uni "unicode"
    "unicode/utf8"
)

func Sentence(sentence string) string {
	result := &str.Builder{}
    var spaces, substr, cons, vowel string
    var first rune
    for len(sentence) > 0 {
        first, _ = utf8.DecodeRuneInString(sentence)
        if uni.IsSpace(first) {
            spaces, sentence = spaceWord(sentence)
            fmt.Fprint(result, spaces)
            continue
        }
        substr, sentence = wordSpace(sentence)
        cons, vowel = consVowel(substr)
        if len(cons) == 0 || str.HasPrefix(cons, "xr") || str.HasPrefix(cons, "yt") {
            fmt.Fprint(result, substr, "ay")
            continue
        }
        ypos := str.Index(cons, "y")
        if ypos > 0 {
            fmt.Fprint(result, substr[ypos:], substr[:ypos], "ay")
            continue
        }
        if str.HasSuffix(cons, "q") && str.HasPrefix(vowel, "u") {
            fmt.Fprint(result, vowel[1:], cons, vowel[:1], "ay")
            continue
        }
        fmt.Fprint(result, vowel, cons, "ay")
    }
    return fmt.Sprint(result)
}

func sepBy(substr string, sepAt func(rune) bool) (string, string) {
    sep := str.IndexFunc(substr, sepAt)
    if sep == -1 {
        sep = len(substr)
    }
    return substr[:sep], substr[sep:]
}

func spaceWord(substr string) (string, string) {
    return sepBy(substr, isWord)
}

func isWord(r rune) bool {
    return !uni.IsSpace(r)
}

func wordSpace(substr string) (string, string) {
    return sepBy(substr, uni.IsSpace)
}

func consVowel(substr string) (string, string) {
    return sepBy(substr, isVowel)
}

func isVowel(r rune) bool {
	switch uni.ToLower(r) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
    default:
        return false
	}
}