package reverse

import "slices"

func Reverse(input string) string {
    runes := []rune(input)
    slices.Reverse(runes)
    return string(runes)
}
