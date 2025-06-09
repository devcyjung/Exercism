package ocr

import "strings"

var (
    charMapping = []byte{
        ' ', '_', ' ',
        '|', '_', '|',
        '|', '_', '|',
        ' ', ' ', ' ',
    }
    digitPositions = []map[int]struct{}{
        {1: {}, 3: {}, 5: {}, 6: {}, 7: {}, 8: {}},
        {5: {}, 8: {}},
        {1: {}, 4: {}, 5: {}, 6: {}, 7: {}},
        {1: {}, 4: {}, 5: {}, 7: {}, 8: {}},
        {3: {}, 4: {}, 5: {}, 8: {}},
        {1: {}, 3: {}, 4: {}, 7: {}, 8: {}},
        {1: {}, 3: {}, 4: {}, 6: {}, 7: {}, 8: {}},
        {1: {}, 5: {}, 8: {}},
        {1: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}},
        {1: {}, 3: {}, 4: {}, 5: {}, 7: {}, 8: {}},
    }
    offsetMapping = [][2]int{
        {0, 0}, {0, 1}, {0, 2},
        {1, 0}, {1, 1}, {1, 2},
        {2, 0}, {2, 1}, {2, 2},
        {3, 0}, {3, 1}, {3, 2},
    }
)

func recognizeDigit(matrix []string, basei, basej int) rune {
    for digit, positions := range digitPositions {
        matchFound := false
        for i, offset := range offsetMapping {
            data := matrix[basei + offset[0]][basej + offset[1]]
            if _, ok := positions[i]; ok {
                if charMapping[i] != data {
                    matchFound = false
                    break
                }
                matchFound = true
            } else {
                if ' ' != data {
                    matchFound = false
                    break
                }
                matchFound = true
            }
        }
        if matchFound {
            return '0' + rune(digit)
        }
    }
    return '?'
}

func Recognize(input string) []string {
    matrix := strings.Split(strings.Trim(input, "\n"), "\n")
	nrows := len(matrix)
    ncols := len(matrix[0])
    nOutRows, nOutCols := nrows / 4, ncols / 3
    result := make([]string, nOutRows)
    var b strings.Builder
    for i := 0; i < nOutRows; i++ {
        b.Reset()
        for j := 0; j < nOutCols; j++ {
            basei, basej := 4 * i, 3 * j
            b.WriteRune(recognizeDigit(matrix, basei, basej))
        }
        result[i] = b.String()
    }
    return result
}