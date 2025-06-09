package yacht

import "slices"

func Score(dice []int, category string) int {
	switch category {
        case "ones":
        	return Dup(dice, 1)
        case "twos":
        	return Dup(dice, 2)
        case "threes":
        	return Dup(dice, 3)
        case "fours":
        	return Dup(dice, 4)
        case "fives":
        	return Dup(dice, 5)
        case "sixes":
        	return Dup(dice, 6)
        case "full house":
        	return FullHouse(dice)
        case "four of a kind":
        	return FourOfaKind(dice)
        case "little straight":
        	return LittleStraight(dice)
        case "big straight":
        	return BigStraight(dice)
        case "choice":
        	return Choice(dice)
        case "yacht":
        	return Yacht(dice)
    }
    return -1
}

func Dup(dice []int, dup int) (r int) {
    for _, d := range dice {
        if d == dup {
            r += dup
        }
    }
    return
}

func FullHouse(dice []int) int {
    slices.Sort(dice)
    if dice[0] == dice[2] && dice[3] == dice[4] && dice[0] != dice[4] {
        return dice[0] * 3 + dice[3] * 2
    }
    if dice[0] == dice[1] && dice[2] == dice[4] && dice[0] != dice[4] {
        return dice[0] * 2 + dice[3] * 3
    }
    return 0
}

func FourOfaKind(dice []int) int {
    slices.Sort(dice)
    if dice[0] == dice[3] {
        return 4 * dice[0]
    }
    if dice[1] == dice[4] {
        return 4 * dice[1]
    }
    return 0
}

func LittleStraight(dice []int) int {
    for i := 1; i <= 5; i++ {
        if !slices.Contains(dice, i) {
            return 0
        }
    }
    return 30
}

func BigStraight(dice []int) int {
    for i := 2; i <= 6; i++ {
        if !slices.Contains(dice, i) {
            return 0
        }
    }
    return 30
}

func Choice(dice []int) (s int) {
    for _, d := range dice {
        s += d
    }
    return
}

func Yacht(dice []int) int {
    slices.Sort(dice)
    if dice[0] == dice[4] {
        return 50
    }
    return 0
}