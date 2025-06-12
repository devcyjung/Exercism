export class GameOfLife {
  constructor(private matrix: number[][]) {
    if (matrix.some(r => r.length !== matrix[0]?.length ?? 0)) {
      throw new Error('Uneven matrix row length')
    }
    matrix.forEach(r => {
      if (r.some(c => c !== 0 && c !== 1)) {
        throw new Error('Invalid number in matrix')
      }
    })
  }

  public tick(): void {
    const rows = this.matrix.length
    const cols = this.matrix[0]?.length ?? 0
    const toOne = []
    const toZero = []
    for (let i = 0; i < rows; ++i) {
      for (let j = 0; j < cols; ++j) {
        switch (this.aliveNeighbors(i, j)) {
          case 2:
            break
          case 3:
            toOne.push([i, j])
            break
          default:
            toZero.push([i, j])
        }
      }
    }
    toOne.forEach(([i, j]) => this.matrix[i][j] = 1)
    toZero.forEach(([i, j]) => this.matrix[i][j] = 0)
  }

  private aliveNeighbors(i: number, j: number): number {
    const rows = this.matrix.length
    const cols = this.matrix[0]?.length ?? 0
    let acc = 0
    for (let di = -1; di <= 1; ++di) {
      for (let dj = -1; dj <= 1; ++dj) {
        if (di === 0 && dj === 0) {
          continue
        }
        if (0 <= i + di && i + di < rows && 0 <= j + dj && j + dj < cols) {
          acc += this.matrix[i + di][j + dj]
        }
      }
    }
    return acc
  }

  public state(): unknown {
    return structuredClone(this.matrix)
  }
}