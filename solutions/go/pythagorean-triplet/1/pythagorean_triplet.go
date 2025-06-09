package pythagorean

import "slices"

type Triplet [3]int

func Range(min, max int) []Triplet {
    if min <= 0 || max <= min {
        return nil
    }
    results := make([]Triplet, 0, 100)
	for i := min; i <= max - 2; i++ {
        for j := i + 1; j <= max - 1; j++ {
            for k := j + 1; k <= max; k++ {
                if k * k == i * i + j * j {
                    results = append(results, Triplet{i, j, k})
                }
            }
        }
    }
    slices.Clip(results)
    return results
}

func Sum(p int) []Triplet {
    if p <= 0 {
        return nil
    }
    results := make([]Triplet, 0, 100)
	for i := 1; i <= (p + 1)/3; i++ {
        for j := i+1; j < p - i - j; j++ {
            k := p - i - j
            if k * k == i * i + j * j {
                results = append(results, Triplet{i, j, k})
            }
        }
    }
    slices.Clip(results)
    return results
}