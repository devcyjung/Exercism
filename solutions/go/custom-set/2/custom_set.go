package stringset

import (
    "fmt"
    "maps"
    "strings"
)

type Set = GenericSet[string]

type GenericSet[E comparable] map[E]bool

const (
    setBeginFmt = '{'
    setEntryFmt = `"%v"`
    setSepFmt = ", "
    setEndFmt = '}'
)

func New() Set {
	return make(GenericSet[string])
}

func NewWithType[E comparable]() GenericSet[E] {
    return make(GenericSet[E])
}

func NewFromSlice(list []string) Set {
    set := make(GenericSet[string])
	for _, value := range list {
        set[value] = true
    }
    return set
}

func NewFromSliceWithType[S ~[]E, E comparable](list S) GenericSet[E] {
    set := make(GenericSet[E])
	for _, value := range list {
        set[value] = true
    }
    return set
}

func (s GenericSet[E]) String() string {
	var b strings.Builder
    b.WriteRune(setBeginFmt)
    for entry := range s {
        if b.Len() > 1 {
            b.WriteString(setSepFmt)
        }
        b.WriteString(fmt.Sprintf(setEntryFmt, entry))
    }
    b.WriteRune(setEndFmt)
    return b.String()
}

func (s GenericSet[E]) IsEmpty() bool {
	return len(s) == 0
}

func (s GenericSet[E]) Has(elem E) bool {
	_, ok := s[elem]
    return ok
}

func (s GenericSet[E]) Add(elem E) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
    return s1.Subset(s2)
}

func (s GenericSet[E]) Subset (other GenericSet[E]) bool {
    if len(s) > len(other) {
        return false
    }
    ok := true
	for entry := range s {
        _, ok = other[entry]
        if !ok {
            break
        }
    }
    return ok
}

func Disjoint(s1, s2 Set) bool {
    return s1.Disjoint(s2)
}

func (s GenericSet[E]) Disjoint(other GenericSet[E]) bool {
	ok := false
	for entry := range s {
        _, ok = other[entry]
        if ok {
            break
        }
    }
    return !ok
}

func Equal(s1, s2 Set) bool {
    return s1.Equal(s2)
}

func (s GenericSet[E]) Equal(other GenericSet[E]) bool {
    return maps.Equal(s, other)
}

func Intersection(s1, s2 Set) Set {
    return s1.Intersection(s2)
}

func (s GenericSet[E]) Intersection(other GenericSet[E]) GenericSet[E] {
	set := make(GenericSet[E])
    var ok bool
    for entry := range s {
        _, ok = other[entry]
        if ok {
            set[entry] = true
        }
    }
    return set
}

func Difference(s1, s2 Set) Set {
    return s1.Difference(s2)
}

func (s GenericSet[E]) Difference(other GenericSet[E]) GenericSet[E] {
	set := make(GenericSet[E])
    var ok bool
    for entry := range s {
        _, ok = other[entry]
        if !ok {
            set[entry] = true
        }
    }
    return set
}

func Union(s1, s2 Set) Set {
    return s1.Union(s2)
}

func (s GenericSet[E]) Union(other GenericSet[E]) GenericSet[E] {
	set := maps.Clone(s)
    for entry := range other {
        set[entry] = true
    }
    return set
}