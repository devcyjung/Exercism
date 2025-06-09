const squaredRadiusToScore = [
  [10 * 10, 0],
  [5 * 5, 1],
  [1 * 1, 5],
  [-1, 10],
] as const

export function score(x: number, y: number): 0 | 1 | 5 | 10 {
  const radiusSquare = x * x + y * y;
  for (const [threshold, score] of squaredRadiusToScore) {
    if (radiusSquare > threshold) {
      return score
    }
  }
  return 0
}