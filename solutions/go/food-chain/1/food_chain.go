package foodchain

import (
    "fmt"
    "strings"
)

const (
    starterFmt = "I know an old lady who swallowed a %v."
    repeaterFmt = "She swallowed the %v to catch the %v."
    enderFmt = "I don't know why she swallowed the %v. Perhaps she'll die."
    spiderFmt = "She swallowed the %v to catch the %v that wriggled and jiggled and tickled inside her."
)

var (
    punchLines = []string{
        "", "", "It wriggled and jiggled and tickled inside her.",
        "How absurd to swallow a bird!", "Imagine that, to swallow a cat!",
        "What a hog, to swallow a dog!", "Just opened her throat and swallowed a goat!",
        "I don't know how she swallowed a cow!", "She's dead, of course!",
    }
    animals = []string{
        "", "fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse",
    }
)

func Verse(v int) string {
    if !(1 <= v && v <= 8) {
        return ""
    }
	lines := make([]string, 0, 9)
    lines = append(lines, fmt.Sprintf(starterFmt, animals[v]))
    if len(punchLines[v]) != 0 {
        lines = append(lines, punchLines[v])
    }
    if v != 8 {
        for i := v - 1; i >= 1; i-- {
            switch i {
            case 2:
                lines = append(lines, fmt.Sprintf(spiderFmt, animals[i + 1], animals[i]))
            default:
                lines = append(lines, fmt.Sprintf(repeaterFmt, animals[i + 1], animals[i]))
            }
        }
        lines = append(lines, fmt.Sprintf(enderFmt, animals[1]))
    }
    return strings.Join(lines, "\n")
}

func Verses(start, end int) string {
    if !(1 <= start && start <= end && end <= 8) {
        return ""
    }
    result := make([]string, 0, end - start + 1)
	for i := start; i <= end; i++ {
        result = append(result, Verse(i))
    }
    return strings.Join(result, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}