package protein

import (
    "errors"
    "slices"
)

var (
    ErrStop error 			= errors.New("Stop")
    ErrInvalidBase error 	= errors.New("Invalid Base")
)

func FromRNA(rna string) ([]string, error) {
    runes := []rune(rna)
    size := len(runes)
    proteins := make([]string, 0, size / 3)
    var protein string
    var err error
	for i := 0; i < size; i += 3 {
        if i + 3 > size {
            return nil, ErrInvalidBase
        }
        protein, err = FromCodon(rna[i:i+3])
        if err == ErrStop {
            break
        }
        if err != nil {
            return nil, err
        }
        proteins = append(proteins, protein)
    }
    slices.Clip(proteins)
    return proteins, nil
}

func FromCodon(codon string) (string, error) {
    switch codon {
    case "AUG":
        return "Methionine", nil
    case "UUU":
        fallthrough
    case "UUC":
        return "Phenylalanine", nil
    case "UUA":
        fallthrough
    case "UUG":
        return "Leucine", nil
    case "UCU":
        fallthrough
    case "UCC":
        fallthrough
    case "UCA":
        fallthrough
    case "UCG":
        return "Serine", nil
    case "UAU":
        fallthrough
    case "UAC":
        return "Tyrosine", nil
    case "UGU":
        fallthrough
    case "UGC":
        return "Cysteine", nil
    case "UGG":
        return "Tryptophan", nil
    case "UAA":
        fallthrough
    case "UAG":
        fallthrough
    case "UGA":
        return "", ErrStop
    default:
        return "", ErrInvalidBase
    }
}