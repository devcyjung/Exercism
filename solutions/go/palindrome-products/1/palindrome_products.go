package palindrome

import (
    "errors"
    "strconv"
)

var (
    ErrInvalidInput = errors.New("fmin > fmax")
    ErrNoResult = errors.New("no palindromes")
)

type Product struct {
    product	int
    Factorizations	[][2]int
}

func Products(fmin, fmax int) (pmin Product, pmax Product, err error) {
	if fmin > fmax {
        err = ErrInvalidInput
        return
    }
    found := false
    for i := fmin; i <= fmax; i++ {
        for j := i; j <= fmax; j++ {
            if isPalindrome(i * j) {
                switch {
                case !found || i * j < pmin.product:
                    pmin.product = i * j
                    pmin.Factorizations = [][2]int{{i, j}}
                case i * j == pmin.product:
                    pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
                case i * j == pmax.product:
                    pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
                case i * j > pmax.product:
                    pmax.product = i * j
                    pmax.Factorizations = [][2]int{{i, j}}
                }
                found = true
            }
        }
    }
    if !found {
        err = ErrNoResult
    }
    return
}

func isPalindrome(n int) bool {
    str := strconv.Itoa(n)
    for i := 0; i <= len(str) / 2; i++ {
        if str[i] != str[len(str) - 1 - i] {
            return false
        }
    }
    return true
}