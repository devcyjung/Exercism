export const rows = n => {
  if (n <= 0)
    return []
  if (n === 1)
    return [[1]]
  const prev = rows(n - 1)
  const last = prev.at(-1)
  prev.push([1, ...last.slice(1).map((v, i) => v + last[i]), 1])
  return prev
}