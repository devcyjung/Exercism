fun translate(rna: String?): List<String> {
    if (rna == null) {
        return emptyList()
    }
    val result = mutableListOf<String>()
    for (codon in rna.chunked(3)) {
        val protein = when (codon) {
            "AUG" -> "Methionine"
            "UUU", "UUC" -> "Phenylalanine"
            "UUA", "UUG" -> "Leucine"
            "UCU", "UCC", "UCA", "UCG" -> "Serine"
            "UAU", "UAC" -> "Tyrosine"
            "UGU", "UGC" -> "Cysteine"
            "UGG" -> "Tryptophan"
            "UAA", "UAG", "UGA" -> break
            else -> throw IllegalArgumentException("Invalid codon")
        }
        result += protein
    }
    return result
}