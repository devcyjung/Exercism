package protein

import "errors"

var CodonProtein = map[string]string{
    "AUG":	"Methionine",
    "UUU":	"Phenylalanine",
    "UUC": 	"Phenylalanine",
    "UUA": 	"Leucine",
    "UUG":	"Leucine",
    "UCU": 	"Serine",
    "UCC":	"Serine",
    "UCA":	"Serine",
    "UCG":	"Serine",
    "UAU": 	"Tyrosine",
    "UAC": 	"Tyrosine",
    "UGU": 	"Cysteine",
    "UGC": 	"Cysteine",
    "UGG": 	"Tryptophan",
    "UAA": 	"STOP",
    "UAG": 	"STOP",
    "UGA": 	"STOP",
}

var (
    ErrStop error 			= errors.New("Stop")
    ErrInvalidBase error 	= errors.New("Invalid Base")
)

func FromRNA(rna string) (r []string, e error) {
	size := len(rna)
    if size % 3 != 0 {
        e = ErrInvalidBase
        return
    }
    for i := 3; i <= size; i += 3 {
        v, err := FromCodon(rna[i-3:i])
        if err != nil {
            if errors.Is(err, ErrStop) {
                break
            }
            r, e = []string{}, errors.Join(err, errors.ErrUnsupported)
            return
        }
        r = append(r, v)
    }
    return
}

func FromCodon(codon string) (r string, e error) {
	v, ok := CodonProtein[codon]
    if !ok {
        e = ErrInvalidBase
        return
    }
    if v == "STOP" {
        e = ErrStop
        return
    }
    r = v
    return
}
