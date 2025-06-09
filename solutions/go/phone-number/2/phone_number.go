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

func Number(phoneNumber string) (string, error) {
	sanitized := strings.Map(removePunct, phoneNumber)
    if strings.HasPrefix(sanitized, "1") {
        sanitized = sanitized[1:]
    }
    if len(sanitized) != 10 {
        return "", ErrWrongLength
    }
    if sanitized[0] < '2' {
        return "", ErrWrongAreaCode
    }
    if sanitized[3] < '2' {
        return "", ErrWrongExchangeCode
    }
    return sanitized, nil
}

func removePunct(r rune) rune {
    switch {
    case '0' <= r && r <= '9':
        return r
    default:
        return -1
    }
}

func AreaCode(phoneNumber string) (string, error) {
	sanitized, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return sanitized[:3], nil
}

func Format(phoneNumber string) (string, error) {
	sanitized, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf(phoneNumberFormat, sanitized[:3], sanitized[3:6], sanitized[6:]), nil
}