package romannumerals

import (
    "errors"
    "strings"
)

var Literals = []rune{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
var LiteralVals = []int{1000, 500, 100, 50, 10, 5, 1}

func ToRomanNumeral(input int) (r string, e error) {
	if input <= 0 || len(Literals) != len(LiteralVals) {
        e = errors.ErrUnsupported
        return
    }
 	var builder strings.Builder
    var digit int
    for index := 0; index < len(Literals); index += 2 {
        digit, input = input / LiteralVals[index], input % LiteralVals[index]
        if digit >= 4 {
            if index == 0 {
                e = errors.ErrUnsupported
                return
            }
            if digit == 9 {
                builder.WriteRune(Literals[index])
                builder.WriteRune(Literals[index-2])
            } else if digit == 4 {
                builder.WriteRune(Literals[index])
                builder.WriteRune(Literals[index-1])
            } else {
                builder.WriteRune(Literals[index-1])
                for j := 0; j < digit - 5; j++ {
                    builder.WriteRune(Literals[index])
                }
            }
        } else {
            for j := 0; j < digit; j++ {
                builder.WriteRune(Literals[index])
            }
        }
    }
    r = builder.String()
    return
}
