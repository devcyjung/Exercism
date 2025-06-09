const isEqual = compareTo => value => Object.is(compareTo, value)
const isOdd = e => e & 1
const toggle = flip => value => flip ^ isOdd(value)

export function cardTypeCheck(stack, card) {
  return stack.filter(isEqual(card)).length
}

export function determineOddEvenCards(stack, countEven) {
  return stack.filter(toggle(countEven)).length
}