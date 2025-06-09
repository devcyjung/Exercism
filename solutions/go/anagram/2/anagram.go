package anagram

import (
    "slices"
    "strings"
)

func Detect(subject string, candidates []string) []string {
    lowerSubject := strings.ToLower(subject)
    sortedSubject := []rune(lowerSubject)
    slices.Sort(sortedSubject)
	return slices.DeleteFunc(candidates, func(candidate string) bool {
        if len(subject) != len(candidate) {
            return true
        }
        lowerCandidate := strings.ToLower(candidate)
        if lowerSubject == lowerCandidate {
            return true
        }
        sortedCandidate := []rune(lowerCandidate)
        slices.Sort(sortedCandidate)
        return !slices.Equal(sortedSubject, sortedCandidate)
    })
}