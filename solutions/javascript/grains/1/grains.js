export const square = sq => {
  const big = BigInt(sq)
  if (big < 1n || 64n < big) {
    throw new Error("square must be between 1 and 64")
  }
  return 1n << (big - 1n)
}

export const total = () => (1n << 64n) - 1n