package bowling

import "errors"

var (
    ErrInvalidRoll = errors.New("Invalid number of pins in a roll")
    ErrFinished = errors.New("The game is already over")
    ErrUnfinished = errors.New("The game is not over yet")
)

type Game struct {
    done						bool
    frameIdx, rollIdx, lastIdx	int
    frames, bonuses				[12][2]int
    bonusCounts					[10]int
}

func NewGame() *Game {
	return &Game{lastIdx: 9}
}

func (g *Game) Roll(pins int) error {
    if g.done {
        return ErrFinished
    }
    frameSum := g.frames[g.frameIdx][0] + pins
	if pins < 0 || pins > 10 || frameSum > 10 {
        return ErrInvalidRoll
    }
    if g.frameIdx >= 10 {
        bonusRollSum := g.frames[10][0] + pins
        if bonusRollSum > 10 && g.frames[10][0] != 10 {
            return ErrInvalidRoll
        }
    }
    g.frames[g.frameIdx][g.rollIdx] = pins
    for _, previous := range []int{g.frameIdx - 1, g.frameIdx - 2} {
        if 0 <= previous && previous < 10 && g.bonusCounts[previous] > 0 {
            g.bonusCounts[previous]--
            g.bonuses[previous][g.bonusCounts[previous]] = pins
        }
    }
    switch {
    case g.frameIdx >= 10:
        g.frameIdx++
    case pins == 10:
        if g.frameIdx == 9 {
            g.lastIdx += 2
        }
        g.bonusCounts[g.frameIdx] = 2
        g.frameIdx++
    case frameSum == 10:
        if g.frameIdx == 9 {
            g.lastIdx++
        }
        g.bonusCounts[g.frameIdx] = 1
        g.frameIdx++
        g.rollIdx = 0
    default:
        if g.rollIdx == 1 {
            g.frameIdx++
            g.rollIdx = 0
        } else {
            g.rollIdx++
        }
    }
    if g.frameIdx == g.lastIdx + 1 {
        g.done = true
    }
    return nil
}

func (g *Game) Score() (int, error) {
	if !g.done {
        return 0, ErrUnfinished
    }
    totalScore := 0
    for frame := 0; frame < 10; frame++ {
        for roll := 0; roll < 2; roll++ {
            totalScore += g.frames[frame][roll] + g.bonuses[frame][roll]
        }
    }
    return totalScore, nil
}