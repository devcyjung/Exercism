export const transform = old => {
  const result = {}
  for (const key in old) {
    const num = Number(key)
    for (const letter of old[key]) {
      result[letter.toLowerCase()] = num
    }
  }
  return result
}