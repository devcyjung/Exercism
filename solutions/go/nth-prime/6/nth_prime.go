package prime

import (
    "errors"
    "slices"
)

var (
    ErrNonPositiveNumber = errors.New("Input must be a positive number")
    primes = make([]int, 0, 128)
)

func Nth(n int) (int, error) {
	if n <= 0 {
        return 0, ErrNonPositiveNumber
    }
    if n == 1 {
        return 2, nil
    }
    slices.Grow(primes, n - 1)
    if len(primes) == 0 {
        primes = append(primes, 3)
    }
nextInt:
    for i := primes[len(primes) - 1] + 2; len(primes) < n - 1; i += 2 {
        for _, prime := range primes {
            if i % prime == 0 {
                continue nextInt
            }
        }
        primes = append(primes, i)
    }
    return primes[n - 2], nil
}