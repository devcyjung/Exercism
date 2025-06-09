package prime

import "errors"

var ErrNonPositiveNumber = errors.New("Input must be a positive number")

func Nth(n int) (int, error) {
	if n <= 0 {
        return 0, ErrNonPositiveNumber
    }
    if n == 1 {
        return 2, nil
    }
    primes := make([]int, 0, n - 1)
nextInt:
    for i := 3; len(primes) < n; i += 2 {
        for _, prime := range primes {
            if i % prime == 0 {
            	continue nextInt
            }
        }
        primes = append(primes, i)
    }
    return primes[n - 2], nil
}
