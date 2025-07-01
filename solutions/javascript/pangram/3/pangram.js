const PANGRAM = 0x3FFFFFF
const charIsLower = ch => ch >= 'a' && ch <= 'z'
const charIdx = ch => ch.charCodeAt(0) - 'a'.charCodeAt(0)
export const isPangram = input => PANGRAM === Array.from(input.toLowerCase())
  .reduce((bits, ch) => !charIsLower(ch) ? bits : bits | 1 << charIdx(ch), 0)