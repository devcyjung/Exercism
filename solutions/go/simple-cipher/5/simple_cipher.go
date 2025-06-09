package cipher

import "strings"

type shift struct {
    vigenere
}

type vigenere struct {
    keys	[]int
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
    var c Cipher
    if distance < 0 {
        c = NewVigenere(string('z' + 1 + rune(distance)))
    } else {
        c = NewVigenere(string('a' + rune(distance)))
    }
    if c == nil {
        return nil
    }
    return shift{ vigenere: c.(vigenere) }
}

func NewVigenere(key string) Cipher {
	keys := make([]int, len(key))
    var hasNonzero bool
    for i, r := range key {
        if r < 'a' || 'z' < r {
            return nil
        }
        if r != 'a' {
            hasNonzero = true
        }
        keys[i] = int(r - 'a')
    }
    if !hasNonzero {
        return nil
    }
    return vigenere{keys}
}

func (v vigenere) Encode(input string) string {
	vlen := len(v.keys)
    var b strings.Builder
    for i, r := range sanitize(input) {
        b.WriteRune(convert(r, v.keys[i%vlen]))
    }
    return b.String()
}

func (v vigenere) Decode(input string) string {
	vlen := len(v.keys)
    var b strings.Builder
    for i, r := range input {
        b.WriteRune(convert(r, -v.keys[i%vlen]))
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