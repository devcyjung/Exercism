package allyourbase

import "errors"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (r []int, e error) {
	if inputBase < 2 {
        e = errors.New("input base must be >= 2")
        return
    }
    if outputBase < 2 {
        e = errors.New("output base must be >= 2")
        return
    }
    num := 0
    pow := 0
    var v int
    for i := len(inputDigits)-1; i>=0; i-- {
        v = inputDigits[i]
        if v < 0 || inputBase <= v {
            r, e = []int{}, errors.New("all digits must satisfy 0 <= d < input base")
            return
        }
        num += power(inputBase, pow) * v
        pow++
    }
    var digit int
    for num > 0 {
        num, digit = num / outputBase, num % outputBase
        r = append([]int{digit}, r...)
    }
    if len(r) == 0 {
        r = append(r, 0)
    }
    return
}

func power(base, pow int) int {
    mul := 1
    for i := 0; i < pow; i++ {
        mul *= base
    }
    return mul
}