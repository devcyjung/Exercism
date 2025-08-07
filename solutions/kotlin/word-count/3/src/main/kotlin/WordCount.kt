object WordCount

fun WordCount.phrase(phrase: String) = phrase
    .lowercase()
    .map { if (it == '\'' || it.isLetterOrDigit()) it else ' ' }
    .joinToString("")
    .split(' ')
    .map { it.trimMatching('\'') }
    .filter { it.isNotEmpty() }
    .groupingBy { it }
    .eachCount()

private fun String.trimMatching(trimChar: Char): String {
    var begin = 0
    var end = this.lastIndex
    while (begin <= end && this[begin] == trimChar && this[end] == this[begin]) {
        ++begin
        --end
    }
    return if (begin <= end) this.slice(begin..end) else ""
}
