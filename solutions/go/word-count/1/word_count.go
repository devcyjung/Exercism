package wordcount

import "strings"

type Frequency map[string]int

func WordCount(phrase string) (f Frequency) {
	f = make(Frequency)
    var begin int
    var end int
    runes := []rune(strings.ToLower(phrase))
    for i, v := range runes {
        if '\'' == v {
        	if i != len(runes)-1 && isAlnum(runes[i+1]) {
                end = i
                if !isAlnum(runes[begin]) {
                    begin = i
                }
            } else {
                if isAlnum(runes[begin]) {
                    f[string(runes[begin:end+1])]++
                    begin = i
                    end = i
                }
            }
        } else if !isAlnum(v) {
            if isAlnum(runes[begin]) {
                f[string(runes[begin:end+1])]++
                begin = i
                end = i
            }
        } else {
            end = i
            if !isAlnum(runes[begin]) {
                begin = i
            }
            if i == len(runes)-1 {
                f[string(runes[begin:end+1])]++
            }
        }
    }
    return
}

func isAlnum(r rune) bool {
    return 'a' <= r && r <= 'z' || '0' <= r && r <= '9'
}
