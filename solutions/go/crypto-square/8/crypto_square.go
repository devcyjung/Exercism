package cryptosquare

import (
    "math"
    "strings"
    "unicode"
)

func Encode(pt string) string {
	plaintext := fString(pt).filter(isalnum).lowercase()
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

type fString string

func (f fString) filter(pred func(rune) bool) fString {
    s := string(f)
    result := strings.Join(strings.FieldsFunc(s, func(r rune) bool {
        return !pred(r)
    }), "")
    return fString(result)
}

func (f fString) lowercase() fString {
    s := string(f)
    result := strings.ToLower(s)
    return fString(result)
}

func isalnum(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}