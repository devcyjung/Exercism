package phonenumber

import (
    "errors"
    "fmt"
    "strings"
)

var (
    ErrWrongLength = errors.New("The input does not have a valid length")
    ErrWrongAreaCode = errors.New("The area code starts with wrong number")
    ErrWrongExchangeCode = errors.New("The exchange code starts with wrong number")
    phoneNumberFormat = "(%v) %v-%v"
)

func Number(phoneNumber string) (r string, err error) {
	sanitized := strings.Map(removePunct, phoneNumber)
    if strings.HasPrefix(sanitized, "1") {
        sanitized = sanitized[1:]
    }
    if len(sanitized) != 10 {
        err = ErrWrongLength
        return
    }
    if sanitized[0] < '2' {
        err = ErrWrongAreaCode
    }
    if sanitized[3] < '2' {
        err = errors.Join(err, ErrWrongExchangeCode)
    }
    if err == nil {
        r = sanitized
    }
    return 
}

func removePunct(r rune) rune {
    switch {
    case '0' <= r && r <= '9':
        return r
    default:
        return -1
    }
}

func AreaCode(phoneNumber string) (r string, err error) {
    var sanitized string
	sanitized, err = Number(phoneNumber)
    if err != nil {
        return
    }
    r = sanitized[:3]
    return
}

func Format(phoneNumber string) (r string, err error) {
    var sanitized string
	sanitized, err = Number(phoneNumber)
    if err != nil {
        return
    }
    r = fmt.Sprintf(phoneNumberFormat, sanitized[:3], sanitized[3:6], sanitized[6:])
    return
}