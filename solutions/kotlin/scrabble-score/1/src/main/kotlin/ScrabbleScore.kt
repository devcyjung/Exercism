object ScrabbleScore {
    fun scoreLetter(c: Char): Int =
        when (c) {
            'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T' -> 1
            'D', 'G' -> 2
            'B', 'C', 'M', 'P' -> 3
            'F', 'H', 'V', 'W', 'Y' -> 4
            'K' -> 5
            'J', 'X' -> 8
            'Q', 'Z' -> 10
            else -> 0
        }

    fun scoreWord(word: String): Int = word.uppercase().map(::scoreLetter).sum()
}