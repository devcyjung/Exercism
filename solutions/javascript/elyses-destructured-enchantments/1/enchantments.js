export function getFirstCard(deck) {
  const [first] = deck
  return first
}

export function getSecondCard(deck) {
  const [_, second] = deck
  return second
}

export function swapTopTwoCards(deck) {
  const [first, second, ...rest] = deck
  return [second, first, ...rest]
}

export function discardTopCard(deck) {
  const [first, ...rest] = deck
  return [first, rest]
}

const FACE_CARDS = ['jack', 'queen', 'king'];

export function insertFaceCards(deck) {
  const [first, ...rest] = deck
  return [first, ...FACE_CARDS, ...rest]
}
