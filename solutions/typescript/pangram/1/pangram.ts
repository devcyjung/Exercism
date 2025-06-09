const ALPHABETS = Array.from({ length: 26 }, (_, i) =>
  String.fromCharCode('A'.charCodeAt(0) + i)
);

export function isPangram(input: string): boolean {
  const capitalizedInput = input.toUpperCase()
  return ALPHABETS.every(ch => capitalizedInput.includes(ch))
}