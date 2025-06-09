package stringset

import (
    "fmt"
    "strings"
)

type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	set := New()
    for _, k := range l {
        set.Add(k)
    }
    return set
}

func (s Set) String() string {
    var builder strings.Builder
    builder.WriteRune('{')
    var w int
	for k := range s {
        if w > 0 {
            builder.WriteString(", ")
        }
        w++
        builder.WriteString(fmt.Sprintf("\"%s\"", k))
    }
    builder.WriteRune('}')
    return builder.String()
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	return s[elem]
}

func (s Set) Add(elem string) {
	s[elem] = true
    return
}

func Subset(s1, s2 Set) bool {
	return Equal(Union(s1, s2), s2)
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
    if len(s1) != len(s2) {
        return false
    }
    for k := range s1 {
        if !s2.Has(k) {
            return false
        }
    }
	return true
}

func Intersection(s1, s2 Set) Set {
	set := New()
    for k := range s1 {
        if s2.Has(k) {
            set.Add(k)
        }
    }
    return set
}

func Difference(s1, s2 Set) Set {
	set := New()
    for k := range s1 {
        if !s2.Has(k) {
            set.Add(k)
        }
    }
    return set
}

func Union(s1, s2 Set) Set {
	set := New()
    for k := range s1 {
        set.Add(k)
    }
    for k := range s2 {
        set.Add(k)
    }
    return set
}
