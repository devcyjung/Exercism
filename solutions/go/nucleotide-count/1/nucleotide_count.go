package dna

import "errors"

type Histogram map[rune]uint

type DNA []rune

func (d DNA) Counts() (h Histogram, e error) {
	h = make(Histogram)
    for _, ch := range "ACGT" {
        h[ch] = 0
    }
    for _, ch := range d {
        switch ch {
            case 'A':
            	h['A']++
            case 'C':
            	h['C']++
            case 'G':
            	h['G']++
            case 'T':
            	h['T']++
            default:
            	h, e = make(Histogram), errors.ErrUnsupported
            	return
        }
    }
    return
}
