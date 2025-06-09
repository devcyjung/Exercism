package encode

import (
    "strconv"
    "strings"
)

func flush(b *strings.Builder, count int, ch rune) {
    if count > 1 {
        b.WriteString(strconv.Itoa(count))
    }
    if count > 0 {
        b.WriteRune(ch)
    }
}

func RunLengthEncode(input string) string {
	var b strings.Builder
    b.Grow(len(input))
    var ch rune
    var count int

    for _, r := range input {
        if r != ch {
            flush(&b, count, ch)
            ch = r
            count = 1
        } else {
            count++
        }
    }
    flush(&b, count, ch)
    
    return b.String()
}

func RunLengthDecode(input string) string {
	var b, numBuffer strings.Builder
    var count int
    var parseErr error
    for _, r := range input {
        if '0' <= r && r <= '9' {
            numBuffer.WriteRune(r)
        } else {
            count, parseErr = strconv.Atoi(numBuffer.String())
            if (parseErr != nil) {
                count = 1
            }
            for i := 0; i < count; i++ {
                b.WriteRune(r)
            }
            numBuffer.Reset()
        }
    }
    return b.String()
}
