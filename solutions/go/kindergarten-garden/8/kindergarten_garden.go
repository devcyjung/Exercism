package kindergarten

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type (
	child      = string
	plant      = string
	plantArray = [4]plant
)

type Garden struct {
	table map[child]plantArray
}

const diagramFmt = "\n[VRCG]{%[1]d}\n[VRCG]{%[1]d}"

var (
    ErrInvalidFormat = errors.New("Invalid diagram format")
    ErrDuplicateNames = errors.New("duplicate children names")
)

func NewGarden(diagram string, children []string) (*Garden, error) {
	pattern := fmt.Sprintf(diagramFmt, 2*len(children))
	r, err := regexp.Compile(pattern)
	if err != nil || !r.MatchString(diagram) {
		return nil, ErrInvalidFormat
	}
	childrenClone := slices.Clone(children)
	slices.Sort(childrenClone)
	if len(childrenClone) != len(slices.Compact(childrenClone)) {
		return nil, ErrDuplicateNames
	}
	table := make(map[child]plantArray)
	sepIdx := strings.LastIndex(diagram, "\n")
	r1 := diagram[1:sepIdx]
	r2 := diagram[sepIdx+1:]
	for i, c := range childrenClone {
		table[c] = plantArray{p(r1[2*i]), p(r1[2*i+1]), p(r2[2*i]), p(r2[2*i+1])}
	}
	return &Garden{table}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.table[child]
	if !ok {
		return nil, false
	}
	return slices.Clone(plants[:]), true
}

func p(initial byte) string {
	switch initial {
	case 'V':
		return "violets"
	case 'R':
		return "radishes"
	case 'C':
		return "clover"
	case 'G':
		return "grass"
	}
	return ""
}