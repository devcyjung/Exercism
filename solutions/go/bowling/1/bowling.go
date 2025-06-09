package bowling

import "errors"

var (
    ErrInvalidNumber = errors.New("Given pin count is out of bounds")
    ErrUnfinishedGame = errors.New("Game is not finished yet")
    ErrFinishedGame = errors.New("Game is already over")
)

type Game struct {
    frames		[11][2]*int
    index		[2]int
    done		bool
}

type frameType int

const (
    openType	frameType = iota
    spareType
    strikeType
)

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Roll(pins int) error {
	if g.done {
        return ErrFinishedGame
    }
    if pins < 0 || pins > 10 {
        return ErrInvalidNumber
    }
    g.frames[g.index[0]][g.index[1]] = &pins
    frameTotal := getFrameTotal(g.frames[g.index[0]])
    if frameTotal > 10 && (!g.isLastFrame() || g.isLastFrame() && *(g.frames[10][0]) != 10) {
        g.frames[g.index[0]][g.index[1]] = nil
        return ErrInvalidNumber
    } else if !g.isLastFrame() && (frameTotal == 10 || g.index[1] == 1) {
        g.index[0], g.index[1] = g.index[0] + 1, 0
    } else {
        g.index[1]++
    }
    if g.index[0] >= 10 {
        switch getFrameType(g.frames[9]) {
        case openType:
            g.done = true
        case strikeType:
            if g.isLastFrame() && g.index[1] == 2 {
                g.done = true
            }
        case spareType:
            if g.index[1] == 1 {
            	g.done = true   
            }
        }
    }
    return nil
}

func (g *Game) Score() (int, error) {
	if !g.done {
        return 0, ErrUnfinishedGame
    }
    acc := 0
    for i := 0; i < 10; i++ {
        base := getFrameTotal(g.frames[i])
        switch getFrameType(g.frames[i]) {
        case openType:
            acc += base
        case strikeType:
            if g.frames[i + 1][1] == nil {
                acc += base + *(g.frames[i + 1][0]) + *(g.frames[i + 2][0])
            } else {
                acc += base + getFrameTotal(g.frames[i + 1])
            }
        case spareType:
            acc += base + *(g.frames[i + 1][0])
        }
    }
    return acc, nil
}

func (g *Game) isLastFrame() bool {
    return g.index[0] == 10
}

func getFrameType(frame [2]*int) frameType {
    if *(frame[0]) == 10 {
        return strikeType
    }
    if getFrameTotal(frame) == 10 {
        return spareType
    }
    return openType
}

func getFrameTotal(frame [2]*int) int {
    if frame[0] == nil {
        return 0
    }
    if frame[1] == nil {
        return *(frame[0])
    }
    return *(frame[0]) + *(frame[1])
}