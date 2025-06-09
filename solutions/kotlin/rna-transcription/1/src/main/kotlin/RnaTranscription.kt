fun transcribeToRna(dna: String): String =
    dna.asSequence().map{ nucleotide ->
        when(nucleotide) {
            'G' -> 'C'
            'C' -> 'G'
            'T' -> 'A'
            'A' -> 'U'
            else -> throw IllegalArgumentException("Invalid DNA strand: $dna")
        }
    }.joinToString(separator = "")