object ResistorColorTrio {
    val UNITS = arrayOf("ohms", "kiloohms", "megaohms", "gigaohms")
    fun text(vararg input: Color): String {
        require(input.size >= 3) {
            "Must provide at least 3 colors"
        }
        var trailingZeros = input.get(2).ordinal
        var digits = input.get(0).ordinal * 10 + input.get(1).ordinal
        while (digits % 10 == 0) {
            digits /= 10
            ++trailingZeros
        }
        val unit = UNITS.get(trailingZeros / 3)
        val zeros = "0".repeat(trailingZeros % 3)
        return "$digits$zeros $unit"
    }
}