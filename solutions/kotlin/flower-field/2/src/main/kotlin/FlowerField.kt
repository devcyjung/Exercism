data class FlowerFieldBoard(val board: List<String>) {
    private val annotated = board.annotated
    fun withNumbers() = annotated
}

private val List<String>.annotated
    get() = this.withIndex().map {(i, row) ->
        row.withIndex().map {(j, char) ->
            if (char == '*') return@map '*'
            val count = ((i - 1).coerceAtLeast(0)..(i + 1).coerceAtMost(this.lastIndex))
                .flatMap {x ->
                    ((j - 1).coerceAtLeast(0)..(j + 1).coerceAtMost(row.lastIndex))
                        .map {y -> x to y}
                }
                .count {(x, y) -> this[x][y] == '*'}
            if (count == 0) ' ' else count.digitToChar()
        }.joinToString("")
    }