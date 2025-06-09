package minesweeper

import (
    "cmp"
    "strings"
    "slices"
)

type pair struct {
	row, col int
}

func Annotate(input []string) []string {
    if len(input) == 0 {
        return input
    }
    nrows := len(input)
    ncols := len([]rune(slices.MaxFunc(input, func(a, b string) int {
        return cmp.Compare(len([]rune(a)), len([]rune(b)))
    })))
    concat := []rune(strings.Join(input, ""))
    if len(concat) != nrows * ncols {
        return input
    }
    get := func(i, j int) rune {
        return concat[i * ncols + j]
    }
    bounds := func(i, j int) bool {
        return i >= 0 && j >= 0 && i < nrows && j < ncols
    }
    count := func(i, j int) rune {
        acc := '0'
        for r := i - 1; r <= i + 1; r++ {
            for c := j - 1; c <= j + 1; c++ {
                if bounds(r, c) && get(r, c) == '*' {
                    acc++
                }
            }
        }
        return acc
    }
    result := make([]string, 0, nrows)
    var b strings.Builder
    for i := 0; i < nrows; i++ {
        b.Reset()
        for j := 0; j < ncols; j++ {
        	switch get(i, j) {
            case '*':
                b.WriteRune('*')
            case ' ':
                cnt := count(i, j)
                switch cnt {
                case '0':
                    b.WriteRune(' ')
                default:
                    b.WriteRune(cnt)
                }
            default:
                return input
            }
        }
        result = append(result, b.String())
    }
    return result
}