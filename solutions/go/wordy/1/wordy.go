package wordy

import (
    "regexp"
    "strconv"
    "strings"
    "unicode"
)

var (
    PLUS = regexp.MustCompile(`plus`)
    MINUS = regexp.MustCompile(`minus`)
    MULTIPLY = regexp.MustCompile(`multiplied by`)
    DIVIDE = regexp.MustCompile(`divided by`)
)

func Answer(question string) (int, bool) {
	if !strings.HasPrefix(question, "What is ") || !strings.HasSuffix(question, "?") {
        return 0, false
    }
    trimmed := strings.TrimPrefix(strings.TrimSuffix(question, "?"), "What is ")
    trimmed = PLUS.ReplaceAllString(trimmed, " + ")
    trimmed = MINUS.ReplaceAllString(trimmed, " - ")
    trimmed = MULTIPLY.ReplaceAllString(trimmed, " * ")
    trimmed = DIVIDE.ReplaceAllString(trimmed, " / ")
    if strings.ContainsFunc(trimmed, nonEquation) {
        return 0, false
    }
    var accumulator, operand int
    var operator string
    var err error
    fields := strings.Fields(trimmed)
    if !(len(fields) > 0 && len(fields) % 2 == 1) {
        return 0, false
    }
    for i, v := range fields {
        switch {
        case i % 2 == 0:
            operand, err = strconv.Atoi(v)
            if err != nil {
                return 0, false
            }
            if i == 0 {
                operator = "+"
            }
            switch operator {
            case "+":
                accumulator += operand
            case "-":
                accumulator -= operand
            case "*":
                accumulator *= operand
            case "/":
                if operand == 0 {
                    return 0, false
                }
                accumulator /= operand
            }
        default:
            if v != "+" && v != "-" && v != "*" && v != "/" {
                return 0, false
            }
            operator = v
        }
    }
    return accumulator, true
}

func nonEquation(r rune) bool {
    return !('0' <= r && r <= '9') && r != '+' && r != '-' && r != '*' && r != '/' && !unicode.IsSpace(r)
}