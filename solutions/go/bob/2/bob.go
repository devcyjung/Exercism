package bob

import (
    "strings"
    "unicode"
)

type Response = string

const (
    QuestionResponse 		Response = "Sure."
    YellingResponse			Response = "Whoa, chill out!"
    YellingQuestionResponse	Response = "Calm down, I know what I'm doing!"
    SilenceResponse			Response = "Fine. Be that way!"
    OtherwiseResponse		Response = "Whatever."
)

func Hey(remark string) Response {
    trimmed := strings.TrimSpace(remark)
    isYelling := strings.ContainsFunc(trimmed, hasLetter) && !strings.ContainsFunc(trimmed, hasLower)
    isAsking := strings.HasSuffix(trimmed, "?")
    switch {
    case len(trimmed) == 0:
        return SilenceResponse
    case isYelling && isAsking:
        return YellingQuestionResponse
    case isYelling:
        return YellingResponse
    case isAsking:
        return QuestionResponse
    default:
        return OtherwiseResponse
    }
}

func hasLetter(r rune) bool {
    return unicode.IsLetter(r)
}

func hasLower(r rune) bool {
    return unicode.IsLower(r)
}