export function find(haystack: number[], needle: number): number | never {
  haystack.sort((a, b) => a - b)
  let left = 0
  let right = haystack.length
  while (left < right) {
    const mid = Math.floor((left + right - 1) / 2)
    if (haystack[mid] === needle) {
      return mid
    }
    if (haystack[mid] < needle) {
      left = mid + 1
    } else {
      right = mid
    }
  }
  throw new Error('Value not in array')
}