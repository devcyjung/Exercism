export const reverseString = str => Array.from(new Intl.Segmenter().segment(str))
  .map(({segment}) => segment).reverse().join('')