package strand

import "strings"

type Gene string

func (g Gene) ToReplaced(from string, to string) Gene {
    return Gene(strings.ReplaceAll(string(g), from, to))
}

func (g Gene) ToUpper() Gene {
    return Gene(strings.ToUpper(string(g)))
}

func ToRNA(dna string) string {
	g := Gene(dna).
    	ToReplaced("G", "c").
    	ToReplaced("C", "g").
    	ToReplaced("T", "a").
    	ToReplaced("A", "u").
    	ToUpper()
    return string(g)
}
