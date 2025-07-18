const PANGRAM = 0b11_1111_1111_1111_1111_1111_1111
const charIsLower = ch => ch >= 'a' && ch <= 'z'
const charIdx = ch => ch.charCodeAt(0) - 'a'.charCodeAt(0)
export const isPangram = input => PANGRAM === Array.from(input.toLowerCase())
  .reduce((bits, ch) => charIsLower(ch) ? (bits | 1 << charIdx(ch)) : bits, 0)