package sublist

import (
    "fmt"
    "strconv"
    "strings"
)

var stringifyFormat = "[%v]"

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
    stringified1 := make([]string, len(l1))
    stringified2 := make([]string, len(l2))
    for i := range stringified1 {
        stringified1[i] = fmt.Sprintf(stringifyFormat, strconv.Itoa(l1[i]))
    }
    for i := range stringified2 {
        stringified2[i] = fmt.Sprintf(stringifyFormat, strconv.Itoa(l2[i]))
    }
    return strings.Contains(strings.Join(stringified2, ","), strings.Join(stringified1, ","))
}