object WordCount

fun WordCount.phrase(phrase: String) = phrase
    .lowercase()
    .splitBy { it != '\'' && !it.isLetterOrDigit() }
    .map { it.trimMatching('\'') }
    .filter { it.isNotEmpty() }
    .groupingBy { it }
    .eachCount()

private fun String.splitBy(isDelimiter: (Char) -> Boolean): List<String> = buildList {
    val sourceStr = this@splitBy
    var beginIndex = -1
    sourceStr.forEachIndexed { index, char ->
        if (isDelimiter(char)) {
            this.add(sourceStr.slice(beginIndex + 1..<index))
            beginIndex = index
        }
    }
    this.add(sourceStr.slice(beginIndex + 1..sourceStr.lastIndex))
}

private fun String.trimMatching(trimChar: Char): String {
    var begin = 0
    var end = this.lastIndex
    while (begin <= end && this[begin] == trimChar && this[end] == this[begin]) {
        ++begin
        --end
    }
    return if (begin <= end) this.slice(begin..end) else ""
}
