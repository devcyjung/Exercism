package school

import (
    "cmp"
    "slices"
)

type Grade struct{
    grade	int
    names	[]string
}

type School struct{
    school	map[int][]string
}

func New() *School {
	return &School{make(map[int][]string)}
}

func (s *School) Add(student string, grade int) {
	if _, ok := s.school[grade]; !ok {
        s.school[grade] = make([]string, 0, 32)
    }
    s.school[grade] = append(s.school[grade], student)
}

func (s *School) Grade(level int) []string {
	if _, ok := s.school[level]; !ok {
        return []string{}
    }
    slices.Sort(s.school[level])
    result := make([]string, len(s.school[level]))
    for i := range result {
        result[i] = s.school[level][i]
    }
    return result
}

func (s *School) Enrollment() []Grade {
    result := make([]Grade, 0, 16)
	for grade, names := range s.school {
        slices.Sort(names)
        newGrade := Grade{grade, names}
        result = append(result, newGrade)
    }
    slices.SortFunc(result, sortByGrade)
    return result
}

func sortByGrade(g1, g2 Grade) int {
    return cmp.Compare(g1.grade, g2.grade)
}