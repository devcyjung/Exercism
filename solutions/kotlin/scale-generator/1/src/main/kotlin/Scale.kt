class Scale(private val tonic: String) {
    init {
        require(tonic.length <= 2 && tonic.length > 0) {
            "Tonic must have length 1 or 2"
        }
        require(tonic.get(0).lowercase() in "a".."g") {
            "Tonic letter is invalid"
        }
        if (tonic.length == 2) {
            require(tonic[1] == '#' || tonic[1] == 'b') {
                "Tonic's second letter must be either '#' or 'b'"
            }
        }
    }

    val isSharp = when (tonic.length) {
        1 -> when (tonic.lowercase()) {
            "a", "b", "e" -> true
            "f" -> false
            else -> tonic.get(0).isUpperCase() 
        }
        else -> tonic.get(1) == '#'
    }

    fun series(): Sequence<String> =
        generateSequence(tonic.replaceFirstChar { it.uppercase() }) { str ->
            when (str.length) {
                2 -> if (str[1] == 'b')
                        str[0].toString()
                     else
                        ('A' + ((str[0] - 'A' + 1) % 7)).toString()
                else -> when (str[0]) {
                    'B', 'E' -> (str[0] + 1).toString()
                    else -> if (isSharp)
                                str + '#'
                            else
                                ('A' + ((str[0] - 'A' + 1) % 7)).toString() + 'b'
                }
            }
        } 
    
    fun chromatic(): List<String> = series().take(12).toList()

    fun interval(intervals: String): List<String> =
        intervals.fold(series().drop(1) to sequenceOf(series().first())) { (chromatic, scale), current ->
            when (current) {
                'm' -> chromatic.drop(1) to (scale + chromatic.first()) 
                'M' -> chromatic.drop(2) to (scale + chromatic.drop(1).first())
                'A' -> chromatic.drop(3) to (scale + chromatic.drop(2).first())
                else -> throw UnsupportedOperationException("Unsupported interval symbol: $current")
            }
        }.second.toList().dropLast(1)
}