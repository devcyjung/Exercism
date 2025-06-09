package armstrong

func IsNumber(n int) bool {
    if n < 0 {
        return false
    }
	temp := n
    pow := 1
    for temp >= 10 {
        temp /= 10
        pow++
    }
    sum := 0
    temp2 := n
    for temp2 > 0 {
        temp2, temp = temp2 / 10, temp2 % 10
        sum += power(temp, pow)
    }
    return sum == n
}

func power(base, pow int) int {
    mul := 1
    for i := 0; i < pow; i++ {
        mul *= base
    }
    return mul
}