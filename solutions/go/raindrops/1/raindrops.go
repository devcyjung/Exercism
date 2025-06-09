package raindrops

import (
    "strconv"
    "strings"
)

type Sound = string

const (
    DivByThree	Sound	= "Pling"
    DivByFive	Sound	= "Plang"
    DivBySeven	Sound	= "Plong"
)

func Convert(number int) string {
	var builder strings.Builder
    if number % 3 == 0 {
        builder.WriteString(DivByThree)
    }
    if number % 5 == 0 {
        builder.WriteString(DivByFive)
    }
    if number % 7 == 0 {
        builder.WriteString(DivBySeven)
    }
    if builder.Len() == 0 {
        builder.WriteString(strconv.Itoa(number))
    }
    return builder.String()
}
