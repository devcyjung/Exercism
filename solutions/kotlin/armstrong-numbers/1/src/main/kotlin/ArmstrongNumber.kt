import kotlin.math.*

object ArmstrongNumber {

    fun check(input: Int): Boolean = input
        .toDouble()
        .let {
            log10(it)
        }
        .toInt()
        .let { nDigits ->
            var input = input
            (0..nDigits)
                .map { i ->
                    (input % 10)
                        .also {
                            input /= 10
                        }
                        .toDouble()
                        .pow(nDigits + 1)
                        .toInt()
                }
                .sum()
        } == input
}
