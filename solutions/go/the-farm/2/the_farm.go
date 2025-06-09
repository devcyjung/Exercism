package thefarm

import (
    "errors"
    "fmt"
)

func DivideFood(f FodderCalculator, n int) (float64, error) {
    fa, errfa := f.FodderAmount(n)
    if errfa != nil {
        return fa, errfa
    }
    ff, errff := f.FatteningFactor()
    if errff != nil {
        return ff, errff
    }
    return float64(ff * fa) / float64(n), nil
}

func ValidateInputAndDivideFood(f FodderCalculator, n int) (float64, error) {
    if n > 0 {
        return DivideFood(f, n)
    }
    return 0.0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
    n int
    message string
}

func (i *InvalidCowsError) Error() string {
    return fmt.Sprintf("%v cows are invalid: %v", i.n, i.message) 
}

func ValidateNumberOfCows(n int) error {
    e := new(InvalidCowsError)
    e.n = n
    if n < 0 {
        e.message = "there are no negative cows"
        return e
    }
    if n == 0 {
        e.message = "no cows don't need food"
        return e
    }
    return nil
}