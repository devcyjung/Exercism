package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return -1, errors.New("Must provide positive number")
    }
    steps := 0
    for n != 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = 3 * n + 1
        }
        steps++
    }
    return steps, nil
}