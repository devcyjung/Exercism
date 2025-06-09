package cryptosquare

import (
    "math"
    "strings"
    "unicode"
)

func Encode(pt string) string {
	plaintext := strings.ToLower(strings.Join(strings.FieldsFunc(pt, onlyAlnum), ""))
    size := len(plaintext)
    root := int(math.Sqrt(float64(size)))
    var r, c int
    if root * root == size {
        r, c = root, root
    } else if root * (root + 1) >= size {
        r, c = root, root + 1
    } else {
        r, c = root + 1, root + 1
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

func onlyAlnum(r rune) bool {
    return !unicode.IsDigit(r) && !unicode.IsLetter(r)
}