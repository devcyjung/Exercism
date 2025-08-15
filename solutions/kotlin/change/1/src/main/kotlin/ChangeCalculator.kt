data class ChangeCalculator(var coins: List<Int>) {
    init {
        coins = coins.sorted()
    }
}

fun ChangeCalculator.computeMostEfficientChange(grandTotal: Int): List<Int> {
    require(grandTotal >= 0) { "Negative totals are not allowed." }
    if (grandTotal == 0) return emptyList()
    val table = Array(grandTotal + 1) {
        Array(coins.size) { 0 }
    }
    for (total in 1..grandTotal) {
        val minCoin = coins.filter { it <= total && (total - it == 0 || table[total - it].sum() > 0) }.minByOrNull {
            table[total - it].sum()
        }
        if (minCoin != null) {
            table[total] = table[total-minCoin].copyOf()
            ++table[total][coins.indexOf(minCoin)]
        }
    }
    require(table[grandTotal].sum() > 0) { "The total $grandTotal cannot be represented in the given currency." }
    return buildList {
        table[grandTotal].withIndex().forEach { (i, numCoins) ->
            repeat(numCoins) {
                add(coins[i])
            }
        }
    }
}
