package armstrong

import (
    "strconv"
    "strings"
)

func IsNumber(n int) bool {
    digits := strconv.Itoa(n)
    size := len(digits)
    armstrongSum := 0
    strings.Map(func (r rune) rune {
        armstrongSum += ipow(int(r - '0'), size)
        return -1
    }, digits)
    return armstrongSum == n
}

func ipow(a, b int) int {
    acc := 1
    for i := 0; i < b; i++ {
        acc *= a
    }
    return acc
}