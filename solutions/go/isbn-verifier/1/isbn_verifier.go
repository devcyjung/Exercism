package isbn

import "strings"

func IsValidISBN(isbn string) (b bool) {
	s := []rune(strings.ReplaceAll(isbn, "-", ""))
    if len(s) == 9 {
        s = append([]rune{'0'}, s...)
    }
    if len(s) != 10 {
        return
    }
    var digit int
    read := 0
    sum := 0
    for mul := 10; mul > 0; mul, read = mul-1, read+1 {
        if mul == 1 && s[read] == 'X' {
            digit = 10
        } else {
            digit = int(s[read] - '0')
            if digit < 0 || digit > 9 {
                return
            }
        }
        sum += digit * mul
        sum %= 11
    }
    if sum == 0 {
        b = true
    }
    return
}
