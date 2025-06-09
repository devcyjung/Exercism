import kotlin.math.pow

object ResistorColorTrio {
    private val UNITS = arrayOf("ohms", "kiloohms", "megaohms", "gigaohms")

    fun text(vararg input: Color): String {
        require(input.size >= 3) { "Must provide at least 3 colors" }

        val (first, second, third) = input
        val baseValue = first.ordinal * 10 + second.ordinal
        val multiplier = 10.0.pow(third.ordinal).toInt()

        val resistance = (baseValue * multiplier)

        var numeral = resistance.toDouble()
        var unitIndex = 0
        
        while (numeral >= 1000 && unitIndex < UNITS.lastIndex) {
            numeral /= 1000
            ++unitIndex
        }

        val formatted = if (numeral % 1 == 0.0)
                            numeral.toInt().toString()
                        else "%.1f".format(numeral)
                        
        return "$formatted ${UNITS[unitIndex]}"
    }
}
