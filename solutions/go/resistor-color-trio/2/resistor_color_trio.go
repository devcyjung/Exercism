package resistorcolortrio

import (
    "fmt"
    "math"
    "slices"
)

var (
    colorList = []string{
        "black", "brown", "red", "orange", "yellow",
        "green", "blue", "violet", "grey", "white",
    }
    quantifiers = []string{
    	"ohms", "kiloohms", "megaohms", "gigaohms",
	}
)

func Label(colors []string) string {
	n1 := slices.Index(colorList, colors[0])
    n2 := slices.Index(colorList, colors[1])
    n3 := slices.Index(colorList, colors[2])
    if n1 == -1 || n2 == -1 || n3 == -1 {
        panic("Invalid color")
    }
    num := float64(n1 * 10 + n2) * math.Pow10(n3)
    qIdx := 0
	for num > 1000 {
        num /= 1000
        qIdx++
    }
    if qIdx > len(quantifiers) - 1 {
        panic("Value out of bounds")
    }
    return fmt.Sprintf("%v %v", formatFloat(num), quantifiers[qIdx])
}

func formatFloat(f float64) string {
	if f == float64(int64(f)) {
		return fmt.Sprintf("%d", int64(f))
	}
	return fmt.Sprintf("%g", f)
}