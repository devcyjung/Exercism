package dndcharacter

import (
    "math"
    "math/rand/v2"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
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

func GenerateCharacter() Character {
    str := Ability()
    dex := Ability()
    con := Ability()
    itl := Ability()
    wis := Ability()
    chr := Ability()
	return Character{
        Strength:		str,
    	Dexterity:    	dex,
    	Constitution: 	con,
    	Intelligence: 	itl,
    	Wisdom:       	wis,
    	Charisma:     	chr,
    	Hitpoints:    	10 + Modifier(con),
    }
}

func DiceRoll() int {
    return rand.N(6) + 1
}
