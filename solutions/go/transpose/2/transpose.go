package transpose

import (
    "cmp"
    "slices"
    "strings"
    "unicode/utf8"
)

const deadRune rune = 0xD9FF

var deadRuneReplacer = strings.NewReplacer(string(deadRune), " ")

func Transpose(input []string) []string {
    nrows := len(input)
    if nrows == 0 {
        return []string{}
    }
	ncols := utf8.RuneCountInString(slices.MaxFunc(input, func(a, b string) int {
        return cmp.Compare(utf8.RuneCountInString(a), utf8.RuneCountInString(b))
    }))
    array := make([]rune, nrows * ncols)
    for i := range array {
        array[i] = ' '
    }
    for i, str := range input {
        for j, r := range str {
            if r == ' ' {
                r = deadRune
            }
            array[j * nrows + i] = r
        }
    }
    result := make([]string, ncols)
    for j := range result {
        trimmed := strings.TrimRight(string(array[j * nrows:(j+1) * nrows]), " ")
        result[j] = deadRuneReplacer.Replace(trimmed)
    }
    return result
}