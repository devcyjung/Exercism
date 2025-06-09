package darts

import "math"

func Score(x, y float64) int {
	radius := math.Sqrt(x*x + y*y)
    if radius > 10 {
        return 0
    }
    if radius > 5 {
        return 1
    }
    if radius > 1 {
        return 5
    }
    return 10
}
