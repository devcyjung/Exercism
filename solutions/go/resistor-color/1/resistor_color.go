package resistorcolor

import "slices"

// Colors returns the list of all colors.
func Colors() (s []string) {
    s = append(s, "black")
    s = append(s, "brown")
    s = append(s, "red")
    s = append(s, "orange")
    s = append(s, "yellow")
    s = append(s, "green")
    s = append(s, "blue")
    s = append(s, "violet")
    s = append(s, "grey")
    s = append(s, "white")
	return 
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return slices.Index(Colors(), color)
}
