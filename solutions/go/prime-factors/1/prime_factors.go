package prime

import "slices"

func Factors(n int64) (s []int64) {
    slices.Grow(s, 100)
    divisor := int64(2)
    for n > 1 {
        if n % divisor == 0 {
            n /= divisor
            s = append(s, divisor)
        } else {
            divisor++
        }
    }
    slices.Clip(s)
	return
}