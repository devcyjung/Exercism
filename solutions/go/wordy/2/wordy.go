package wordy

import (
    "strconv"
    "strings"
)

const (
    prefix = "What is"
    suffix = "?"
)

var replacer = strings.NewReplacer("plus", "+", "minus", "-", "multiplied by", "*", "divided by", "/")

func Answer(question string) (int, bool) {
	prefixTrimmed, ok1 := strings.CutPrefix(question, prefix)
    trimmed, ok2 := strings.CutSuffix(prefixTrimmed, suffix)
    if !ok1 || !ok2 {
        return 0, false
    }
    fields := strings.Fields(replacer.Replace(trimmed))
    var operator string
    acc := 0
    if len(fields) % 2 != 1 {
        return 0, false
    }
    for i, field := range fields {
        switch i % 2 {
        case 0:
            operand, err := strconv.Atoi(field)
            if err != nil {
                return 0, false
            }
            switch operator {
            case "":
                acc += operand
            case "+":
                acc += operand
            case "-":
                acc -= operand
            case "*":
                acc *= operand
            case "/":
                if operand == 0 {
                    return 0, false
                }
                acc /= operand
            default:
                return 0, false
            }
        default:
            operator = field
        }
    }
    return acc, true
}