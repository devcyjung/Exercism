package raindrops

import (
    "strconv"
    "strings"
)

var soundMap = [3]struct{
    divisor	int
    sound	string
} {
    {3, "Pling"},
    {5, "Plang"},
    {7, "Plong"},
}

func Convert(number int) string {
	var builder strings.Builder
    for _, mapping := range soundMap {
        if number % mapping.divisor == 0 {
            builder.WriteString(mapping.sound)
        }
    }
    if builder.Len() == 0 {
        builder.WriteString(strconv.Itoa(number))
    }
    return builder.String()
}