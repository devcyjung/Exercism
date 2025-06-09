export function timeToMixJuice(name) {
  switch (name) {
    case 'Pure Strawberry Joy':
      return 0.5
    case 'Energizer':
      return 1.5
    case 'Green Garden':
      return 1.5
    case 'Tropical Island':
      return 3
    case 'All or Nothing':
      return 5
    default:
      return 2.5
  }
}

const reduceFn = (acc, cur) => {
  if (acc[0] > 0) {
    acc[0] -= cur
    ++acc[1]
  }
  return acc
}

export function limesToCut(wedgesNeeded, limes) {
  return limes.map(v => {
    switch (v) {
      case 'small':
        return 6
      case 'medium':
        return 8
      case 'large':
        return 10
    }
  }).reduce(reduceFn, [wedgesNeeded, 0])[1]
}

export function remainingOrders(timeLeft, orders) {
  return orders.slice(
    orders
    .map(timeToMixJuice)
    .reduce(reduceFn, [timeLeft, 0])[1]
  )
}