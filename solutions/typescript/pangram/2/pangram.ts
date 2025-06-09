const ALPHABETS = [
  'Z', 'J', 'Q', 'X', 'K', 'V', 'B', 'P', 'G', 'W', 'Y', 'F', 'M',
  'C', 'U', 'L', 'D', 'H', 'R', 'S', 'N', 'I', 'O', 'A', 'T', 'E',
];

export function isPangram(input: string): boolean {
  const capitalizedInput = input.toUpperCase()
  return ALPHABETS.every(ch => capitalizedInput.includes(ch))
}