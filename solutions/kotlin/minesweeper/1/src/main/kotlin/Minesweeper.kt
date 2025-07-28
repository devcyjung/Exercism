import kotlin.math.max
import kotlin.math.min

data class MinesweeperBoard(val board: List<String>) {
    val annotated: List<String> = board.let { board ->
        val nrows = board.size
        val ncols = if (nrows == 0) { 0 } else { board[0].length }
        require(board.all { row -> row.length == ncols })
        val isMine: (Int, Int) -> Boolean = { rowIdx, colIdx -> board[rowIdx][colIdx] == '*' }
        val adjacentMines: (Int, Int) -> Int = { rowIdx, colIdx ->
            (max(0, rowIdx - 1)..min(rowIdx + 1, nrows - 1)).flatMap { adjRowIdx ->
                (max(0, colIdx - 1)..min(colIdx + 1, ncols - 1)).map { adjColIdx -> 
                    adjRowIdx to adjColIdx
                }
            }.filter { (adjRowIdx, adjColIdx) -> isMine(adjRowIdx, adjColIdx) }.count()
        }
        board.mapIndexed { rowIdx, row ->
            row.asSequence().mapIndexed { colIdx, cell ->
                if (cell == '*') { '*' } else {
                    adjacentMines(rowIdx, colIdx).let { adjMineCount -> 
                        if (adjMineCount == 0) { ' ' } else { adjMineCount.digitToChar() }
                    }
                }
            }.joinToString(separator = "")
        }
    }

    fun withNumbers(): List<String> = annotated
}