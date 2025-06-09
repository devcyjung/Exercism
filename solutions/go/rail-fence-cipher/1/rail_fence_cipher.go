package railfence

import (
    "cmp"
    "slices"
    "strings"
    "unicode/utf8"
)

const deadRune rune = 0xDEAD

func Encode(message string, rails int) string {
    array, matrix, _ := allocator(message, rails)
    posFunc := positionManager(rails)
    for _, r := range message {
        i, j := posFunc()
        matrix[i][j] = r
    }
    var b strings.Builder
    for _, r := range array {
        if r != deadRune {
            b.WriteRune(r)
        }
    }
    return b.String()
}

func Decode(message string, rails int) string {
	_, matrix, width := allocator(message, rails)
    posFunc := positionManager(rails)
    positions := make([][2]int, 0, width)
    for range message {
        i, j := posFunc()
        positions = append(positions, [2]int{i, j})
    }
    slices.SortFunc(positions, sortByRowThenCol)
    for idx, r := range message {
        pos := positions[idx]
        matrix[pos[0]][pos[1]] = r
    }
    var b strings.Builder
    newPosFunc := positionManager(rails)
    for range message {
        i, j := newPosFunc()
        b.WriteRune(matrix[i][j])
    }
    return b.String()
}

func sortByRowThenCol(a, b [2]int) int {
    if a[0] != b[0] {
        return cmp.Compare(a[0], b[0])
    }
    return cmp.Compare(a[1], b[1])
}

func positionManager(rails int) func() (int, int) {
    dirs := [2][2]int{{1, 1}, {-1, 1}}
    dirIdx := 0
    pos := [2]int{0, 0}
    dir := dirs[dirIdx]
    return func() (int, int) {
        defer func() {
            pos[0] += dir[0]
            pos[1] += dir[1]
            if pos[0] < 0 || pos[0] >= rails {
                pos[0] -= dir[0]
            	pos[1] -= dir[1]
                dirIdx += 1
                dirIdx %= 2
                dir = dirs[dirIdx]
                pos[0] += dir[0]
            	pos[1] += dir[1]
            }
        }()
        return pos[0], pos[1]
    }
}

func allocator(message string, rails int) ([]rune, [][]rune, int) {
    width := utf8.RuneCountInString(message)
	array := make([]rune, rails * width)
    for i := range array {
        array[i] = deadRune
    }
    matrix := make([][]rune, rails)
    for i := range matrix {
        matrix[i] = array[i * width:(i + 1) * width]
    }
    return array, matrix, width
}