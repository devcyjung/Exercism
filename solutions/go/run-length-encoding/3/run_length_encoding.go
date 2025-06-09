package encode

import (
    "strconv"
    "strings"
)

func RunLengthEncode(input string) string {
    var b strings.Builder
	var run rune
    var length int
    for _, r := range input {
        switch {
        case r == run:
            length++
        default:
            writeEncoding(&b, run, length)
        	run = r
            length = 1
        }
    }
    writeEncoding(&b, run, length)
    return b.String()
}

func writeEncoding(b *strings.Builder, run rune, length int) {
    if length > 1 {
        b.WriteString(strconv.Itoa(length))
    }
    if length > 0 {
        b.WriteRune(run)
    }
}

func RunLengthDecode(input string) string {
	var b strings.Builder
    var length int
    for _, r := range input {
        switch {
        case '0' <= r && r <= '9':
            length *= 10
            length += int(r - '0')
        case length == 0:
            b.WriteRune(r)
        default:
            for length > 0 {
                b.WriteRune(r)
                length--
            }
        }
    }
    return b.String()
}