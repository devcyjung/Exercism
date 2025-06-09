object Pangram {
    fun isPangram(input: String): Boolean {
        return input.fold(Array(26){ false }) { array, ch ->
            when (ch) {
                in 'A'..'Z' -> array[ch - 'A'] = true
                in 'a'..'z' -> array[ch - 'a'] = true
            }
            array
        }.all{ it }
    }
}