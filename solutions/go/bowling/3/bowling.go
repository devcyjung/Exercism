package bowling

import "errors"

var (
    ErrInvalidRoll = errors.New("Invalid number of pins in a roll")
    ErrFinished = errors.New("The game is already over")
    ErrUnfinished = errors.New("The game is not over yet")
)

type Game struct {
    done				bool
    frameIdx, rollIdx	int
    frames, bonuses		[12][2]int
    lastIdx				int
    bonusCounts			[10]int
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
    if g.frameIdx >= 10 { // BONUS ROUNDS
        bonusRollSum := g.frames[10][0] + pins
        if bonusRollSum > 10 && g.frames[10][0] != 10 {
            return ErrInvalidRoll
        }
        g.updateScores(pins)
        g.frameIdx++
    } else if pins == 10 { // STRIKE
        if g.frameIdx == 9 {
            g.lastIdx += 2
        }
        g.updateScores(pins)
        g.bonusCounts[g.frameIdx] = 2
        g.frameIdx++
    } else if frameSum == 10 { // SPARE
        if g.frameIdx == 9 {
            g.lastIdx++
        }
        g.updateScores(pins)
        g.bonusCounts[g.frameIdx] = 1
        g.frameIdx++
        g.rollIdx = 0
    } else { // OPEN
        g.updateScores(pins)
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
    score := 0
    for i := 0; i < 10; i++ {
        for j := 0; j < 2; j++ {
            score += g.frames[i][j] + g.bonuses[i][j]
        }
    }
    return score, nil
}

func (g *Game) updateScores(pins int) {
    g.frames[g.frameIdx][g.rollIdx] = pins
    for _, previous := range []int{g.frameIdx - 1, g.frameIdx - 2} {
        if 0 <= previous && previous < 10 && g.bonusCounts[previous] > 0 {
            g.bonusCounts[previous]--
            g.bonuses[previous][g.bonusCounts[previous]] = pins
        }
    }
}