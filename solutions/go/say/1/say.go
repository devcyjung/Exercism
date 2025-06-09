package say

import "strings"

var (
    units = []string{"billion", "million", "thousand", ""}
    singles = []string{
        "", "one", "two", "three", "four", "five", "six",
        "seven", "eight", "nine",
    }
    tens = []string{
        "", "", "twenty", "thirty", "forty", "fifty", "sixty",
        "seventy", "eighty", "ninety",
    }
    teens = []string{
        "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen",
        "sixteen", "seventeen", "eighteen", "nineteen",
    }
)

func Say(n int64) (string, bool) {
	if n < 0 || n > 999_999_999_999 {
        return "", false
    }
    if n == 0 {
        return "zero", true
    }
    var terms [4]int64
    for i := 3; i >= 0; i-- {
        n, terms[i] = n / 1000, n % 1000
    }
    result := make([]string, 0, 7)
    for i := 0; i < 4; i++ {
        if terms[i] == 0 {
            continue
        }
        result = append(result, parseTerm(terms[i]))
        if units[i] != "" {
            result = append(result, units[i])
        }
    }
    return strings.Join(result, " "), true
}

func parseTerm(n int64) string {
    hundreds, subHundred := n / 100, n % 100
    tenth, last := subHundred / 10, subHundred % 10
    result := make([]string, 0, 3)
    if hundreds > 0 {
        result = append(result, singles[hundreds])
        result = append(result, "hundred")
    }
    if subHundred > 0 {
        switch {
        case subHundred >= 20 && last == 0:
            result = append(result, tens[tenth])
        case subHundred >= 20:
            result = append(result, tens[tenth] + "-" + singles[last])
        case subHundred >= 10:
            result = append(result, teens[last])
        default:
            result = append(result, singles[last])
        }
    }
    return strings.Join(result, " ")
}