object WordCount {
    fun phrase(phrase: String): Map<String, Int> = phrase
        .lowercase()
        .map { if (it == '\'' || it.isLetterOrDigit()) it else ' ' }
        .joinToString("")
        .split(' ')
        .map {
            var begin = 0
            var end = it.lastIndex
            while (begin <= end && it[begin] == '\'' && it[end] == it[begin]) {
                ++begin
                --end
            }
            if (begin <= end) it.slice(begin..end) else ""
        }
        .filter { it.isNotEmpty() }
        .groupingBy { it }
        .eachCount()
}
