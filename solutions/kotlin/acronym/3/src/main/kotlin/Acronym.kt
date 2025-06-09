object Acronym {
    fun generate(phrase: String)
        = phrase
            .replace(Regex("""[^\p{L}-\s]"""), "")
            .splitToSequence(Regex("""[-\s]+"""))
            .map{ word ->
                word.getOrNull(0)?.uppercase()
            }
            .filterNotNull()
            .joinToString(separator = "")
}