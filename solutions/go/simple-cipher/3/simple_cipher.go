package cipher

import "strings"

type vigenere []int

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
    if distance < 0 {
        return NewVigenere(string('z' + 1 + rune(distance)))
    }
	return NewVigenere(string('a' + rune(distance)))
}

func NewVigenere(key string) Cipher {
	v := make(vigenere, 0)
    var hasNonzero bool
    for _, r := range key {
        if r < 'a' || 'z' < r {
            return nil
        }
        if r != 'a' {
            hasNonzero = true
        }
        v = append(v, int(r - 'a'))
    }
    if !hasNonzero {
        return nil
    }
    return v
}

func (v vigenere) Encode(input string) string {
	vlen := len(v)
    var b strings.Builder
    for i, r := range sanitize(input) {
        b.WriteRune(convert(r, v[i%vlen]))
    }
    return b.String()
}

func (v vigenere) Decode(input string) string {
	vlen := len(v)
    var b strings.Builder
    for i, r := range input {
        b.WriteRune(convert(r, -v[i%vlen]))
    }
    return b.String()
}

func convert(r rune, offset int) rune {
    return (r + rune(offset) - 'a' + 26) % 26 + 'a'
}

func sanitize(input string) string {
    return strings.ToLower(strings.Join(strings.FieldsFunc(input, outOfBounds), ""))
}

func outOfBounds(r rune) bool {
    return !('a' <= r && r <= 'z') && !('A' <= r && r <= 'Z')
}