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

func ShuffleAnimals() (s []string) {
    perm := rand.Perm(animalLen)
    animals[0], animals[1], animals[2], animals[3], animals[4], animals[5], animals[6], animals[7] = animals[perm[0]], animals[perm[1]], animals[perm[2]], animals[perm[3]], animals[perm[4]], animals[perm[5]], animals[perm[6]], animals[perm[7]]
    s = make([]string, animalLen)
    copy(s, animals[:])
    return
}
