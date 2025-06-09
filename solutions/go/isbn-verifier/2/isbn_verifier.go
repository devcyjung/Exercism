package isbn

import (
    "index/suffixarray"
    "regexp"
)

var (
    invalidPattern = regexp.MustCompile(`[^0-9-Xx]`)
    checkPattern = regexp.MustCompile(`[Xx]`)
    digitPattern = regexp.MustCompile(`[0-9Xx]`)
)

func IsValidISBN(isbn string) (b bool) {
    bytes := []byte(isbn)
    sfx := suffixarray.New(bytes)
    if len(sfx.FindAllIndex(invalidPattern, -1)) > 0 {
        return false
    }
    checks := sfx.FindAllIndex(checkPattern, -1)
    if len(checks) > 1 || len(checks) == 1 && checks[0][1] != len(bytes) {
        return false
    }
    digitIndices := sfx.FindAllIndex(digitPattern, -1)
    if len(digitIndices) != 10 {
        return false
    }
    acc := 0
    for i, indicesPair := range sfx.FindAllIndex(digitPattern, -1) {
        acc += byteToInt(bytes[indicesPair[0]]) * (10 - i)
    }
    acc %= 11
    return acc == 0
}

func byteToInt(b byte) int {
    switch b {
        case 'x':
        	return 10
        case 'X':
        	return 10
        default:
        	return int(b - '0')
    }
}