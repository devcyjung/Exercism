package cipher

import "strings"

type shift rune
type vigenere []rune

func NewCaesar() Cipher {
	return shift(3)
}

func NewShift(distance int) Cipher {
    if distance == 0 || distance > 25 || distance < -25 {
        return nil
    }
	return shift(distance)
}

func (c shift) Encode(input string) string {
	var r []rune
    var n rune
    for _, x := range input {
        if 'a' <= x && x <= 'z' {
            n = x + rune(c)
        } else if 'A' <= x && x <= 'Z' {
            n = x + rune(c) - 'A' + 'a'
        } else {
            continue
        }
        if n > 'z' {
            n += 'a' - 'z' - 1
        } else if n < 'a' {
            n += 'z' + 1 - 'a'
        }
        r = append(r, n)
    }
    return string(r)
}

func (c shift) Decode(input string) string {
    t := -c
	return t.Encode(input)
}

func NewVigenere(key string) Cipher {
    var r []rune
    var foundUsefulKey bool
    for _, x := range key {
        if x < 'a' || x > 'z' {
            return nil
        }
        if x != 'a' {
        	foundUsefulKey = true
        }
        r = append(r, x - 'a')
    }
    if !foundUsefulKey {
        return nil
    }
    return vigenere(r)
}

func (v vigenere) Encode(input string) string {
    key := []rune(v)
    keysize := len(key)
    var builder strings.Builder
    var i int
	for _, s := range input {
        if !('a' <= s && s <= 'z') && !('A' <= s && s <= 'Z') {
            continue
        }
        sh := shift(key[i % keysize])
        en := sh.Encode(string(s))
        builder.WriteString(en)
        i++
    }
    return builder.String()
}

func (v vigenere) Decode(input string) string {
    key := []rune(v)
    keysize := len(key)
    var builder strings.Builder
    var i int
	for _, s := range input {
        if !('a' <= s && s <= 'z') && !('A' <= s && s <= 'Z') {
            continue
        }
        sh := shift(key[i % keysize])
        en := sh.Decode(string(s))
        builder.WriteString(en)
        i++
    }
    return builder.String()
}
