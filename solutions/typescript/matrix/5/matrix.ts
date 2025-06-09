export class Matrix {
  private readonly rowMat: number[][]
  private readonly colMat: number[][]
  
  constructor(input: string) {
    let maxRowLen = 0
    this.rowMat = input.split('\n')
      .map(rowStr => rowStr.split(' '))
      .map(row => {
        maxRowLen = Math.max(row.length, maxRowLen)
        return row.map(Number)
      })
    this.colMat = Array.from({length: maxRowLen})
      .map((_, i) => this.rowMat.map(row => row[i] ?? NaN))
    Object.freeze(this.rowMat)
    Object.freeze(this.colMat)
  }

  get rows(): number[][] {
    return this.rowMat
  }

  get columns(): number[][] {
    return this.colMat
  }
}
