export class Bowling {
  private static NegativeRollError = new Error('Negative roll is invalid')
  private static PinCountOverflowError = new Error('Pin count exceeds pins on the lane')
  private static UnfinishedGameError = new Error('Score cannot be taken until the end of the game')
  private static FinishedGameError = new Error('Cannot roll after game is over')
  
  private done: boolean = false
  private frameIdx: number = 0
  private rollIdx: number = 0
  private lastIdx: number = 9
  private frames: number[][] = Array.from(
    { length: 12 },
    () => Array.from({ length: 2 }, () => 0)
  )
  private bonuses: number[][] = Array.from(
    { length: 12 },
    () => Array.from({ length: 2 }, () => 0)
  )
  private bonusCounts: number[] = Array.from({ length: 12 }, () => 0)
  
  public roll(pins: number): void | never {
    if (this.done) {
      throw Bowling.FinishedGameError
    }
    if (pins < 0) {
      throw Bowling.NegativeRollError
    }
    if (pins > 10) {
      throw Bowling.PinCountOverflowError
    }
    if (this.frameIdx >= 10) {
      const bonusRollSum = this.frames[10][0] + pins
      if (bonusRollSum > 10 && this.frames[10][0] !== 10) {
        throw Bowling.PinCountOverflowError
      }
    }
    const frameSum = this.frames[this.frameIdx][0] + pins
    if (frameSum > 10) {
      throw Bowling.PinCountOverflowError
    }
    this.frames[this.frameIdx][this.rollIdx] = pins
    for (const previous of [this.frameIdx - 1, this.frameIdx - 2]) {
      if (0 <= previous && previous < 10 && this.bonusCounts[previous] > 0) {
        --(this.bonusCounts[previous])
        this.bonuses[previous][this.bonusCounts[previous]] = pins
      }
    }
    if (this.frameIdx >= 10) {
      ++(this.frameIdx)
    } else if (pins === 10) {
      if (this.frameIdx === 9) {
        this.lastIdx += 2
      }
      this.bonusCounts[this.frameIdx] = 2
      ++(this.frameIdx)
    } else if (frameSum === 10) {
      if (this.frameIdx === 9) {
        ++(this.lastIdx)
      }
      this.bonusCounts[this.frameIdx] = 1
      ++(this.frameIdx)
      this.rollIdx = 0
    } else {
      if (this.rollIdx === 1) {
        ++(this.frameIdx)
        this.rollIdx = 0
      } else {
        ++(this.rollIdx)
      }
    }
    if (this.frameIdx === this.lastIdx + 1) {
      this.done = true
    }
  }

  public score(): number | never {
    if (!this.done) {
      throw Bowling.UnfinishedGameError
    }
    let totalScore = 0
    for (let frame = 0; frame < 10; ++frame) {
      for (let roll = 0; roll < 2; ++roll) {
        totalScore += this.frames[frame][roll] + this.bonuses[frame][roll]
      }
    }
    return totalScore
  }
}