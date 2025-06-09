package darts

import "math"

func Score(x, y float64) int {
	radius := math.Hypot(x, y)
    switch {
    case radius > 10:
        return 0
    case radius > 5:
        return 1
    case radius > 1:
        return 5
    default:
        return 10
    }
}
