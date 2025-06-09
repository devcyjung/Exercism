package strand

import "strings"

type Gene string

func (g Gene) toReplaced(from string, to string) Gene {
    return Gene(strings.ReplaceAll(string(g), from, to))
}

func (g Gene) toUpper() string {
    return strings.ToUpper(string(g))
}

func ToRNA(dna string) string {
	return Gene(dna).
    	toReplaced("G", "c").
    	toReplaced("C", "g").
    	toReplaced("T", "a").
    	toReplaced("A", "u").
    	toUpper()
}
