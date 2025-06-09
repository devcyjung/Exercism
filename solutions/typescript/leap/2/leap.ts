type NonNegativeFiniteInteger<T extends number> =
typeof Infinity extends T ? never :
`${T}` extends `${any}e+${any}` ? T : 
    `${T}` extends `-${any}` | `${any}.${any}` | `${any}e-${any}` ? never : T;

export function isLeap<T extends number>(year:NonNegativeFiniteInteger<T>): boolean {
  return (year%4 === 0 && year%100 !== 0) || year%400===0;
}
