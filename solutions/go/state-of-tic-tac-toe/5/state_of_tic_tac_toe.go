package stateoftictactoe

import (
    "errors"
    "slices"
    "strings"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
    Err		State = ""
)

var (
    ErrInvalidBoardSize = errors.New("Board Size is not squared")
    ErrInvalidBoardRune = errors.New("Board has invalid character")
    ErrInvalidMarkCount = errors.New("A player made invalid moves")
    ErrInvalidGameState = errors.New("Game progressed beyond the end")
)

func StateOfTicTacToe(board []string) (State, error) {
	switch {
    case len(board) == 0:
        return Err, ErrInvalidBoardSize
    case len(slices.MaxFunc(board, byLength)) != len(board):
        return Err, ErrInvalidBoardSize
    case len(slices.MinFunc(board, byLength)) != len(board):
        return Err, ErrInvalidBoardSize
    case slices.ContainsFunc(board, invalidString):
        return Err, ErrInvalidBoardRune
    }
    state := newGameState(len(board))
    for i, row := range board {
        for j, cell := range row {
            state.update(cell, i, j)
        }
    }
    return state.tallyState()
}

func byLength(a, b string) int {
    return len(a) - len(b)
}

func invalidString(str string) bool {
    return strings.ContainsFunc(str, invalidRune)
}

func invalidRune(r rune) bool {
    switch r {
    case ' ', 'O', 'X':
        return false
    default:
        return true
    }
}

type gameState struct {
    size				int
	rows, cols, diag	[]rune
    mark				map[rune]int
}

func newGameState(size int) *gameState {
    cellCounter := map[rune]int{
        'O':	0,
        'X':	0,
        ' ':	0,
    }
    return &gameState{
        size:	size,
        rows:	make([]rune, size),
        cols:	make([]rune, size),
        diag:	make([]rune, 2),
        mark:	cellCounter,
    }
}

func (g *gameState) update(ch rune, i, j int) {
    g.mark[ch]++
    switch ch {
    case ' ':
        g.rows[i] = 'D'
        g.cols[j] = 'D'
    default:
        switch g.rows[i] {
        case 0, ch:
            g.rows[i] = ch
        default:
            g.rows[i] = 'D'
        }
    	switch g.cols[j] {
        case 0, ch:
            g.cols[j] = ch
        default:
            g.cols[j] = 'D'
        }
    }
    if i == j {
        switch ch {
        case ' ':
            g.diag[0] = 'D'
        default:
            switch g.diag[0] {
            case 0, ch:
                g.diag[0] = ch
            default:
                g.diag[0] = 'D'
            }
        }
    }
    if i + j == g.size - 1 {
        switch ch {
        case ' ':
            g.diag[1] = 'D'
        default:
            switch g.diag[1] {
            case 0, ch:
                g.diag[1] = ch
            default:
                g.diag[1] = 'D'
            }
        }
    }
}

func (g *gameState) tallyState() (State, error) {
    oMarks, xMarks, blanks := g.mark['O'], g.mark['X'], g.mark[' ']
    if !(oMarks <= xMarks && xMarks <= oMarks + 1) {
        return Err, ErrInvalidMarkCount
    }
    rowWinner := 'D'
    colWinner := 'D'
    var rowWinCount, colWinCount, rowIdx, colIdx int
    for i, rowWin := range g.rows {
        colWin := g.cols[i]
        if rowWin != 'D' {
            rowWinCount++
            rowWinner = rowWin
            rowIdx = i
        }
        if colWin != 'D' {
            colWinCount++
            colWinner = colWin
            colIdx = i
        }
    }
    if rowWinCount > 1 || colWinCount > 1 {
        return Err, ErrInvalidGameState
    }
    winner := 'D'
    for _, win := range []rune{rowWinner, colWinner, g.diag[0], g.diag[1]} {
        switch win {
        case 'D':
            continue
        default:
            if winner != 'D' && winner != win {
                return Err, ErrInvalidGameState
            }
            winner = win
        }
    }
    if winner == 'D' {
        if blanks == 0 {
            return Draw, nil
        }
        return Ongoing, nil
    }
    switch {
    case winner == g.diag[0] && winner == g.diag[1]:
        if winner == rowWinner && (g.size % 2 != 1 || rowIdx != g.size / 2) {
            return Err, ErrInvalidGameState
        }
        if winner == colWinner && (g.size % 2 != 1 || colIdx != g.size / 2) {
            return Err, ErrInvalidGameState
        }
        return Win, nil
    case winner == g.diag[0]:
        if winner == rowWinner && winner == colWinner && rowIdx != colIdx {
            return Err, ErrInvalidGameState
        }
        return Win, nil
    case winner == g.diag[1]:
        if winner == rowWinner && winner == colWinner && rowIdx + colIdx != g.size - 1 {
            return Err, ErrInvalidGameState
        }
        return Win, nil
    default:
        return Win, nil
    }
}