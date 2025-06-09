package logs

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
    r := []rune(log)
	for i, v := range r {
        if v == oldRune {
            r[i] = newRune
        }
    }
    return string(r)
}

func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
