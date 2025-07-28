class Squares(private val number: Int) {
    private val sumOfSquares = (1..number).map{ it * it }.sum()
    private val sum = (1..number).sum()
    private val squareOfSum = sum * sum
    private val difference = squareOfSum - sumOfSquares
    
    fun sumOfSquares(): Int = sumOfSquares

    fun squareOfSum(): Int = squareOfSum

    fun difference(): Int = difference
}
