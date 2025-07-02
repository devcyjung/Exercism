const _str = digit => digit === 0 ? ' ' : digit.toString()
const _count = (rows) => ''.concat(...rows).split('').filter(c => c === '*').length
const _cell = (pre, cur, next, i) => _str(_count([pre, cur, next].map(v => v?.substring(i - 1, i + 2) ?? '')))
const _row = (pre, cur, next) => cur.split('').map((c, i) => c === '*' ? '*' : _cell(pre, cur, next, i))
export const annotate = input => input.map((row, i) => _row(input[i - 1], row, input[i + 1]).join(''))