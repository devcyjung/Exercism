package matrix

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

var ErrInvalidFormat = errors.New("input has invalid format")

type Matrix = *matrix

type matrix struct {
	rowwise, colwise []int
	nrows, ncols     int
}

func New(s string) (Matrix, error) {
	var result Matrix
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		fields := strings.Fields(line)
		switch i {
		case 0:
			result = &matrix{
				rowwise: make([]int, len(lines)*len(fields)),
				colwise: make([]int, len(lines)*len(fields)),
				nrows:   len(lines),
				ncols:   len(fields),
			}
		default:
			if len(fields) != result.ncols {
				return nil, ErrInvalidFormat
			}
		}
		for j, value := range fields {
			cell, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			result.Set(i, j, cell)
		}
	}
	return result, nil
}

func (m Matrix) Rows() [][]int {
	clone := slices.Clone(m.rowwise)
	result := make([][]int, m.nrows)
	offset := 0
	var nextOffset int
	for i := range result {
		nextOffset = offset + m.ncols
		result[i] = clone[offset:nextOffset]
		offset = nextOffset
	}
	return result
}

func (m Matrix) Cols() [][]int {
	clone := slices.Clone(m.colwise)
	result := make([][]int, m.ncols)
	offset := 0
	var nextOffset int
	for j := range result {
		nextOffset = offset + m.nrows
		result[j] = clone[offset:nextOffset]
		offset = nextOffset
	}
	return result
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= m.nrows || col >= m.ncols {
		return false
	}
	m.row(row)[col] = val
	m.col(col)[row] = val
	return true
}

func (m Matrix) row(i int) []int {
	return m.rowwise[i*m.ncols : (i+1)*m.ncols]
}

func (m Matrix) col(j int) []int {
	return m.colwise[j*m.nrows : (j+1)*m.nrows]
}