package rectangles

import "slices"

func Count(diagram []string) int {
    rows, cols := make(map[int]int), make(map[int]int)
	for i, row := range diagram {
        for j, ch := range row {
            if ch == '+' {
                rows[i]++
                cols[j]++
            }
        }
    }
    rcand, ccand := make([]int, 0, len(rows)), make([]int, 0, len(cols))
    for r, cnt := range rows {
        if cnt > 1 {
            rcand = append(rcand, r)
        }
    }
    for c, cnt := range cols {
        if cnt > 1 {
            ccand = append(ccand, c)
        }
    }
    slices.Sort(rcand)
    slices.Sort(ccand)
    rectCount := 0
    for r1 := 0; r1 < len(rcand) - 1; r1++ {
        for r2 := r1 + 1; r2 < len(rcand); r2++ {
            for c1 := 0; c1 < len(ccand) - 1; c1++ {
                for c2 := c1 + 1; c2 < len(ccand); c2++ {
                    i1, i2, j1, j2 := rcand[r1], rcand[r2], ccand[c1], ccand[c2]
                    if checkRect(diagram, i1, i2, j1, j2) {
                        rectCount++
                    }
                }
            }
        }
    }
    return rectCount
}

func checkRect(diagram []string, i1, i2, j1, j2 int) bool {
    if j1 >= len(diagram[i1]) || j1 >= len(diagram[i2]) ||
    	j2 >= len(diagram[i1]) || j2 >= len(diagram[i2]) {
        return false
    }
    if diagram[i1][j1] != '+' || diagram[i1][j2] != '+' ||
    	diagram[i2][j1] != '+' || diagram[i2][j2] != '+' {
        return false
    }
    for i := i1 + 1; i < i2; i++ {
        if (diagram[i][j1] != '|' && diagram[i][j1] != '+') ||
        	(diagram[i][j2] != '|' && diagram[i][j2] != '+') {
            return false
        }
    }
	for j := j1 + 1; j < j2; j++ {
        if (diagram[i1][j] != '-' && diagram[i1][j] != '+') ||
        	(diagram[i2][j] != '-' && diagram[i2][j] != '+') {
            return false
        }
    }
    return true
}