package matrix

import (
    "errors"
    "slices"
    "strconv"
    "strings"
)

var ErrInvalidFormat = errors.New("Input has invalid format")

type Matrix = *matrix

type matrix struct {
    rowwise, colwise []int
    nrows, ncols int
}

func New(s string) (Matrix, error) {
	lines := strings.Split(s, "\n")
    var result Matrix
    var err error
    var fields []string
    var cell int
    for i, line := range lines {
        fields = strings.Fields(line)
        if i == 0 {
            result = &matrix{
                rowwise: make([]int, len(lines) * len(fields)),
                colwise: make([]int, len(lines) * len(fields)),
                nrows: len(lines),
                ncols: len(fields),
            }
        }
        if len(fields) != result.ncols {
            return nil, ErrInvalidFormat
        }
        for j, value := range fields {
            cell, err = strconv.Atoi(value)
            if err != nil {
                return nil, err
            }
            result.rowwise[i * result.ncols + j] = cell
            result.colwise[j * result.nrows + i] = cell
        }
    }
    return result, nil
}

func (m Matrix) Cols() [][]int {
	cloned := make([][]int, m.ncols)
    for j := range cloned {
        cloned[j] = slices.Clone(m.colwise[j * m.nrows : (j + 1) * m.nrows])
    }
    return cloned
}

func (m Matrix) Rows() [][]int {
	cloned := make([][]int, m.nrows)
    for i := range cloned {
        cloned[i] = slices.Clone(m.rowwise[i * m.ncols : (i + 1) * m.ncols])
    }
    return cloned
}

func (m Matrix) Set(row, col, val int) bool {
    if row < 0 || col < 0 || row >= m.nrows || col >= m.ncols {
        return false
    }
	m.rowwise[row * m.ncols + col] = val
    m.colwise[col * m.nrows + row] = val
    return true
}