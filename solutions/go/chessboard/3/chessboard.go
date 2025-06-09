package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
    count := 0
    for _, ok := range cb[file] {
        if ok {
            count++
        }
    }
    return count
}

func CountInRank(cb Chessboard, rank int) int {
    count := 0
    index := rank - 1
    if index < 0 {
        return count
    }
    for _, file := range cb {
        if index < len(file) && file[index] {
            count++
        }
    }
    return count
}

func CountAll(cb Chessboard) int {
    count := 0
    for _, file := range cb {
        count += len(file)
    }
    return count
}

func CountOccupied(cb Chessboard) int {
    count := 0
    for _, file := range cb {
        for _, ok := range file {
            if ok {
                count++
            }
        }
    }
    return count
}