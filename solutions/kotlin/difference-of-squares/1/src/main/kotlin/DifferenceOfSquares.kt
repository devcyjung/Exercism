class Squares(val number: Int) {
    val sumOfSquares = (1..number).map{ it * it }.sum()
    val sum = (1..number).sum()
    val squareOfSum = sum * sum
    val difference = squareOfSum - sumOfSquares
    
    fun sumOfSquares(): Int {
        return sumOfSquares
    }

    fun squareOfSum(): Int {
        return squareOfSum
    }

    fun difference(): Int {
        return difference
    }
}
