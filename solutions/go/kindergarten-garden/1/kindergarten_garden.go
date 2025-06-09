package kindergarten

import (
    "fmt"
    "sort"
    "strings"
)

type Garden struct {
    sortedChildren	[]string
    plants			[]strings.Builder
}

type duplicateChildrenNameError struct {
    children	[]string
    duplicate	string
}

func (e duplicateChildrenNameError) Error() string {
    quoted := make([]string, len(e.children))
    for i, s := range e.children {
        quoted[i] = fmt.Sprintf("%q", s)
    }
    return fmt.Sprintf(`duplicate children name:
in the given list of children name: [ %s ]
found duplicate name: %q`, strings.Join(quoted, ", "), e.duplicate)
}

// Errors
// When diagram does not follow "\n[CGRV]{len(children)}\n[CGRV]{len(children)}" patern.
func NewGarden(diagram string, children []string) (g *Garden, err error) {
    population := len(children)
    plants := make([]strings.Builder, population)
    err = checkDiagram(diagram, population, plants)
    if err != nil {
        return
    }
    sortedChildren := make([]string, population)
    copy(sortedChildren, children)
    sort.Strings(sortedChildren)
	for i := 0; i < len(sortedChildren) - 1; i++ {
        if sortedChildren[i] == sortedChildren[i + 1] {
            err = duplicateChildrenNameError{children, sortedChildren[i]}
            return
        }
    }
    g = &Garden {
        sortedChildren,
        plants,
    }
    return
}

// ok = false when child is not found in the children list.
func (g *Garden) Plants(child string) (result []string, ok bool) {
	index := sort.SearchStrings(g.sortedChildren, child)
    if index == len(g.sortedChildren) || g.sortedChildren[index] != child {
        return
    }
    result, ok = make([]string, 4), true
    var name string
    for i, p := range (&g.plants[index]).String() {
        name, _ = lookupPlantName(p)
        result[i] = name
    }
    return
}

type invalidFormatError struct {
    input	 string
    position int
    expected []rune
    found	 rune
}

func (e invalidFormatError) Error() string {
    quoted := make([]string, len(e.expected))
    for i, s := range e.expected {
        quoted[i] = fmt.Sprintf("%q", s)
    }
    return fmt.Sprintf(`invalid format:
in the given input: %q,
at rune index:%d,
expected: one of [ %s ],
but found: %q`, e.input, e.position, strings.Join(quoted, ", or "), e.found)
}

type invalidLengthError struct {
    input			 string
    population, expected, length int
}

func (e invalidLengthError) Error() string {
    return fmt.Sprintf(`invalid length:
in the given input: %q,
with student count: %d,
the expected input length was: %d
but the actual length is: %d`, e.input, e.population, e.expected, e.length)
}

// Check that diagram follows "\n[CGRV]{len(children)}\n[CGRV]{len(children)}" patern
func checkDiagram(diagram string, population int, plants []strings.Builder) (err error) {
    if len(diagram) != 2 + 4 * population {
        err = invalidLengthError{diagram, population, 2 + 4 * population, len(diagram)}
    	return
    }
    var ok bool
    var plantIndex int
    for i, ch := range diagram {
        switch i {
        case 0: // the first '\n'
            fallthrough
        case population * 2 + 1:
        	if ch != '\n' {
                err = invalidFormatError{diagram, i, []rune{'\n'}, ch}
                return
            }
        default: // first row is [1, population*2], and second is [population*2+2, 4*population+1]
            _, ok = lookupPlantName(ch)
            if !ok {
                err = invalidFormatError{diagram, i, []rune{'G', 'C', 'R', 'V'}, ch}
                return
            }
            if i <= 2 * population {
                plantIndex = (i - 1) / 2
            } else {
                plantIndex = (i - 2 * population - 2) / 2
            }
            (&plants[plantIndex]).WriteRune(ch)
        }
    }
    return
}

func lookupPlantName(alphabet rune) (result string, ok bool) {
    switch alphabet {
    case 'G':
        return "grass", true
    case 'C':
        return "clover", true
    case 'R':
        return "radishes", true
    case 'V':
        return "violets", true
    default:
        return "", false
    }
}