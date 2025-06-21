package phonenumber

import (
    "fmt"
    "strings"
)

func Number(phoneNumber string) (string, error) {
	sanitized := strings.TrimPrefix(strings.Map(removePunct, phoneNumber), "1")
    if len(sanitized) != 10 || sanitized[0] < '2' || sanitized[3] < '2' {
        return "", fmt.Errorf("invalid input: %s", phoneNumber)
    }
    return sanitized, nil
}

func removePunct(r rune) rune {
    if '0' <= r && r <= '9' {
        return r
    }
    return -1
}

func AreaCode(phoneNumber string) (string, error) {
	s, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return s[:3], nil
}

func Format(phoneNumber string) (string, error) {
	s, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("(%v) %v-%v", s[:3], s[3:6], s[6:]), nil
}