package kindergarten

import (
	"errors"
	"slices"
)

type (
	child      = string
	plant      = string
	plantArray = [4]plant
)

type Garden struct {
	table map[child]plantArray
}

var (
	ErrInvalidFormat  = errors.New("invalid diagram format")
	ErrDuplicateNames = errors.New("duplicate children names")
)

func NewGarden(diagram string, children []string) (*Garden, error) {
	rowLength := (len(children) << 1) + 1
	if len(diagram) != (rowLength << 1) {
		return nil, ErrInvalidFormat
	}
	for i, char := range diagram {
		switch i {
		case 0, rowLength:
			if char != '\n' {
				return nil, ErrInvalidFormat
			}
		default:
			switch char {
			case 'V', 'R', 'C', 'G':
			default:
				return nil, ErrInvalidFormat
			}
		}
	}
	table := make(map[child]plantArray)
	sorted := slices.Clone(children)
	slices.Sort(sorted)
	lastName := ""
	for i, c := range sorted {
		if c == lastName {
			return nil, ErrDuplicateNames
		}
		table[c] = plantArray{p(diagram[(i<<1)+1]), p(diagram[(i<<1)+2]), p(diagram[(i<<1)+rowLength+1]), p(diagram[(i<<1)+rowLength+2])}
		lastName = c
	}
	return &Garden{table}, nil
}

func (g *Garden) Plants(child child) ([]plant, bool) {
	plants, ok := g.table[child]
	if !ok {
		return nil, false
	}
	return slices.Clone(plants[:]), true
}

func p(initial byte) plant {
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
