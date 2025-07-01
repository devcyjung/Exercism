export const toRna = dna => Array.from(dna).map(ch => {
  switch(ch) {
    case 'G':
      return 'C'
    case 'C':
      return 'G'
    case 'T':
      return 'A'
    case 'A':
      return 'U'
  }
}).join('')