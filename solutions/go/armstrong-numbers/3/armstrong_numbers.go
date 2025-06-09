package armstrong

import (
    "math"
    "strconv"
    "strings"
)

func IsNumber(n int) bool {
    digits := strconv.Itoa(n)
    size := float64(len(digits))
    armstrongSum := 0
    _ = strings.Map(func (r rune) rune {
        armstrongSum += int(math.Pow(float64(r - '0'), size))
        return -1
    }, digits)
    return armstrongSum == n
}