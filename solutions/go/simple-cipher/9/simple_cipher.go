package cipher

import "strings"

type shift struct {
    vigenere
}

type vigenere struct {
    keys	[]rune
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
    if strings.ContainsFunc(key, invalidKey) || !strings.ContainsFunc(key, meaningfulKey) {
        return nil
    }
    keys := make([]rune, len(key))
    for i, r := range key {
        keys[i] = r - 'a'
    }
    return vigenere{keys}
}

func invalidKey(r rune) bool {
    return r < 'a' || 'z' < r
}

func meaningfulKey(r rune) bool {
    return r != 'a'
}

func (v vigenere) Encode(input string) string {
    vlen := len(v.keys)
    var b strings.Builder
    for i, r := range sanitize(input) {
        b.WriteRune(shiftAndRotate(r, v.keys[i%vlen]))
    }
    return b.String()
}

func (v vigenere) Decode(input string) string {
    vlen := len(v.keys)
    var b strings.Builder
    for i, r := range input {
        b.WriteRune(shiftAndRotate(r, -v.keys[i%vlen]))
    }
    return b.String()
}

func shiftAndRotate(r, offset rune) rune {
    return (r + offset - 'a' + 26) % 26 + 'a'
}

func sanitize(input string) string {
    return strings.Map(sanitizeMapping, input)
}

func sanitizeMapping(r rune) rune {
    if ('a' <= r && r <= 'z') {
        return r
    }
    if ('A' <= r && r <= 'Z') {
        return r - 'A' + 'a'
    }
    return -1
}