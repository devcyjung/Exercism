package phonenumber

import (
    "fmt"
    "strings"
)

func Number(phoneNumber string) (string, error) {
	sanitized := strings.Map(removePunct, phoneNumber)
    if strings.HasPrefix(sanitized, "1") {
        sanitized = sanitized[1:]
    }
    if len(sanitized) != 10 || sanitized[0] < '2' || sanitized[3] < '2' {
        return "", fmt.Errorf("Invalid input: %s", phoneNumber)
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
    return fmt.Sprintf("(%v) %v-%v", sanitized[:3], sanitized[3:6], sanitized[6:]), nil
}