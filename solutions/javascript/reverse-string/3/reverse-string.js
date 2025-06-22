//
// This is only a SKELETON file for the 'Reverse String' exercise. It's been provided as a
// convenience to get you started writing code faster.


export const reverseString = input => Array.from(new Intl.Segmenter().segment(input))
  .map(({ segment }) => segment).reverse().join('')
