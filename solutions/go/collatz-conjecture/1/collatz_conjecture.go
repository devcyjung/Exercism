package collatzconjecture

import "errors"

func CollatzConjecture(n int) (steps int, e error) {
    if n <= 0 {
        e = errors.ErrUnsupported
        return
    }
    for n != 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = 3*n + 1
        }
        steps++
    }
    return
}
