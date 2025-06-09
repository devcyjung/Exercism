package resistorcolortrio

import (
    "slices"
    "strings"
)

var Colors = []string{
    "black",
    "brown",
    "red",
    "orange",
    "yellow",
    "green",
    "blue",
    "violet",
    "grey",
    "white",
}

var Quantifiers = []string{
    "ohms",
    "kiloohms",
    "megaohms",
    "gigaohms",
}

func Label(colors []string) string {
	n1 := slices.Index(Colors, colors[0])
    n2 := slices.Index(Colors, colors[1])
    n3 := slices.Index(Colors, colors[2])
    var additionalDigits int
    if n1 != 0 {
        additionalDigits += 2
    } else if n2 != 0 {
        additionalDigits++
    }
    var quantifier int
    if n3 + additionalDigits > 3 {
        quantifier++
    }
    if n3 + additionalDigits > 6 {
        quantifier++
    }
    if n3 + additionalDigits > 9 {
        quantifier++
    }
    if additionalDigits == 0 {
        quantifier = 0
    }
    nDigits := n3 + additionalDigits - 3 * quantifier
    var builder strings.Builder
    if n1 != 0 {
        builder.WriteRune('0' + rune(n1))
    }
    if n1 == 0 || n1 != 0 && nDigits >= 2 {
        builder.WriteRune('0' + rune(n2))
    }
    if nDigits == 3 {
        builder.WriteRune('0')
    }
    builder.WriteRune(' ')
    builder.WriteString(Quantifiers[quantifier])
    return builder.String()
}
