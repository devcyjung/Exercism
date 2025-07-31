class Anagram(val source: String) {
    val lowercaseSource by lazy { source.lowercase() }
    val sortedSource by lazy { lowercaseSource.toList().sorted() }

    fun match(anagrams: Collection<String>): Set<String> = anagrams
        .filter {
            val lowercaseIt = it.lowercase()
            lowercaseIt != lowercaseSource && lowercaseIt.toList().sorted() == sortedSource
        }
        .toSet()
}
