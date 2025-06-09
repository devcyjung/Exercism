package chance

import "math/rand/v2"

func RollADie() int {
	return rand.IntN(20) + 1
}

func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

var animals = [...]string{
    "ant",
    "beaver",
    "cat",
    "dog",
    "elephant",
    "fox",
    "giraffe",
    "hedgehog",
}

const animalLen = len(animals)

func ShuffleAnimals() []string {
	shuffled := [animalLen]string{}
    for dstIdx, srcIdx := range rand.Perm(animalLen) {
        shuffled[dstIdx] = animals[srcIdx]
    }
    return shuffled[:]
}
