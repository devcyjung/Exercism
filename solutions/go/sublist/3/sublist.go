package sublist

func Sublist(l1, l2 []int) Relation {
	diff := len(l1) - len(l2)
    if diff > 0 {
        if sub(l2, l1) {
            return RelationSuperlist
        }
    } else if diff < 0 {
        if sub(l1, l2) {
            return RelationSublist
        }
    } else {
        if sub(l1, l2) {
            return RelationEqual
        }
    }
    return RelationUnequal
}

func sub(pattern, against []int) bool {
    if len(pattern) == 0 {
        return true
    }
    if len(against) == 0 {
        return false
    }
    pLen, aLen := len(pattern), len(against)
    p := processPattern(pattern)
    matchLen := 0
    i := 0
    for i < aLen {
        if against[i] == pattern[matchLen] {
            matchLen++
            i++
            if matchLen == pLen {
                return true
            }
        } else if matchLen == 0 {
            i++
        } else {
            matchLen = p[matchLen - 1]
        }
    }
    return false
}

func processPattern(pattern []int) []int {
    pLen := len(pattern)
    p := make([]int, pLen)
    matchLen := 0
    i := 1
    for i < pLen {
        if pattern[i] == pattern[matchLen] {
            matchLen++
            p[i] = matchLen
            i++
        } else if matchLen == 0 {
            p[i] = 0
            i++
        } else {
            matchLen = p[matchLen - 1]
        }
    }
    return p
}