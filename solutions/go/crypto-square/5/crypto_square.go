package cryptosquare

import (
    "math"
    "strings"
    "unicode"
)

func alnumOnly(r rune) bool {
    return !unicode.IsLetter(r) && !unicode.IsDigit(r)
}

func Encode(pt string) string {
	plaintext := strings.ToLower(strings.Join(strings.FieldsFunc(pt, alnumOnly), ""))
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