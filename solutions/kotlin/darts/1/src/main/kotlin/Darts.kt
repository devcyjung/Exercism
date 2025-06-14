import kotlin.math.hypot

object Darts {
    fun score(x: Number, y: Number): Int {
        val radius = hypot(x.toDouble(), y.toDouble())
        return when {
            radius <= 1.0 -> 10
            radius <= 5.0 -> 5
            radius <= 10.0 -> 1
            else -> 0
        }
    }
}