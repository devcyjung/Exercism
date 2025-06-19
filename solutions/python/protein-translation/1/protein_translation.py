from itertools import takewhile

def proteins(strand):
    return [
        amino for amino in takewhile(lambda x: x, (
            _codon_to_amino(strand[i:i+3]) for i in range(0, len(strand), 3)
        ))
    ]

def _codon_to_amino(codon):
    match codon:
        case 'AUG':
            return 'Methionine'
        case 'UUU' | 'UUC':
            return 'Phenylalanine'
        case 'UUA' | 'UUG':
            return 'Leucine'
        case 'UCU' | 'UCC' | 'UCA' | 'UCG':
            return 'Serine'
        case 'UAU' | 'UAC':
            return 'Tyrosine'
        case 'UGU' | 'UGC':
            return 'Cysteine'
        case 'UGG':
            return 'Tryptophan'
        case 'UAA' | 'UAG' | 'UGA':
            return None
        case _:
            raise ValueError(f'Invalid codon: {codon}')