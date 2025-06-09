package matrix

import (
	"errors"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Matrix struct {
	values, rowMax, colMin []int
}

type Pair struct {
	row, col int
}

func min(a, b int) int { // can delete in 1.21+
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int { // can delete in 1.21+
	if a > b {
		return a
	} else {
		return b
	}
}

func lineSep(r rune) bool {
  return r == '\n'
}

func New(s string) (retM *Matrix, retE error) {
	var matrix Matrix
	var rows, cells []string
	rows = strings.FieldsFunc(s, lineSep)
	rowCount := len(rows)
	colCount := 0
	temp := 0
	rowOffset := 0
	const MAX = math.MaxInt
	for i, row := range rows {
		cells = strings.Fields(row)
		if i == 0 {
			colCount = len(cells)
			matrix.values = make([]int, rowCount*colCount)
			matrix.rowMax = make([]int, rowCount)
			matrix.colMin = make([]int, colCount)
			for c := range matrix.colMin {
				matrix.colMin[c] = MAX
			}
		}
		if len(cells) != colCount {
			retE = errors.New("unequal length of rows")
			return
		}
		for j, cell := range cells {
			temp, retE = strconv.Atoi(cell)
			if retE != nil {
				return
			}
			matrix.values[rowOffset+j], matrix.rowMax[i], matrix.colMin[j] =
            	temp, max(temp, matrix.rowMax[i]), min(temp, matrix.colMin[j])
		}
		rowOffset += colCount
	}
	retM = &matrix
	return
}

func (m *Matrix) Saddle() (retP []Pair) {
	slices.Grow(retP, len(m.values))
	rowLen := len(m.colMin)
	rowOffset := 0
	var rowIdx1 int
	for i, rm := range m.rowMax {
		rowIdx1 = i + 1
		for j, cm := range m.colMin {
			if rm == cm && rm == m.values[rowOffset+j] {
				retP = append(retP, Pair{rowIdx1, j + 1})
			}
		}
		rowOffset += rowLen
	}
	slices.Clip(retP)
	return
}