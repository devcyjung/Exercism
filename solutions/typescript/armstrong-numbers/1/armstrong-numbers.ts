export function isArmstrongNumber(num: number | bigint): boolean {
  const digits = Array.from(num.toString(), ch => BigInt(ch))
  const original = BigInt(num)
  const len = BigInt(digits.length)
  const armstrongSum = digits.reduce((acc, cur) => acc + cur ** len, 0n)
  return armstrongSum === original
}