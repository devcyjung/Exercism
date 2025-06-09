package sublist

import "slices"

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

func sub(small, big []int) bool {
    size := len(small)
    if size == 0 {
        return true
    }
    head := small[0]
    index := 0
    for len(big) >= size {
        index = slices.Index(big, head)
        if index < 0 {
            break
        }
        big = big[index:]
        if len(big) < size {
            break
        }
        if slices.Equal(small, big[:size]) {
            return true
        }
        big = big[1:]
    }
    return false
}