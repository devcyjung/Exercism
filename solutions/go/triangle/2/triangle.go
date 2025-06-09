package triangle

import "slices"

type Kind int

const (
    NaT Kind = iota
    Equ
    Iso
    Sca
)

func KindFromSides(a, b, c float64) (k Kind) {
    sum := a + b + c
    sides := []float64{a, b, c}
    minimum := slices.Min(sides)
    maximum := slices.Max(sides)
    medium := sum - minimum - maximum
    switch {
    case maximum >= minimum + medium || minimum <= 0:
        return NaT
    case minimum == maximum:
        return Equ
    case minimum == medium || medium == maximum:
        return Iso
    default:
        return Sca
    }
}