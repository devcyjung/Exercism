package dndcharacter

import (
    "math/rand/v2"
    "slices"
)

type Character struct {
	Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma, Hitpoints int
}

func Modifier(score int) int {
	return (score - 10) >> 1
}

func Ability() int {
	rolls := make([]int, 4)
    sum := 0
    for i := range rolls {
        rolls[i] = diceRoll()
        sum += rolls[i]
    }
    sum -= slices.Min(rolls)
    return sum
}

func GenerateCharacter() Character {
    str, dex, con, itl, wis, chr := abilities()
	return Character{
        str, dex, con, itl, wis, chr, 10 + Modifier(con),
    }
}

func diceRoll() int {
    return rand.N(6) + 1
}

func abilities() (str, dex, con, itl, wis, chr int) {
    return Ability(), Ability(), Ability(), Ability(), Ability(), Ability()
}