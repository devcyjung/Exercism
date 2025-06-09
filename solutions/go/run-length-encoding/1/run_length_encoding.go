package encode

import (
    "strconv"
    "strings"
)

func RunLengthEncode(input string) string {
	var b strings.Builder
    b.Grow(len(input))
    var ch rune
    var count int
    
    flush := func() {
        if count > 1 {
            b.WriteString(strconv.Itoa(count))
        }
        if count > 0 {
            b.WriteRune(ch)
        }
    }
    
    for _, r := range input {
        if r != ch {
            flush()
            ch = r
            count = 1
        } else {
            count++
        }
    }
    flush()
    
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
