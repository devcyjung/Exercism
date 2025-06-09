package stateoftictactoe

import (
    "errors"
    "unicode"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (retS State, retE error) {
    b := make([]rune, 9)
    var temp []rune
    var tempRune rune
    for i := 0; i < 3; i++ {
        temp = []rune(board[i])
        if len(temp) != 3 {
            retE = errors.New("not a 3x3 board")
            return
        }
        for j := 0; j < 3; j++ {
            tempRune = temp[j]
            if !unicode.IsSpace(tempRune) && tempRune != 'X' && tempRune != 'O' {
                retE = errors.New("illegal character in board")
                return
            }
        }
        copy(b[i*3:(i+1)*3], temp)
    }
    retE = checkLegalTurnOrder(b)
    if retE != nil {
        return
    }
    var win bool
    win, retE = isWin(b)
    if retE != nil {
        return
    }
    if win {
        retS = Win
        return
    }
	if isFilled(b) {
        retS = Draw
        return
    }
    retS = Ongoing
    return
}

func isFilled(board []rune) bool {
    for _, c := range board {
        if unicode.IsSpace(c) {
            return false
        }
    }
    return true
}

var winningLines = [][]int{
    {0,1,2},
    {3,4,5},
    {6,7,8},
    {0,3,6},
    {1,4,7},
    {2,5,8},
    {0,4,8},
    {2,4,6},
}

func checkLegalTurnOrder(board []rune) (retE error) {
    var oCount, xCount int
    for _, v := range board {
        switch v {
            case 'O': oCount++
            case 'X': xCount++
        }
    }
    if oCount == xCount + 1 {
        retE = errors.New("player O started")
        return
    }
    if xCount > oCount + 1 {
        retE = errors.New("player X took multiple turns in a row")
        return
    }
    if oCount > xCount + 1 {
        retE = errors.New("player O took multiple turns in a row")
        return
    }
    return
}

func isWin(board []rune) (retB bool, retE error) {
    var saw, cur, winner rune
    var winCount int
    for _, line := range winningLines {
        for i, pos := range line {
            cur = board[pos]
            if i == 0 {
                saw = cur
            }
            if unicode.IsSpace(cur) {
                break
            }
            if saw != cur {
                break
            }
            if i == 2 {
                winCount++
                if winCount > 2 {
                    retE = errors.New("players kept playing after a win")
                    return
                }
                if winCount == 2 {
                    if winner != cur {
                        retE = errors.New("players kept playing after a win")
                    	return
                    }
                }
                winner = cur
            }
        }
    }
    
    if winner == 'O' || winner == 'X' {
        retB = true
        return
    }
    return
}