const _str = digit => digit === 0 ? ' ' : digit.toString()
const _count = rows => ''.concat(...rows).split('').filter(c => c === '*').length
const _cell = (r1, r2, r3, i) => _str(_count([r1, r2, r3].map(v => v.substring(i - 1, i + 2))))
const _row = (r1 = '', r2, r3 = '') => r2.split('').map((c, i) => c === '*' ? '*' : _cell(r1, r2, r3, i))
export const annotate = input => input.map((row, i) => _row(input[i - 1], row, input[i + 1]).join(''))