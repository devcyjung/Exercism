export function isLeap(year:number):boolean {
  if (Number.isInteger(year) && year >= 0){
    return ((year%4 === 0 && year%100 !== 0) || year%400===0)
  }
  throw new Error("Year must be integer and not negative.");
}
