package wordcount

import (
    "regexp"
    "strings"
)

type Frequency map[string]int

const apos rune = '\''

var quoteRemover = regexp.MustCompile(`(\s'|'\s|^'|'$)`)

func WordCount(phrase string) Frequency {
    sanitizedString := quoteRemover.ReplaceAllString(strings.Map(lowerAndReplacePunct, phrase), " ")
    frequency := make(Frequency)
    for _, field := range strings.Fields(sanitizedString) {
        frequency[field]++
    }
    return frequency
}

func lowerAndReplacePunct(r rune) rune {
    switch {
    case 'A' <= r && r <= 'Z':
        return r - 'A' + 'a'
    case 'a' <= r && r <= 'z':
        return r
    case '0' <= r && r <= '9':
        return r
    case r == apos || r == '-':
        return r
    default:
        return ' '
    }
}