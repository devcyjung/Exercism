export function getFirstCard([first]) {
  return first
}

export function getSecondCard([_, second]) {
  return second
}

export function swapTwoCards([first, second, ...rest]) {
  return [second, first, ...rest]
}

export function shiftThreeCardsAround([first, second, third, ...rest]) {
  return [second, third, first, ...rest]
}

export function pickNamedPile({ chosen }) {
  return chosen
}

export function swapNamedPile({ chosen, disregarded }) {
  return { chosen: disregarded, disregarded: chosen }
}