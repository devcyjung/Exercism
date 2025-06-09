const charMapping = [
  ' ', '_', ' ',
  '|', '_', '|',
  '|', '_', '|',
  ' ', ' ', ' ',
]

const digitPositions = [
  new Set([1, 3, 5, 6, 7, 8]),
  new Set([5, 8]),
  new Set([1, 4, 5, 6, 7]),
  new Set([1, 4, 5, 7, 8]),
  new Set([3, 4, 5, 8]),
  new Set([1, 3, 4, 7, 8]),
  new Set([1, 3, 4, 6, 7, 8]),
  new Set([1, 5, 8]),
  new Set([1, 3, 4, 5, 6, 7, 8]),
  new Set([1, 3, 4, 5, 7, 8]),
]

const offsetMapping = [
  [0, 0], [0, 1], [0, 2],
  [1, 0], [1, 1], [1, 2],
  [2, 0], [2, 1], [2, 2],
  [3, 0], [3, 1], [3, 2],
]

function recognizeDigit(matrix, basei, basej) {
  for (const [digit, positions] of digitPositions.entries()) {
    let matchFound = false
    for (const [i, offset] of offsetMapping.entries()) {
      const data = matrix[basei + offset[0]][basej + offset[1]]
      if (positions.has(i)) {
        if (charMapping[i] !== data) {
          matchFound = false
          break
        }
        matchFound = true
      } else {
        if (data !== ' ') {
          matchFound = false
          break
        }
        matchFound = true
      }
    }
    if (matchFound === true) {
      return String.fromCodePoint('0'.codePointAt(0) + digit) 
    }
  }
  return '?'
}

export function convert (input) {
  const matrix = input.split('\n')
  const [nrows, ncols] = [matrix.length, matrix[0].length]
  const [nOutRows, nOutCols] = [Math.floor(nrows / 4), Math.floor(ncols / 3)]
  const result = Array.from({length: nOutRows})
  const buffer = []
  for (let i = 0; i < nOutRows; ++i) {
    buffer.length = 0
    for (let j = 0; j < nOutCols; ++j) {
      buffer.push(recognizeDigit(matrix, 4 * i, 3 * j))
    }
    result[i] = buffer.join('')
  }
  return result.join(',')
}