type UpperChar = 'A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' |
                 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z'
type LowerChar = Lowercase<UpperChar>

export function transform(old: Record<number, UpperChar[]>): Partial<Record<LowerChar, number>> {
  return Object.fromEntries(
    Object.entries(old).flatMap(
      ([digit, chars]) => chars.map(upper => [upper.toLowerCase(), Number(digit)])
    )
  )
}