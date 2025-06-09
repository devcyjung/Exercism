package matrix

import (
    "strings"
    "strconv"
    "fmt"
    "math"
)

type Matrix	[][]int
type Pair	[2]int

func New(s string) (retM *Matrix, retE error) {
    var mat Matrix
    retM = &mat
    if len(s) == 0 {
        return
    }
    for i, line := range strings.Split(s, "\n") {
        var row []int
        for j, value := range strings.Split(line, " ") {
            intValue, err := strconv.Atoi(value)
            if err != nil {
                retM, retE = nil, fmt.Errorf("String parse error at %d %d %v", i, j, err)
                return
            }
           	row = append(row, intValue)
        }
        mat = append(mat, row)
    }
    return
}

func (m *Matrix) Saddle() (retP []Pair) {
    var maxEachRow []int
    var minEachCol []int
    for i, row := range *m {
        for j, val := range row {
            if len(maxEachRow) <= i {
                maxEachRow = append(maxEachRow, 0)
            }
            if len(minEachCol) <= j {
                minEachCol = append(minEachCol, math.MaxInt)
            }
            if maxEachRow[i] < val {
                maxEachRow[i] = val
            }
            if minEachCol[j] > val {
                minEachCol[j] = val
            }
        }
    }
    for i, row := range *m {
        for j, val := range row {
            if maxEachRow[i] == val && minEachCol[j] == val {
                retP = append(retP, Pair{ i+1, j+1 })
            }
        }
    }
    return
}

