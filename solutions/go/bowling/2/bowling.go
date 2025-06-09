package bowling

import "errors"

var (
    ErrInvalidNumber = errors.New("Given pin count is out of bounds")
    ErrUnfinishedGame = errors.New("Game is not finished yet")
    ErrFinishedGame = errors.New("Game is already over")
)

type Game struct {
    frames		[11][2]int
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
    g.frames[g.index[0]][g.index[1]] = pins
    frameTotal := g.getFrameTotal(g.index[0])
    if frameTotal > 10 && (!g.isLastFrame() || g.isLastFrame() && g.frames[10][0] != 10) {
        g.frames[g.index[0]][g.index[1]] = 0
        return ErrInvalidNumber
    } else if !g.isLastFrame() && (frameTotal == 10 || g.index[1] == 1) {
        g.moveToNextFrame()
    } else {
        g.index[1]++
    }
    if g.index[0] >= 10 {
        switch g.getFrameType(9) {
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
        base := g.getFrameTotal(i)
        switch g.getFrameType(i) {
        case openType:
            acc += base
        case strikeType:
            if g.frames[i + 1][1] == 0 {
                acc += base + g.frames[i + 1][0] + g.frames[i + 2][0]
            } else {
                acc += base + g.getFrameTotal(i + 1)
            }
        case spareType:
            acc += base + g.frames[i + 1][0]
        }
    }
    return acc, nil
}

func (g *Game) isLastFrame() bool {
    return g.index[0] == 10
}

func (g *Game) moveToNextFrame() {
    g.index[0], g.index[1] = g.index[0] + 1, 0
}

func (g *Game) moveToSecondShot() {
    g.index[1] = 1
}

func (g *Game) getFrameType(frameIdx int) frameType {
    if g.frames[frameIdx][0] == 10 {
        return strikeType
    }
    if g.getFrameTotal(frameIdx) == 10 {
        return spareType
    }
    return openType
}

func (g *Game) getFrameTotal(frameIdx int) int {
    return g.frames[frameIdx][0] + g.frames[frameIdx][1]
}