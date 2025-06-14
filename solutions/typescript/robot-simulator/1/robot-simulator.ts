export class InvalidInputError extends Error {
  constructor(message?: string) {
    super(message || 'Invalid Input')
  }
}

const DIRECTIONS = ['north', 'east', 'south', 'west'] as const
type Direction = typeof DIRECTIONS[number]
type Coordinates = [number, number]

export class Robot {
  private x: number = 0
  private y: number = 0
  private dirIdx: number = 0
  
  get bearing(): Direction {
    return DIRECTIONS[this.dirIdx]
  }

  get coordinates(): Coordinates {
    return [this.x, this.y]
  }

  place({ x, y, direction }: { x: number; y: number; direction: Direction }) {
    const dirIdx = DIRECTIONS.indexOf(direction)
    if (dirIdx < 0) { 
      throw new InvalidInputError()
    }
    this.x = x
    this.y = y
    this.dirIdx = dirIdx
  }

  evaluate(instructions: string) {
    Array.from(instructions).forEach(i => {
      switch(i) {
        case 'R':
          this.dirIdx = (this.dirIdx + 1) % 4
          break
        case 'L':
          this.dirIdx = (this.dirIdx + 3) % 4
          break
        case 'A':
          if ((this.dirIdx & 1) === 0) {
            if (((this.dirIdx >> 1) & 1) === 0) {
              ++(this.y)
            } else {
              --(this.y)
            }
          } else {
            if (((this.dirIdx >> 1) & 1) === 0) {
              ++(this.x)
            } else {
              --(this.x)
            }
          }
          break
        default:
          throw new InvalidInputError()
      }
    })
  }
}
