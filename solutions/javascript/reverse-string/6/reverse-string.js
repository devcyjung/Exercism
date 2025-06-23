//
// This is only a SKELETON file for the 'Reverse String' exercise. It's been provided as a
// convenience to get you started writing code faster.

const segmenter = new Intl.Segmenter()

export const reverseString = input => Array.from(segmenter.segment(input))
  .map(({ segment }) => segment).reverse().join('')
