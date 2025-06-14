export const square = (sq: number): bigint | never => {
  if (sq < 1 || 64 < sq) {
    throw new Error('Square number out of valid range')
  }
  return 1n << (BigInt(Math.floor(sq)) - 1n)
}

export const total = (): bigint => {
  return (1n << 64n) - 1n
}