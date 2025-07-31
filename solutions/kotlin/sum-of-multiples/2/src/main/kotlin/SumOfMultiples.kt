import kotlin.math.abs

object SumOfMultiples {
    fun sum(factors: Set<Int>, limit: Int): Int = factors
        .filter { it != 0 }
        .map { abs(it) }
        .flatMap { factor ->
            (1..((limit - 1) / factor))
                .map { factor * it }
        }
        .distinct()
        .sum()
}