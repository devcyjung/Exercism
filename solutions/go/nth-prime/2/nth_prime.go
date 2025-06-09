package prime

import "errors"

func Nth(n int) (r int, e error) {
	if n <= 0 {
        e = errors.ErrUnsupported
        return
    }
    primes := []int{}
    cur := 2
	var isPrime bool
    for len(primes) < n {
        isPrime = true
        for _, p := range primes {
            if cur % p == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            primes = append(primes, cur)
        }
        cur++
    }
    r = primes[n-1]
    return
}
