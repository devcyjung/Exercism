package diamond

import (
    "errors"
    "strings"
)

func Gen(char byte) (string, error) {
    order := char - 'A'
    if !(0 <= order && order <= 25) {
        return "", errors.New("input is not in valid range")
    }
    side := 2 * order + 1
    buf := make([]byte, side)
    for i := range buf {
        buf[i] = ' '
    }
    diamond := make([]string, side)
    left, right := order, order
    top, bottom := 0, side - 1
    for ch := byte('A'); ch <= char; ch++ {
        buf[left], buf[right] = ch, ch
        diamond[top] = string(buf)
        diamond[bottom] = diamond[top]
        buf[left], buf[right] = ' ', ' '
        left--
        top++
        bottom--
        right++
    }
    return strings.Join(diamond, "\n"), nil
}