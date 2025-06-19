package piglatin

import (
    "strings"
    "unicode"
    "unicode/utf8"
)

func Sentence(sentence string) string {
	var b strings.Builder
    for len(sentence) > 0 {
        first, _ := utf8.DecodeRuneInString(sentence)
        if unicode.IsSpace(first) {
            var spaces string
            spaces, sentence = spaceWord(sentence)
            b.WriteString(spaces)
            continue
        }
        var substr string
        substr, sentence = wordSpace(sentence)
        cons, vowel := consVowel(substr)
        if len(cons) == 0 || strings.HasPrefix(cons, "xr") || strings.HasPrefix(cons, "yt") {
            b.WriteString(substr)
            b.WriteString("ay")
            continue
        }
        ypos := strings.Index(cons, "y")
        if ypos > 0 {
            b.WriteString(substr[ypos:])
            b.WriteString(substr[:ypos])
            b.WriteString("ay")
            continue
        }
        if strings.HasSuffix(cons, "q") && strings.HasPrefix(vowel, "u") {
            b.WriteString(vowel[1:])
            b.WriteString(cons)
            b.WriteString(vowel[:1])
            b.WriteString("ay")
            continue
        }
        b.WriteString(vowel)
        b.WriteString(cons)
        b.WriteString("ay")
    }
    return b.String()
}

func sepUtil(substr string, sepAt func(rune) bool) (string, string) {
    sep := strings.IndexFunc(substr, sepAt)
    if sep == -1 {
        sep = len(substr)
    }
    return substr[:sep], substr[sep:]
}

func spaceWord(substr string) (string, string) {
    return sepUtil(substr, isWord)
}

func isWord(r rune) bool {
    return !unicode.IsSpace(r)
}

func wordSpace(substr string) (string, string) {
    return sepUtil(substr, unicode.IsSpace)
}

func consVowel(substr string) (string, string) {
    return sepUtil(substr, isVowel)
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
    default:
        return false
	}
}