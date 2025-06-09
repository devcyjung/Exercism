fun translate(rna: String?): List<String> {
    return rna?.let { rna ->
        rna.chunkedSequence(3) { it }.takeWhile { codon ->
            codon !in setOf("UAA", "UAG", "UGA")
        }.map{ codon ->
            when (codon) {
                "AUG" -> "Methionine"
                "UUU", "UUC" -> "Phenylalanine"
                "UUA", "UUG" -> "Leucine"
                "UCU", "UCC", "UCA", "UCG" -> "Serine"
                "UAU", "UAC" -> "Tyrosine"
                "UGU", "UGC" -> "Cysteine"
                "UGG" -> "Tryptophan"
                else -> throw IllegalArgumentException("Invalid codon")
            }
        }.toList()
    } ?: emptyList()
}