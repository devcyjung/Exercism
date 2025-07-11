package logs

// Application identifies the application emitting the given log.
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

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    r := []rune(log)
	for i, v := range r {
        if v == oldRune {
            r[i] = newRune
        }
    }
    return string(r)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
