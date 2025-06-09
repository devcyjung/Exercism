package spiralmatrix

func SpiralMatrix(size int) [][]int {
    if size <= 0 {
        return [][]int{}
    }
    array := make([]int, size * size)
    matrix := make([][]int, size)
    for i := range matrix {
        matrix[i] = array[i * size:(i + 1) * size]
    }
    dirFunc := directionManager(size)
    boundsFunc := boundsManager(size)
    rPos, cPos := 0, 0
    rDir, cDir := dirFunc()
    minR, maxR, minC, maxC := boundsFunc()
    for i := 1; i <= size * size; i++ {
        matrix[rPos][cPos] = i
        rPos += rDir
        cPos += cDir
        if (minR <= rPos && rPos <= maxR && minC <= cPos && cPos <= maxC) {
            continue
        }
        rPos -= rDir
        cPos -= cDir
        rDir, cDir = dirFunc()
        rPos += rDir
        cPos += cDir
        minR, maxR, minC, maxC = boundsFunc()
    }
    return matrix
}

func directionManager(size int) func() (int, int) {
    dirIdx := 0
    directions := [][]int {
        {0, 1}, {1, 0}, {0, -1}, {-1, 0},
    }
    return func() (int, int) {
        defer func() {
           dirIdx = (dirIdx + 1) % 4 
        }()
        dir := directions[dirIdx]
        return dir[0], dir[1]
    }
}

func boundsManager(size int) func() (int, int, int, int) {
    minR, maxR, minC, maxC := 0, size - 1, 0, size - 1
    deltaIdx := 0
    deltas := [][]int {
        {1, 0, 0, 0}, {0, 0, 0, -1}, {0, -1, 0, 0}, {0, 0, 1, 0}, 
    }
    return func() (int, int, int, int) {
        defer func() {
            delta := deltas[deltaIdx]
            minR += delta[0]
            maxR += delta[1]
            minC += delta[2]
            maxC += delta[3]
            deltaIdx = (deltaIdx + 1) % 4
        }()
        return minR, maxR, minC, maxC
    }
}