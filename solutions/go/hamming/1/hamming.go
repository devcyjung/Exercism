package hamming

import "errors"

func Distance(a, b string) (d int, e error) {
    runeA := []rune(a)
    runeB := []rune(b)
	if len(runeA) != len(runeB) {
        e = errors.ErrUnsupported
        return
    }
    for i, j := range runeA {
        if runeB[i] != j {
            d++
        }
    }
    return
}
