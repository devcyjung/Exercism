package lsproduct

import (
    "errors"
    "slices"
    "strings"
)

var (
    ErrNegativeSpan = errors.New("span must not be negative")
    ErrNonDigitInput = errors.New("digits input must only contain digits")
    ErrSpanLongerThanInput = errors.New("span must be smaller than string length")
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    switch {
    case span < 0:
        return 0, ErrNegativeSpan
    case span > len(digits):
        return 0, ErrSpanLongerThanInput
    case strings.ContainsFunc(digits, nonDigit):
        return 0, ErrNonDigitInput
    }
    products := make([]int64, len(digits) - span + 1)
    var product int64
    for i := 0; i < len(products); i++ {
        product = 1
        for j := i; j < i + span; j++ {
            product *= int64(digits[j] - '0')
        }
        products[i] = product
    }
    return slices.Max(products), nil
}

func nonDigit(r rune) bool {
    return !('0' <= r && r <= '9')
}