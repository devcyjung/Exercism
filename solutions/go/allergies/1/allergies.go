package allergies

import "slices"

var ALLERGENS = [8]string {
    "eggs", "peanuts", "shellfish", "strawberries",
    "tomatoes", "chocolate", "pollen", "cats",
}

func Allergies(allergies uint) []string {
	a := make([]string, 0, len(ALLERGENS))
    for i, v := range ALLERGENS {
        if allergies >> i & 1 == 1 {
            a = append(a, v)
        }
    }
    slices.Clip(a)
    return a
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies >> slices.Index(ALLERGENS[:], allergen) & 1 == 1
}