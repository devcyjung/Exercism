package matrix

import (
    "errors"
    "slices"
    "strconv"
    "strings"
)

type matrix struct {
    values []int
    rowCount, colCount int
}

type Matrix = *matrix

func New(s string) (mat Matrix, err error) {
    var m matrix
	rows := strings.Split(s, "\n")
    var cells []string
    m.rowCount = len(rows)
    rowOffset := 0
    for i, row := range rows {
        cells = strings.Fields(row)
        if i == 0 {
            m.colCount = len(cells)
            m.values = make([]int, m.rowCount * m.colCount)
        }
        if len(cells) != m.colCount {
            err = errors.New("unequal row length")
            return
        }
        for j, cell := range cells {
            m.values[rowOffset+j], err = strconv.Atoi(cell)
            if err != nil {
                return
            }
        }
        rowOffset += m.colCount
    }
    mat = &m
    return
}

func (m Matrix) Cols() (colMat [][]int) {
	colMat = make([][]int, m.colCount)
    for j := 0; j < m.colCount; j++ {
        colMat[j] = make([]int, m.rowCount)
    }
    rowOffset := 0
    for i := 0; i < m.rowCount; i++ {
        for j := 0; j < m.colCount; j++ {
            colMat[j][i] = m.values[rowOffset+j]
        }
        rowOffset += m.colCount
    }
    return
}

func (m Matrix) Rows() (rowMat [][]int) {
	slices.Grow(rowMat, m.rowCount)
    rowOffset := 0
    for i := 0; i < m.rowCount; i++ {
        rowMat = append(rowMat, slices.Clone(m.values[rowOffset : rowOffset+m.colCount]))
        rowOffset += m.colCount
    }
    return
}

func (m Matrix) Set(row, col, val int) bool {
    if 0 <= row && row < m.rowCount && 0 <= col && col < m.colCount {
        m.values[row*m.rowCount+col] = val
        return true
    }
	return false
}

func (m Matrix) Get(row, col int) int {
    return m.values[row*m.rowCount+col]
}
