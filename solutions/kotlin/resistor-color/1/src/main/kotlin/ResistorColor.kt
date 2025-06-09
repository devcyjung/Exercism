object ResistorColor {
    val COLORS = listOf("black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white")

    fun colorCode(input: String): Int =
        COLORS.indexOf(input)

    fun colors(): List<String> =
        COLORS
}
