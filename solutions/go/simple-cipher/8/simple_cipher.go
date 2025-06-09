package cipher

import (
    "slices"
    "strings"
)

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	return newCipher([]rune{rune(distance)})
}

func NewVigenere(key string) Cipher {
	return newCipher([]rune(strings.Map(encodingGenerater, key)))
}

func encodingGenerater(r rune) rune {
    if 'a' <= r && r <= 'z' {
        return r - 'a'
    }
    return r
}

type cipher []rune

func newCipher(encoding []rune) Cipher {
    if slices.ContainsFunc(encoding, invalidEncoding) {
        return nil
    }
    if !slices.ContainsFunc(encoding, meaningfulEncoding) {
        return nil
    }
    return cipher(encoding)
}

func invalidEncoding(r rune) bool {
    return r < rune(-25) || rune(25) < r
}

func meaningfulEncoding(r rune) bool {
    return r != rune(0)
}

func (c cipher) Encode(input string) string {
    index := 0
    size := len(c)
    var encoding rune
    return strings.Map(func(r rune) rune {
        encoding = c[index]
        switch {
        case 'a' <= r && r <= 'z':
            index = (index + 1) % size
            return ((r - 'a' + encoding) + 26) % 26 + 'a'
        case 'A' <= r && r <= 'Z':
            index = (index + 1) % size
            return ((r - 'A' + encoding) + 26) % 26 + 'a'
        default:
            return -1
        }
    }, input)
}

func (c cipher) Decode(input string) string {
    index := 0
    size := len(c)
    var encoding rune
    return strings.Map(func(r rune) rune {
        encoding = c[index]
        switch {
        case 'a' <= r && r <= 'z':
            index = (index + 1) % size
            return ((r - 'a' - encoding) + 26) % 26 + 'a'
        case 'A' <= r && r <= 'Z':
            index = (index + 1) % size
            return ((r - 'A' - encoding) + 26) % 26 + 'a'
        default:
            return -1
        }
    }, input)
}