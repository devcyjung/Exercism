package matrix

import (
    "fmt"
    "strconv"
    "strings"
)

type matrix struct {
    rowwise, colwise [][]int
    rowarr, colarr []int
    nRow, nCol int
}

type Matrix interface {
    Cols() [][]int
    Rows() [][]int
    Set(row, col, val int) bool
}

func New(s string) (m Matrix, e error) {
    var nrow, ncol, val int
    var lines, line []string

    var rowarr, colarr []int
    
    lines = strings.Split(s, "\n")
    nrow = len(lines)
    for i := 0; i < nrow; i++ {
        line = strings.Fields(lines[i])
        if i == 0 {
            ncol = len(line)
            rowarr, colarr = make([]int, nrow*ncol), make([]int, nrow*ncol)
        }
        if len(line) != ncol {
            e = fmt.Errorf(
`error creating Matrix: uneven length of rows
input: "%v"
current line: %+v
expected row length: %v
current line row length: %v`,
                s, line, ncol, len(line))
            return
        }
        for j := 0; j < ncol; j++ {
            val, e = strconv.Atoi(line[j])
            if e != nil {
                return
            }
            rowarr[i*ncol + j] = val
            colarr[j*nrow + i] = val
        }
    }
    
    var matrix matrix
    matrix.rowwise = make([][]int, nrow)
    matrix.colwise = make([][]int, ncol)

    for i := 0; i < nrow; i++ {
        matrix.rowwise[i] = rowarr[i*ncol : (i+1)*ncol]
    }

    for j := 0; j < ncol; j++ {
        matrix.colwise[j] = colarr[j*nrow : (j+1)*nrow]
    }

    matrix.nRow, matrix.nCol = nrow, ncol
	matrix.rowarr, matrix.colarr = rowarr, colarr
    
    m = matrix
    
    return
}

func (m matrix) Cols() (c [][]int) {
    c = make([][]int, m.nCol)
    ca := make([]int, m.nCol*m.nRow)
    copy(ca, m.colarr)
    for j := 0; j < m.nCol; j++ {
        c[j] = ca[j * m.nRow : (j+1) * m.nRow]
    }
	return
}

func (m matrix) Rows() (r [][]int) {
    r = make([][]int, m.nRow)
    ra := make([]int, m.nCol*m.nRow)
    copy(ra, m.rowarr)
    for i := 0; i < m.nRow; i++ {
        r[i] = ra[i * m.nCol : (i+1) * m.nCol]
    }
	return
}

func (m matrix) Set(row, col, val int) bool {
	if !(0 <= row && row < m.nRow) || !(0 <= col && col < m.nCol) {
        return false
    }
    m.rowwise[row][col], m.colwise[col][row] = val, val
    return true
}
