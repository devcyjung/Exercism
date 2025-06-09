export function getFirstCard([first]) {
  return first
}

export function getSecondCard([_, second]) {
  return second
}

export function swapTopTwoCards([first, second, ...rest]) {
  return [second, first, ...rest]
}

export function discardTopCard([first, ...rest]) {
  return [first, rest]
}

const FACE_CARDS = ['jack', 'queen', 'king'];

export function insertFaceCards([first, ...rest]) {
  return [first, ...FACE_CARDS, ...rest]
}
