package queenattack

import "fmt"

func CanQueenAttack(whitePosition, blackPosition string) (res bool, e error) {
	wX, wY, err1 := getXY(whitePosition)
    bX, bY, err2 := getXY(blackPosition)
    if err1 != nil || err2 != nil {
        e = fmt.Errorf("Possible parsing errors for white: (%w) for black: (%w)", err1, err2)
        return
    }
    ok1, ok2 := checkXY(wX, wY), checkXY(bX, bY)
    if !ok1 || !ok2 {
        e = fmt.Errorf("Position out of bounds W: (%d, %d) B: (%d, %d)", wX, wY, bX, bY)
        return
    }
    if wX == bX && wY == bY {
        e = fmt.Errorf("Positions overlap W: (%d, %d) B: (%d, %d)", wX, wY, bX, bY)
        return
    }
    res = wX == bX || wY == bY || wX + wY == bX + bY || wX - wY == bX - bY
    return
}

func getXY(pos string) (x, y rune, e error) {
    r := []rune(pos)
    if len(r) < 2 {
        e = fmt.Errorf("Invalid format: %s", pos)
        return
    }
	x, y = r[0]-'a', r[1]-'1'
    return
}

func checkXY(x, y rune) bool {
    return 0 <= x && x < 8 && 0 <= y && y < 8
}