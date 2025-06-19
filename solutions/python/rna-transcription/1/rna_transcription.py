def to_rna(dna_strand):
    return dna_strand.translate(_mapping)

_mapping = str.maketrans("GCTA", "CGAU")