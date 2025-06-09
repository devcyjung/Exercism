export class GameOfLife {
  #matrix
  #nrow
  #ncol
  
  constructor(matrix) {
    this.#matrix = matrix
    this.#nrow = matrix.length
    this.#ncol = matrix[0]?.length ?? 0
  }

  tick() {
    const aliveNeighbors = [...Array(this.#nrow).keys()]
      .flatMap(
        i => [...Array(this.#ncol).keys()]
                .map(j => [i, j, this.liveNeighbors(i, j)])
      )
    const toZero = aliveNeighbors.filter(([i, j, n]) => n !== 2 && n !== 3)
    const toOne = aliveNeighbors.filter(([i, j, n]) => n === 3)
    this.#matrix = toOne.reduce((mat, [i, j, n]) => {
      mat[i][j] = 1
      return mat
    }, toZero.reduce((mat, [i, j, n]) => {
      mat[i][j] = 0
      return mat
    }, this.#matrix))
  }

  liveNeighbors(i, j) {
    return [
      -1,0,1
    ].flatMap(
      di => [-1,0,1].map(dj => [i + di, j + dj])
    ).filter(
      ([x, y]) => (x !== i || y !== j)
        && (0 <= x && x < this.#nrow) && (0 <= y && y < this.#ncol)
        && !!this.#matrix[x][y]
    ).length
  }

  state() {
    return this.#matrix
  }
}
