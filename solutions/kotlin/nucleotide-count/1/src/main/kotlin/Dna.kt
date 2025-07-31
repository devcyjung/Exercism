class Dna(dna: String) {
    val counter: Map<Char, Int>
    init {
        counter = dna
            .groupingBy(Char::uppercaseChar)
            .eachCountTo(mutableMapOf('A' to 0, 'C' to 0, 'G' to 0, 'T' to 0))
        require(counter.keys.all { "ACGT".contains(it) }) { "Invalid DNA" }
    }
    
    val nucleotideCounts: Map<Char, Int>
        get() = counter
}
