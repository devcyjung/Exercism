package yacht

import "slices"

func Score(dice []int, category string) int {
    if len(dice) != 5 {
        return -1
    }
    maximum, minimum := slices.Max(dice), slices.Min(dice)
    if maximum > 6 || minimum < 1 {
        return -1
    }
	switch category {
        case "ones":
        	return count(dice, 1) * 1
        case "twos":
        	return count(dice, 2) * 2
        case "threes":
        	return count(dice, 3) * 3
        case "fours":
        	return count(dice, 4) * 4
        case "fives":
        	return count(dice, 5) * 5
        case "sixes":
        	return count(dice, 6) * 6
        case "full house":
        	s := sortedByCount(dice)
        	if len(s) == 2 && s[0].count == 3 {
                return s[0].num * 3 + s[1].num * 2
            }
        	return 0
        case "four of a kind":
        	s := sortedByCount(dice)
        	if s[0].count >= 4 {
                return 4 * s[0].num
            }
        	return 0
        case "little straight":
        	s := sortedByCount(dice)
        	if len(s) == 5 && maximum == 5 && minimum == 1 {
                return 30
            }
        	return 0
        case "big straight":
        	s := sortedByCount(dice)
        	if len(s) == 5 && maximum == 6 && minimum == 2 {
                return 30
            }
        	return 0
        case "choice":
        	acc := 0
        	for _, roll := range dice {
                acc += roll
            }
        	return acc
        case "yacht":
        	s := sortedByCount(dice)
        	if len(s) == 1 {
                return 50
            }
        	return 0
    }
    return -1
}

func count(dice []int, value int) int {
    acc := 0
    for _, roll := range dice {
        if roll == value {
            acc++
        }
    }
    return acc
}

type pair struct {
    num, count	int
}

func sortedByCount(dice []int) []pair {
    rolls := make([]pair, 0, 6)
    for i := 1; i <= 6; i++ {
        rolls = append(rolls, pair{i, count(dice, i)})
    }
    rolls = slices.DeleteFunc(rolls, func(p pair) bool {
        return p.count == 0
    })
    slices.SortFunc(rolls, func(a, b pair) int {
        return b.count - a.count
    })
    return rolls
}