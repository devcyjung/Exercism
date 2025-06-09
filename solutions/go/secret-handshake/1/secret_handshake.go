package secret

import "slices"

var actions = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) []string {
    result := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
        if (code >> i) & 1 == 1 {
            result = append(result, actions[i])
        }
    }
    if (code >> 4) & 1 == 1 {
        slices.Reverse(result)
    }
    slices.Clip(result)
    return result
}