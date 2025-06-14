declare global {
  interface String {
    replaceAll(pattern: string | RegExp, replacement: string): string
  }
}

const punctRemover = /[^A-Za-z0-9\-']/g
const quoteRemover = /(\s'|'\s|^'|'$)/g
const splitSpace = /\s/

export function count(input: string): Map<string, number> {
  const wordcount: Map<string, number> = new Map()
  input
    .toLowerCase()
    .replaceAll(punctRemover, ' ')
    .replaceAll(quoteRemover, ' ')
    .split(/\s/)
    .filter(s => s.length > 0)
    .forEach(word => {
      wordcount.set(word, (wordcount.get(word) ?? 0) + 1)
    })
  return wordcount
}