package summultiples

func SumMultiples(limit int, divisors ...int) int {
	multiples := make(map[int]bool)
    for _, divisor := range divisors {
        if divisor == 0 {
            continue
        }
        for i := 1; i * divisor < limit; i++ {
            multiples[i * divisor] = true
        }
    }
    acc := 0
    for divisor := range multiples {
        acc += divisor
    }
    return acc
}