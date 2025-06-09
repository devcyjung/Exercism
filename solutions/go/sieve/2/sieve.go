package sieve

import "slices"

var primes = make([]int, 0, 128)

func Sieve(limit int) []int {
    if len(primes) == 0 {
        primes = append(primes, 2, 3)
    }
    if limit <= primes[len(primes) - 1] {
        index, ok := slices.BinarySearch(primes, limit)
        if ok {
            return slices.Clone(primes[:index + 1])
        } else {
            return slices.Clone(primes[:index])
        }
    }
nextInt:
    for i := primes[len(primes) - 1] + 2; i <= limit; i += 2 {
        for _, prime := range primes[1:] {
            if i % prime == 0 {
                continue nextInt
            }
        }
        primes = append(primes, i)
    }
    return slices.Clone(primes)
}