package parsinglogfiles

import (
    "regexp"
    "fmt"
)

var regexmap map[string]*regexp.Regexp

func init() {
    regexmap = make(map[string]*regexp.Regexp)
    regexmap["IsValidLine"] = regexp.MustCompile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
    regexmap["SplitLogLine"] = regexp.MustCompile(`<[~*=-]*>`)
    regexmap["CountQuotedPasswords"] = regexp.MustCompile(`(?i)".*password.*"`)
    regexmap["RemoveEndOfLineText"] = regexp.MustCompile(`end-of-line[\d]+`)
    regexmap["TagWithUserName"] = regexp.MustCompile(`User[\s]+(\S*)\b`)
}

func IsValidLine(text string) bool {
    return regexmap["IsValidLine"].MatchString(text)
}

func SplitLogLine(text string) []string {
    return regexmap["SplitLogLine"].Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    cnt := 0
    for _, v := range lines {
        if regexmap["CountQuotedPasswords"].MatchString(v) {
            cnt++
        }
    }
    return cnt
}

func RemoveEndOfLineText(text string) string {
    return regexmap["RemoveEndOfLineText"].ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    res := make([]string, 0)
    for _, v := range lines {
        mat := regexmap["TagWithUserName"].FindStringSubmatch(v)
        if len(mat) > 1 {
            res = append(res, fmt.Sprintf("[USR] %v %v", mat[1], v))
        } else {
            res = append(res, v)
        }
    }
    return res
}
