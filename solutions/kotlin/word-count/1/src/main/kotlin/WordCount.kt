object WordCount {
    fun phrase(phrase: String): Map<String, Int> = phrase
        .lowercase()
        .map { if (it == '\'' || it.isLetterOrDigit()) it else ' ' }
        .joinToString("")
        .split(" ")
        .map { it.dropWhile { it == '\'' }.dropLastWhile { it == '\'' } }
        .filter { it.isNotEmpty() }
        .groupingBy { it }
        .eachCount()
}
