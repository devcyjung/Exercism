const SQUARE_RADIUS_TO_SCORE = [
  [10 * 10, 0],
  [5 * 5, 1],
  [1 * 1, 5],
  [-1, 10],
] as const

function assertUnreachable(...values: unknown[]): never {
  throw new Error("Unreachable code reached with value: " + values);
}

type Score = typeof SQUARE_RADIUS_TO_SCORE[number][1]

export function score(x: number, y: number): Score {
  const radiusSquare = x * x + y * y;
  for (const [threshold, score] of SQUARE_RADIUS_TO_SCORE) {
    if (radiusSquare > threshold) {
      return score
    }
  }
  return assertUnreachable(x, y)
}