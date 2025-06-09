object Raindrops {
    val soundMapping = arrayOf(3 to "Pling", 5 to "Plang", 7 to "Plong")
    fun convert(n: Int): String =
        soundMapping.filter{ (divisor, _) ->
            n % divisor == 0
        }.map{ (_, sound) ->
            sound
        }.joinToString(separator = "").let{ result ->
            if (result.isEmpty()) n.toString() else result
        }
}
