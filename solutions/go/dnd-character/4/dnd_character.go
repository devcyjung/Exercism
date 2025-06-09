package dndcharacter

import "math/rand/v2"

type Character struct {
	Strength, Dexterity, Constitution,
    Intelligence, Wisdom, Charisma, Hitpoints int
}

func Modifier(score int) int {
    if (score - 10) % 2 < 0 {
        return (score - 10) / 2 - 1
    }
	return (score - 10) / 2
}

func Ability() (sum int) {
	min := DiceRoll()
    for i := 0; i < 3; i++ {
        d := DiceRoll()
        if min <= d {
            sum += d
        } else {
            min, sum = d, sum + min
        }
    }
    return
}

func Abilities() (str, dex, con, itl, wis, chr int) {
    return Ability(), Ability(), Ability(),
    	Ability(), Ability(), Ability()
}

func GenerateCharacter() Character {
    str, dex, con, itl, wis, chr := Abilities()
	return Character{
        str, dex, con, itl,
        wis, chr, 10 + Modifier(con),
    }
}

func DiceRoll() int {
    return rand.N(6) + 1
}