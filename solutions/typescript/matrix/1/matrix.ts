export class Matrix {
  private readonly matrix: number[][]
  private readonly nCol: number
  
  constructor(input: string) {
    const rows = input.split('\n')
    this.matrix = rows.map(rowString => rowString.split(' ')).map(row => row.map(Number))
    this.nCol = this.matrix[0]?.length ?? 0
  }

  get rows(): number[][] {
    return this.matrix
  }

  get columns(): number[][] {
    return Array.from({length: this.nCol}).map((_, i) => this.matrix.map(row => row[i]))
  }
}
