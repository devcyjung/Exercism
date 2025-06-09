package summultiples

func SumMultiples(limit int, divisors ...int) (s int) {
    m := make(map[int]bool)
	for _, d := range divisors {
        if d == 0 {
            continue
        }
        for i := 1; i <= (limit-1) / d; i++ {
            m[i*d] = true
        }
    }
    for k := range m {
        s += k
    }
    return
}
