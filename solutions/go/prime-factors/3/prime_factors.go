package prime

import (
    "math"
    "slices"
)

func Factors(n int64) []int64 {
    factors := make([]int64, 0, int(math.Log2(float64(n))))
    for n % 2 == 0 {
        n /= 2
        factors = append(factors, 2)
    }
    var divisor int64 = 3
    for n > 1 {
        for n % divisor == 0 {
            n /= divisor
            factors = append(factors, divisor)
        }
        divisor += 2
    }
    slices.Clip(factors)
    return factors
}