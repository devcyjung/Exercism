package house

import (
    "fmt"
    "strings"
)

const (
    starterFmt = "This is the %v"
    repeaterFmt = "that %v the %v"
)

var (
    nouns = []string{
        "", "house that Jack built", "malt", "rat", "cat", "dog", "cow with the crumpled horn",
        "maiden all forlorn", "man all tattered and torn", "priest all shaven and shorn",
        "rooster that crowed in the morn", "farmer sowing his corn", "horse and the hound and the horn",
    }
    verbs = []string{
        "", "lay in", "ate", "killed", "worried", "tossed", "milked", "kissed", "married", "woke",
        "kept", "belonged to",
    }
)

func Verse(v int) string {
    if !(1 <= v && v <= 12) {
        return ""
    }
    result := make([]string, 0, 12)
    result = append(result, fmt.Sprintf(starterFmt, nouns[v]))
    for i := v - 1; i >= 1; i-- {
        result = append(result, fmt.Sprintf(repeaterFmt, verbs[i], nouns[i]))
    }
	return strings.Join(result, "\n") + "."
}

func Song() string {
	result := make([]string, 0, 12)
    for i := 1; i <= 12; i++ {
        result = append(result, Verse(i))
    }
    return strings.Join(result, "\n\n")
}