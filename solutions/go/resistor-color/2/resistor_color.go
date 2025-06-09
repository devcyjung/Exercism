package resistorcolor

import "slices"

var colors = [10]string {
    "black", "brown", "red", "orange", "yellow",
    "green", "blue", "violet", "grey", "white",
}

func Colors() (s []string) {
    return slices.Clone(colors[:])
}

func ColorCode(color string) int {
	return slices.Index(Colors(), color)
}