const isOdd = e => 1 & e
const isEven = e => 1 & ~e

export function getCardPosition(stack, card) {
  return stack.indexOf(card)
}

export function doesStackIncludeCard(stack, card) {
  return stack.includes(card)
}

export function isEachCardEven(stack) {
  return stack.every(isEven)
}

export function doesStackIncludeOddCard(stack) {
  return stack.some(isOdd)
}

export function getFirstOddCard(stack) {
  return stack.find(isOdd)
}

export function getFirstEvenCardPosition(stack) {
  return stack.findIndex(isEven)
}
