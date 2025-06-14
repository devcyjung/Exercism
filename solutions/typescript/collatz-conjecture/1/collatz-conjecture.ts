export function steps(num: number): number | never {
  if (num < 1 || !Number.isInteger(num)) {
    throw new Error('Only positive integers are allowed')
  }
  let steps = 0
  let cur = num
  while (cur !== 1) {
    ++steps
    if ((cur & 1) === 0) {
      cur >>= 1
    } else {
      cur += (cur << 1) + 1
    }
  }
  return steps
}