enum class Color {
    BLACK, BROWN, RED, ORANGE, YELLOW,
    GREEN, BLUE, VIOLET, GREY, WHITE,
}

object ResistorColor {
    fun colorCode(input: String): Int =
        Color.valueOf(input.uppercase()).ordinal

    fun colors(): List<String> =
        Color.entries.map{ entry ->
            entry.name.lowercase()
        }.toList()
}
