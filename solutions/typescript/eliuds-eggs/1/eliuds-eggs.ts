export const eggCount = (displayValue: number): number => {
  return Array.from(Math.floor(displayValue).toString(2)).filter(ch => ch === '1').length
}