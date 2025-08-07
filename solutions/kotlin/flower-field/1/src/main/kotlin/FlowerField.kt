data class FlowerFieldBoard(val board: List<String>) {
    private val annotated = board.annotated
    fun withNumbers() = annotated
}

private val List<String>.annotated
    get() = this.withIndex().map { (i, row) ->
        row.withIndex().map { (j, char) ->
            if (char == '*') return@map '*'
            var count = 0
            for (x in (i - 1).coerceAtLeast(0)..(i + 1).coerceAtMost(this.lastIndex)) {
                for (y in (j - 1).coerceAtLeast(0)..(j + 1).coerceAtMost(row.lastIndex)) {
                    if (this[x][y] == '*') ++count
                }
            }
            if (count == 0) ' ' else count.digitToChar()
        }.joinToString("")
    }