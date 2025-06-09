package diamond

import (
    "errors"
    "strings"
)

var ErrInvalidRange = errors.New("Input is not in valid range")

func Gen(char byte) (string, error) {
    order := rune(char - 'A')
    if !(0 <= order && order <= 25) {
        return "", ErrInvalidRange
    }
    size := order * 2 + 1
    var b strings.Builder
    for i := rune(0); i < size; i++ {
        for j := rune(0); j < size; j++ {
            switch {
            case i + j == order:
                b.WriteRune('A' + i)
        	case i + j == 3 * order:
                b.WriteRune('A' + 2 * order - i)
        	case i == j + order:
                b.WriteRune('A' + 2 * order - i)
            case j == i + order:
                b.WriteRune('A' + i)
            default:
                b.WriteRune(' ')
            }
        }
        if i != size - 1 {
            b.WriteRune('\n')
        }
    }
    return b.String(), nil
}