package logs

import "strings"

func Application(log string) string {
	for _, r := range log {
        switch r {
        case '\u2757':
            return "recommendation"
        case '\U0001F50D':
            return "search"
        case '\u2600':
            return "weather"
        }
    }
    return "default"
}

func Replace(log string, oldRune, newRune rune) string {
    var b strings.Builder
	for _, r := range log {
        if r == oldRune {
        	b.WriteRune(newRune)
        } else {
            b.WriteRune(r)
        }
    }
    return b.String()
}

func WithinLimit(log string, limit int) bool {
    if len(log) <= limit {
        return true
    }
	runeCount := 0
    for range log {
        runeCount++
        if runeCount > limit {
            return false
        }
    }
    return true
}