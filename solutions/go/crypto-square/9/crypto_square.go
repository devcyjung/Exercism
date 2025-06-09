package cryptosquare

import (
    "math"
    "strings"
    "unicode"
)

func Encode(pt string) string {
	plaintext := text(pt).filter(isalnum).lowercase()
    size := len(plaintext)
    r := int(math.Sqrt(float64(size)))
    c := r
	var changeR bool
    for r * c < size {
        if changeR {
            r++
        } else {
            c++
        }
        changeR = !changeR
    }
    var b strings.Builder
    var textIdx int
    for ci := 0; ci < c; ci++ {
        for ri := 0; ri < r; ri++ {
            textIdx = ri * c + ci
            if textIdx < size {
                b.WriteByte(plaintext[textIdx])
            } else {
                b.WriteByte(' ')
            }
        }
        if ci != c - 1 {
            b.WriteByte(' ')
        }
    }
    return b.String()
}

type text string

func (t text) filter(pred func(rune) bool) text {
    s := string(t)
    result := strings.Join(strings.FieldsFunc(s, func(r rune) bool {
        return !pred(r)
    }), "")
    return text(result)
}

func (t text) lowercase() text {
    s := string(t)
    result := strings.ToLower(s)
    return text(result)
}

func isalnum(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}