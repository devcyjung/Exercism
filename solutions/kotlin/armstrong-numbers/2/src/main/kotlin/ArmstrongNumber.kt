object ArmstrongNumber {
    fun check(input: Int): Boolean {
        require(input >= 0) {
            "Must provide non-negative integer"
        }
        val numOfDigits = series(input).count()
        val armstrongSum = series(input)
            .map { digit ->
                var power = digit
                repeat(numOfDigits - 1) {
                    power *= digit
                }
                power
            }.sum()
        return input == armstrongSum
    }

    fun series(input: Int): Sequence<Int> =
        generateSequence(0 to input) { (_, n) ->
            if (n == 0) null else n % 10 to n / 10
        }.drop(1).map { (digit, _) ->
            digit
        }
}
