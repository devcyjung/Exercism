package triangle

import "slices"

type Kind string

const (
    NaT Kind = "not a triangle"
    Equ Kind = "equilateral"
    Iso Kind = "isosceles"
    Sca Kind = "scalene"
)

func KindFromSides(a, b, c float64) (k Kind) {
    sorted := []float64{a, b, c}
    slices.Sort(sorted)
    if sorted[2] > sorted[1] + sorted[0] || sorted[0] <= 0 {
        k = NaT
        return
    }
    k = Sca
    if sorted[0] == sorted[1] || sorted[1] == sorted[2] {
        k = Iso
    }
    if sorted[0] == sorted[2] {
        k = Equ
    }
	return
}
