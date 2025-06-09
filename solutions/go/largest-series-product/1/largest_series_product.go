package lsproduct

import (
    "errors"
    "strconv"
    "strings"
)

func LargestSeriesProduct(digits string, span int) (r int64, e error) {
	if len(digits) < span || span < 0 {
        e = errors.ErrUnsupported
        return
    }
    for _, s := range strings.Split(digits, "0") {
        v := []rune(s)
        if len(v) < span {
            continue
        }
        prod := int64(1)
        var m, n int
        for i := 0; i < span; i++ {
            n, e = strconv.Atoi(string(v[i]))
            if e != nil {
                r = int64(0)
                return
            }
            prod *= int64(n)
        }
        local := prod
        for i := 1; i+span <= len(v); i++ {
            n, e = strconv.Atoi(string(v[i+span-1]))
            if e != nil {
                r = int64(0)
                return
            }
            m, e = strconv.Atoi(string(v[i-1]))
            if e != nil {
                r = int64(0)
                return
            }
            prod /= int64(m)
            prod *= int64(n)
            if local < prod {
                local = prod
            }
        }
        if r < local {
            r = local
        }
    }
    return
}
