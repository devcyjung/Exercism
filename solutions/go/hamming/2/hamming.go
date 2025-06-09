package hamming

import "errors"

func Distance(a, b string) (int, error) {
    runeA, runeB := []rune(a), []rune(b)
	if len(runeA) != len(runeB) {
        return -1, errors.New("Two strands must have equal length")
    }
    dist := 0
    for i, v := range runeA {
        if runeB[i] != v {
            dist++
        }
    }
    return dist, nil
}