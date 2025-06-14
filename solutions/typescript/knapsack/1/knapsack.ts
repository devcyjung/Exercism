type Item = {
  weight: number
  value: number
}

export function maximumValue(
  { maximumWeight, items }: { maximumWeight: number; items: Item[] }
): number {
  const memo: number[][] = Array.from(
    { length: items.length + 1 }, () => Array.from({ length: maximumWeight + 1 }, () => 0)
  )
  for (let i = 1; i <= items.length; ++i) {
    const currentItem = items[i - 1]
    for (let j = 1; j <= maximumWeight; ++j) {
      const withoutCurrent = memo[i - 1][j] 
      const withCurrent = memo[i - 1][j - currentItem.weight] + currentItem.value
      if (Number.isNaN(withCurrent) || withoutCurrent > withCurrent) {
        memo[i][j] = withoutCurrent
      } else {
        memo[i][j] = withCurrent
      }
    }
  }
  return memo[items.length][maximumWeight]
}