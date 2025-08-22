data class FlowerFieldBoard(private val board: List<String>) {
    private val annotated = board.annotated
    fun withNumbers() = annotated
}

private val List<String>.annotated
    get() = this.withIndex().map {(i, row) ->
        row.withIndex().asSequence().map {(j, char) ->
            if (char == '*') '*'
            else ((i - 1)..(i + 1))
                .flatMap {x -> ((j - 1)..(j + 1)).map {y -> x to y}}
                .count {(x, y) -> this.getOrNull(x)?.getOrNull(y) == '*'}
                .let { if (it == 0) ' ' else it.digitToChar() }
        }.joinToString("")
    }