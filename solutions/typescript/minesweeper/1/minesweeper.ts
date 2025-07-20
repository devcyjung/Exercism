export function annotate(field: string[]): string[] {
  const nrow = field.length
  const ncol = field[0]?.length ?? 0
  for (const row of field) {
    if (row.length !== ncol) {
      throw new Error("Uneven fields")
    }
  }
  const count = (i: number, j: number): number => {
    let mines = 0
    for (let x = Math.max(0, i - 1); x <= Math.min(i + 1, nrow - 1); ++x) {
      for (let y = Math.max(0, j - 1); y <= Math.min(j + 1, ncol - 1); ++y) {
        if (field[x][y] === '*') ++mines
      }
    }
    return mines
  }
  const buf = Array.from({length: ncol})
  for (let i = 0; i < nrow; ++i) {
    for (let j = 0; j < ncol; ++j) {
      if (field[i][j] === '*') {
        buf[j] = '*'
        continue
      }
      const mines = count(i, j)
      buf[j] = mines === 0 ? ' ' : String(mines)
    }
    field[i] = buf.join('')
  }
  return field
}
