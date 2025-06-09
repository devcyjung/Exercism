package kindergarten

import (
    "errors"
    "fmt"
    "regexp"
    "slices"
    "strings"
)

type child = string
type plant = string
type plantArray = [4]plant

type Garden struct {
    table map[child]plantArray
}

func p(initial byte) string {
    switch (initial) {
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

func NewGarden(diagram string, children []string) (g *Garden, e error) {
	pattern := fmt.Sprintf("\n[VRCG]{%[1]d}\n[VRCG]{%[1]d}", 2 * len(children))
    r, err := regexp.Compile(pattern)
    if err != nil || !r.MatchString(diagram) {
        e = errors.New("invalid diagram format")
        return
    }
    childrenClone := slices.Clone(children)
    slices.Sort(childrenClone)
    if len(childrenClone) != len(slices.Compact(childrenClone)) {
        e = errors.New("duplicate children names")
        return
    }
    table := make(map[child]plantArray)
    sepIdx := strings.LastIndex(diagram, "\n")
    r1 := diagram[1:sepIdx]
    r2 := diagram[sepIdx + 1:]
    for i, c := range childrenClone {
        table[c] = plantArray{p(r1[2 * i]), p(r1[2 * i + 1]), p(r2[2 * i]), p(r2[2 * i + 1])}
    }
    g = &Garden{table}
    return
}

func (g *Garden) Plants(child string) (plants []string, found bool) {
	arr, ok := g.table[child]
    if !ok {
        return
    }
    plants = make([]plant, 4)
    copy(plants, arr[:])
    found = true
    return
}