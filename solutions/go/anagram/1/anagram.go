package anagram

import (
    "slices"
    "strings"
)

func Detect(subject string, candidates []string) (r []string) {
    subject = strings.ToUpper(subject)
    target := []rune(subject)
    size := len(target)
    targetSorted := slices.Clone(target)
    slices.Sort(targetSorted)
    for _, v := range candidates {
        upperV := strings.ToUpper(v)
        if upperV == subject {
            continue
        }
        runes := []rune(upperV)
        if len(runes) != size {
            continue
        }
        runesSorted := slices.Clone(runes)
        slices.Sort(runesSorted)
        if !slices.Equal(runesSorted, targetSorted) {
            continue
        }
        r = append(r, v)
    }
    return
}
