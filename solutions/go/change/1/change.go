package change

import (
    "errors"
    "slices"
)

var ErrInvalid = errors.New("Invalid target number")

func Change(coins []int, target int) ([]int, error) {
    if target < 0 {
        return nil, ErrInvalid
    }
    nCoins := make([]int, target + 1)
    coinLists := make([][]int, target + 1)
    nCoins[0], coinLists[0] = 0, []int{}
    for amount := 1; amount <= target; amount++ {
        var minCoin, minCoinIdx int
        var found bool
        for i, coin := range coins {
            if amount - coin >= 0 && coinLists[amount - coin] != nil {
                nc := 1 + nCoins[amount - coin]
                if !found || nc < minCoin {
                    minCoin, minCoinIdx, found = nc, i, true
                } 
            }
        }
        if found {
            nCoins[amount] = minCoin
            cloned := slices.Clone(coinLists[amount - coins[minCoinIdx]])
            coinLists[amount] = append(cloned, coins[minCoinIdx])
        }
    }
    if coinLists[target] == nil {
        return nil, ErrInvalid
    }
    slices.Sort(coinLists[target])
    return coinLists[target], nil
}