package bob

import "strings"

type Response = string

const (
    NormalQuestionResponse 	Response = "Sure."
    YellingResponse			Response = "Whoa, chill out!"
    YellingQuestionResponse	Response = "Calm down, I know what I'm doing!"
    SilenceResponse			Response = "Fine. Be that way!"
    OtherwiseResponse		Response = "Whatever."
)

func Hey(remark string) Response {
    trimmed := strings.TrimSpace(remark)
    question := strings.HasSuffix(trimmed, "?")
    yell := remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
    silent := trimmed == ""
    if !yell && question {
        return NormalQuestionResponse
    }
    if yell && !question {
        return YellingResponse
    }
    if yell && question {
        return YellingQuestionResponse
    }
    if silent {
        return SilenceResponse
    }
    return OtherwiseResponse
}
