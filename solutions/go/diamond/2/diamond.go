package diamond

import (
    "errors"
    "strings"
)

var ErrInvalidRange = errors.New("Input is not in valid range")

func Gen(char byte) (string, error) {
    order := char - 'A'
    if !(0 <= order && order <= 25) {
        return "", ErrInvalidRange
    }
    size := order * 2 + 1
    var b strings.Builder
    for i := byte(0); i < size; i++ {
        for j := byte(0); j < size; j++ {
            switch {
            case i + j == order:
                b.WriteByte('A' + i)
        	case i + j == 3 * order:
                b.WriteByte('A' + 2 * order - i)
        	case i == j + order:
                b.WriteByte('A' + 2 * order - i)
            case j == i + order:
                b.WriteByte('A' + i)
            default:
                b.WriteByte(' ')
            }
        }
        if i != size - 1 {
            b.WriteByte('\n')
        }
    }
    return b.String(), nil
}