package diffsquares

func SquareOfSum(n int) int {
	switch {
    case n == 0:
        return 0
    case n > 0:
        s := 0
        for i := 1; i <= n; i++ {
            s += i
        }
        return s * s
    default:
        return SquareOfSum(-n)
    }
}

func SumOfSquares(n int) int {
	switch {
    case n == 0:
        return 0
    case n > 0:
        s := 0
        for i := 1; i <= n; i++ {
            s += i * i
        }
        return s
    default:
        return SumOfSquares(-n)
    }
}

func Difference(n int) int {
	diff := SquareOfSum(n) - SumOfSquares(n)
    switch {
    case diff < 0:
        return -diff
    default:
        return diff
    }
}
