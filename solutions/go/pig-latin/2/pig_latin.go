package piglatin

import (
    "strings"
    "unicode"
)

type cache struct {
    i	int
}

func Sentence(sentence string) string {
    c := &cache{}
    var b strings.Builder
    words := strings.Fields(sentence)
    spaces := strings.FieldsFunc(sentence, nonspace)
    var isWord bool
    if strings.IndexFunc(sentence, space) != 0 {
        isWord = true
    }
    size := len(words) + len(spaces)
    for i := 0; i < size; i++ {
        if isWord {
            b.WriteString(translateWord(words[0], c))
            words = words[1:]
        } else {
            b.WriteString(spaces[0])
            spaces = spaces[1:]
        }
        isWord = !isWord
    }
    return b.String()
}

func space(r rune) bool {
    return unicode.IsSpace(r)
}

func nonspace(r rune) bool {
    return !unicode.IsSpace(r)
}

func translateWord(word string, c *cache) string {
	switch {
    case rule1(word):
        return word + "ay"
    case rule3(word, c):
        return word[c.i+2:] + word[:c.i+2] + "ay"
    case rule4(word, c):
        return word[c.i:] + word[:c.i] + "ay"
    case rule2(word, c):
        return word[c.i:] + word[:c.i] + "ay"
    default:
        return word
    }
}

func rule1(s string) bool {
    return firstVowel(s) == 0 || prefix(s, "xr", "yt")
}

func rule2(s string, c *cache) bool {
    c.i = leadingConsonants(s)
    return c.i > 0
}

func rule3(s string, c *cache) bool {
    c.i = strings.Index(s, "qu")
    return c.i != -1 && c.i + 1 == firstVowel(s)
}

func rule4(s string, c *cache) bool {
    c.i = strings.Index(s, "y")
    return c.i != -1 && c.i > 0 && c.i + 1 <= leadingConsonants(s)
}

func prefix(s string, prefixes ...string) bool {
    for _, pre := range prefixes {
        if strings.HasPrefix(s, pre) {
            return true
        }
    }
    return false
}

func firstVowel(s string) int {
    return strings.IndexAny(s, "aeiou")
}

func leadingConsonants(s string) int {
    f := firstVowel(s)
    switch f {
    case -1:
        return len(s)
    default:
        return f
    }
}