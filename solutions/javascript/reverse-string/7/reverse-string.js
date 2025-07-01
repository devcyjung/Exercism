const SEGMENTER = new Intl.Segmenter()
export const reverseString = str => Array.from(SEGMENTER.segment(str))
  .map(({segment}) => segment).reverse().join('')