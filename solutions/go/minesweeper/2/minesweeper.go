package minesweeper

import "strings"

type pair struct {
	row, col int
}

const (
	ONE byte = iota + '1'
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	MINE     byte = '*'
	UNTAPPED byte = ' '
	NONE     byte = ' '
)

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

	isValidPosition := func(p *pair) bool {
		return p.row >= 0 && p.row < nRows && p.col >= 0 && p.col < nCols
	}

	getMark := func(p *pair) *byte {
		return &board[p.row][p.col]
	}

	mergePair := func(p1, p2 *pair) *pair {
		return &pair{p1.row + p2.row, p1.col + p2.col}
	}

	updateCell := func(p *pair) {
		if !isValidPosition(p) {
			return
		}

		mark := getMark(p)

		if *mark != UNTAPPED {
			return
		}

		deltas := []*pair{
			{-1, -1}, {-1, 0}, {-1, 1}, // above

			{0, -1}, {0, 1}, // sides

			{1, -1}, {1, 0}, {1, 1}, // below
		}

		var next *pair
		for _, delta := range deltas {
			next = mergePair(p, delta)
			if !isValidPosition(next) {
				continue
			}
			if MINE == *getMark(next) {
				switch *mark {
				case UNTAPPED:
					*mark = ONE
				case MINE:
					panic("current pair has been concurrently modified")
				// case NONE:
				// 	*mark = ONE
				default:
					*mark++
				}
			}
		}
		if *mark == UNTAPPED {
			*mark = NONE
		}
	}

	for i, row := range board {
		for j := range row {
			updateCell(&pair{i, j})
		}
	}

	for i := 0; i < nRows; i++ {
		input[i] = string(board[i])
	}

	return input
}
