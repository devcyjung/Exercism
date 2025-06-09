package romannumerals

import (
    "errors"
    "slices"
    "strings"
)

var numerals = [7]string{
    "I", "V", "X", "L", "C", "D", "M",
}

func ToRomanNumeral(input int) (string, error) {
    if input <= 0 || 4000 <= input {
        return "", errors.New("Input out of range")
    }
    digits := make([]string, 0, 4)
    var rem int
    for i := 0; input > 0; i += 2 {
        rem = input % 10
        switch {
        case rem < 4:
            digits = append(digits, strings.Repeat(numerals[i], rem))
        case rem == 4:
            digits = append(digits, numerals[i] + numerals[i + 1])
        case rem < 9:
            digits = append(digits, numerals[i + 1] + strings.Repeat(numerals[i], rem - 5))
        case rem == 9:
            digits = append(digits, numerals[i] + numerals[i + 2])
        }
        input /= 10
    }
    slices.Reverse(digits)
    return strings.Join(digits, ""), nil
}