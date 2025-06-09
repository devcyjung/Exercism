package parsinglogfiles

import (
    "regexp"
    "fmt"
)

var RegexIsValidLine = regexp.MustCompile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
var RegexSplitLogLine = regexp.MustCompile(`<[~*=-]*>`)
var RegexCountQuotedPasswords = regexp.MustCompile(`(?i)".*password.*"`)
var RegexRemoveEndOfLineText = regexp.MustCompile(`end-of-line[\d]+`)
var RegexTagWithUserName = regexp.MustCompile(`User[\s]+(\S*)\b`)

func IsValidLine(text string) bool {
    return RegexIsValidLine.MatchString(text)
}

func SplitLogLine(text string) []string {
    return RegexSplitLogLine.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    cnt := 0
    for _, v := range lines {
        if RegexCountQuotedPasswords.MatchString(v) {
            cnt++
        }
    }
    return cnt
}

func RemoveEndOfLineText(text string) string {
    return RegexRemoveEndOfLineText.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    res := make([]string, 0)
    for _, v := range lines {
        mat := RegexTagWithUserName.FindStringSubmatch(v)
        if len(mat) > 1 {
            res = append(res, fmt.Sprintf("[USR] %v %v", mat[1], v))
        } else {
            res = append(res, v)
        }
    }
    return res
}
