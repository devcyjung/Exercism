package grains

import "errors"

func Square(number int) (r uint64, e error) {
	if number <= 0 || 64 < number {
        e = errors.New("Number out of range")
        return
    }
    r = 1 << (number-1)
    return
}

func Total() uint64 {
	return ^uint64(0)
}