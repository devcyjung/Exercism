package minesweeper

import "strings"

func Annotate(input []string) []string {
	if len(input) == 0 || len(input[0]) == 0 {
		return input
	}

	nRows, nCols := len(input), len(input[0])

	rawBytes := []byte(strings.Join(input, ""))

	board := make([][]byte, nRows)
	for i := 0; i < nRows; i++ {
		board[i] = rawBytes[i*nCols : (i+1)*nCols]
	}

	for i, row := range board {
		for j := range row {
			if board[i][j] == '*' {
				updateAdjacentCells(board, i, j)
			}
		}
	}

	for i := 0; i < nRows; i++ {
		input[i] = string(board[i])
	}

	return input
}

func updateAdjacentCells(board [][]byte, row, col int) {
	type pair struct {
		row, col int
	}

	nRow, nCol := len(board), len(board[0])
	isValidPosition := func(p *pair) bool {
		return p.row >= 0 && p.row < nRow && p.col >= 0 && p.col < nCol
	}

	deltas := []pair{
		{-1, -1}, {-1, 0}, {-1, 1}, // above

		{0, -1}, {0, 1}, // sides

		{1, -1}, {1, 0}, {1, 1}, // below
	}

	var nextPair *pair
	var nextCell *byte
	for _, delta := range deltas {
		nextPair = &pair{row + delta.row, col + delta.col}

		if isValidPosition(nextPair) {
			nextCell = &board[nextPair.row][nextPair.col]
			switch *nextCell {
			case '*':
			case ' ':
				*nextCell = '1'
			default:
				*nextCell++
			}
		}
	}
}
