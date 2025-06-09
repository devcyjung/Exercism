package sublist

func Sublist(firstList, secondList []int) Relation {
	len1, len2 := len(firstList), len(secondList)
    if len1 == len2 {
        firstInSecond := matchAgainst(firstList, secondList)
        secondInFirst := matchAgainst(secondList, firstList)
        if firstInSecond && secondInFirst {
            return RelationEqual
        } else if firstInSecond {
            return RelationSublist
        } else if secondInFirst {
            return RelationSuperlist
        }
    } else if len1 < len2 {
        if matchAgainst(firstList, secondList) {
            return RelationSublist
        }
    } else {
        if matchAgainst(secondList, firstList) {
            return RelationSuperlist
        }
    }
    return RelationUnequal
}

func matchAgainst(pattern, against []int) bool {
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