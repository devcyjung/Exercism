package allyourbase

import (
    "errors"
    "math"
    "slices"
)

var (
    ErrInvalidInputBase = errors.New("input base must be >= 2")
    ErrInvalidOutputBase = errors.New("output base must be >= 2")
    ErrInvalidDigit = errors.New("all digits must satisfy 0 <= d < input base")
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	switch {
    case inputBase < 2:
        return nil, ErrInvalidInputBase
    case outputBase < 2:
        return nil, ErrInvalidOutputBase
    }
    num := 0
    for _, digit := range inputDigits {
        if digit < 0 || inputBase <= digit {
            return nil, ErrInvalidDigit
        }
        num *= inputBase
        num += digit
    }
    if num == 0 {
        return []int{0}, nil
    }
    output := make([]int, 0, int(math.Log(float64(num)) / math.Log(float64(outputBase))) + 1)
    for num > 0 {
        output = append(output, num % outputBase)
        num /= outputBase
    }
    slices.Clip(output)
    slices.Reverse(output)
    return output, nil
}