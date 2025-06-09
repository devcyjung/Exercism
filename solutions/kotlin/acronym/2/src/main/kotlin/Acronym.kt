object Acronym {
    fun generate(phrase: String) : String = phrase
        .splitToSequence(" ", "-")
        .fold(StringBuilder()) {
            acc, string ->
            string.dropWhile{
                !it.isLetter()
            }
                .firstOrNull()
                ?.uppercase()
                ?.let {
                    acc.append(it)
                }
            acc
        }
        .toString()
}
