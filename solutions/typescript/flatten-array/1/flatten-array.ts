export function flatten(input: unknown[]): unknown[] {
  return input.flat(Infinity).filter(e => e !== null && e !== undefined)
}