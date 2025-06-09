export function twoSum(array1, array2) {
  return Number(array1.reduce((a, b) => a + b, "")) + Number(array2.reduce((a, b) => a + b, ""))
}

export function luckyNumber(value) {
  const original = String(value).split("")
  const reversed = original.toReversed()
  return original.every((v, i) => reversed.at(i) === v)
}

export function errorMessage(input) {
  if (!input) {
    return 'Required field'
  }
  const num = Number(input)
  if (Number.isNaN(num) || num === 0) {
    return 'Must be a number besides 0'
  }
  return ''
}
