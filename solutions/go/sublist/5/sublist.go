package sublist

import (
    "fmt"
    "strconv"
    "strings"
)

const (
    elemFormat = "[%v]"
    joinSep = ","
)

func Sublist(l1, l2 []int) Relation {
    diff := len(l1) - len(l2)
    switch {
    case diff > 0 && sub(l2, l1):
        return RelationSuperlist
    case diff < 0 && sub(l1, l2):
    	return RelationSublist
    case diff == 0 && sub(l1, l2):
        return RelationEqual
    default:
        return RelationUnequal
    }
}

func sub(l1, l2 []int) bool {
    strList1 := make([]string, len(l1))
    strList2 := make([]string, len(l2))
    for i := range strList1 {
        strList1[i] = fmt.Sprintf(elemFormat, strconv.Itoa(l1[i]))
    }
    for i := range strList2 {
        strList2[i] = fmt.Sprintf(elemFormat, strconv.Itoa(l2[i]))
    }
    return strings.Contains(strings.Join(strList2, joinSep), strings.Join(strList1, joinSep))
}