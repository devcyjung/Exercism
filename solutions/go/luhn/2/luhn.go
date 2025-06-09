package luhn

import (
    "strings"
    "slices"
    "unicode"
)

func Valid(id string) bool {
    if strings.ContainsFunc(id, validator) {
        return false
    }
	nums := []rune(strings.Map(sanitizer, id))
    size := len(nums)
	if size <= 1 {
        return false
    }
    var luhn rune
    slices.Reverse(nums)
    for i, v := range nums {
        switch {
        case i % 2 == 0:
            luhn += v
        case 2 * v > 9:
            luhn += 2 * v - 9
        default:
            luhn += 2 * v
        }
    }
    return luhn % 10 == 0
}

func validator(r rune) bool {
    return !unicode.IsSpace(r) && !unicode.IsDigit(r)
}

func sanitizer(r rune) rune {
    if unicode.IsSpace(r) {
        return -1
    }
    return r - '0'
}