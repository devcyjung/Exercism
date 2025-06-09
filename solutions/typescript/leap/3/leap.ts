
export function isLeap(year:number): boolean {
  if (Number.isNaN(year) || !Number.isInteger(year) || year < 0 || year === Infinity) {
    throw new Error("Floating point numbers, NaN, Infinity, negative numbers are not valid year numbers.");
  }
  return (year % 4 === 0 && year % 100 !== 0) || year % 400 === 0;
}
