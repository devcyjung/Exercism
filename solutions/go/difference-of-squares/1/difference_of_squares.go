package diffsquares

func SquareOfSum(n int) (r int) {
	if n <= 0 {
        return
    }
    r = n * n * (n+1) * (n+1) / 4
    return
}

func SumOfSquares(n int) (r int) {
	if n <= 0 {
        return
    }
    for i := 1; i <= n; i++ {
        r += i * i
    }
    return
}

func Difference(n int) (r int) {
	if n <= 0 {
        return
    }
    return SquareOfSum(n) - SumOfSquares(n)
}
