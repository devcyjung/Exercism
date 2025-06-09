package phonenumber

import (
    "errors"
    "fmt"
)

func Number(phoneNumber string) (s string, e error) {
	var r []rune
    for _, v := range phoneNumber {
        if '0' <= v && v <= '9' {
            r = append(r, v)
        }
    }
    if len(r) == 11 {
        if r[0] == '1' {
            r = r[1:]
        } else {
            e = errors.ErrUnsupported
            return
        }
    }
    if len(r) != 10 || r[0]-'0' < 2 || r[3]-'0' < 2 {
        e = errors.ErrUnsupported
        return
    }
    s = string(r)
    return
}

func AreaCode(phoneNumber string) (r string, e error) {
	var s string
    s, e = Number(phoneNumber)
    if e != nil {
        return
    }
    r = s[:3]
    return
}

func Format(phoneNumber string) (r string, e error) {
	var s string
    s, e = Number(phoneNumber)
    if e != nil {
        return
    }
    r = fmt.Sprintf("(%s) %s-%s", s[:3], s[3:6], s[6:])
    return
}
