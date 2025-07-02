const $ = Object.freeze({
  MINE: '*',
  NIL: '',
  ZERO: ' ',
  str: digit => digit === 0 ? $.ZERO : digit.toString(),
  count: rows => $.NIL.concat(...rows).split($.NIL).filter(c => c === $.MINE).length,
  cell: (r1, r2, r3, i) => $.str($.count([r1, r2, r3].map(v => v.substring(i - 1, i + 2)))),
  row: (r1 = $.NIL, r2, r3 = $.NIL) => r2.split($.NIL).map((c, i) => c === $.MINE ? c : $.cell(r1, r2, r3, i))
})
export const annotate = input => input.map((row, i) => $.row(input[i - 1], row, input[i + 1]).join($.NIL))