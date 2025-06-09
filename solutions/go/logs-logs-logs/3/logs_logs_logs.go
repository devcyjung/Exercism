package logs

import "strings"

var b strings.Builder

func Application(log string) string {
	for _, v := range log {
        switch v {
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
    b.Reset()
	for _, v := range log {
        if v == oldRune {
        	b.WriteRune(newRune)
        } else {
            b.WriteRune(v)
        }
    }
    return b.String()
}

func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}