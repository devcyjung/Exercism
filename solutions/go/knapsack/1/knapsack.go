package knapsack

type Item struct {
	Weight, Value int
}

func Knapsack(maximumWeight int, items []Item) int {
    rowLen := maximumWeight + 1
    memo := make([]int, (len(items) + 1) * rowLen)
    for i := 1; i <= len(items); i++ {
        itemWeight := items[i - 1].Weight
        itemValue := items[i - 1].Value
        for j := 1; j <= maximumWeight; j++ {
            withoutCurrent := memo[(i - 1) * rowLen + j]
            if itemWeight > j {
                memo[i * rowLen + j] = withoutCurrent
                continue
            }
            withCurrent := memo[(i - 1) * rowLen + j - itemWeight] + itemValue
            switch {
            case withoutCurrent > withCurrent:
                memo[i * rowLen + j] = withoutCurrent
            default:
                memo[i * rowLen + j] = withCurrent
            }
        }
    }
    return memo[len(items) * rowLen + maximumWeight]
}