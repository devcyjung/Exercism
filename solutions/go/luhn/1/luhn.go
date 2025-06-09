package luhn

import "strings"

func Valid(id string) (b bool) {
	runes := []rune(strings.Join(strings.Fields(id), ""))
    size := len(runes)
	if size <= 1 {
        return
    }
    offset := size % 2
    sum := 0
    for i, rune := range runes {
        num := int(rune - '0')
        if num < 0 || num > 9 {
            return
        }
        if i % 2 == offset {
            doubled := num * 2
            sum += doubled
            if doubled > 9 {
                sum -= 9
            }
        } else {
            sum += num
        }
    }
    if sum % 10 == 0 {
        b = true
    }
    return
}
