type Position = readonly [number, number]

type Positions = {
  white: Position
  black: Position
}

export class QueenAttack {
  public readonly black: Position
  public readonly white: Position

  constructor({ white = [7, 3], black = [0, 3] }: Partial<Positions> = {}) {
    QueenAttack.validatePosition(white)
    QueenAttack.validatePosition(black)
    QueenAttack.validateDuplicate(white, black)
    this.white = white
    this.black = black
  }

  toString(): string {
    const board = Array.from({length: 8}, () => Array.from({length: 8}).fill('_'))
    board[this.white[0]][this.white[1]] = 'W'
    board[this.black[0]][this.black[1]] = 'B'
    return board.map(row => row.join(' ')).join('\n')
  }

  get canAttack(): boolean {
    const rowDiff = this.white[0] - this.black[0]
    const colDiff = this.white[1] - this.black[1]
    return rowDiff * colDiff * (rowDiff + colDiff) * (rowDiff - colDiff) === 0
  }

  static validatePosition(position: Position): void {
    for (const coord of position) {
      if (coord < 0 || coord >= 8) {
        throw new Error('Queen must be placed on the board')
      }
    }
  }

  static validateDuplicate(position1: Position, position2: Position): void {
    if (position1[0] === position2[0] && position1[1] === position2[1]) {
      throw new Error('Queens cannot share the same space')
    }
  }
}