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

function wedgesPerLime(size) {
  switch (size) {
  case "small":
    return 6
  case "medium":
    return 8
  case "large":
    return 10
  }
}

export function limesToCut(wedgesNeeded, limes) {
  let totalWedges = 0
  let index = 0
  for (const lime of limes) {
    if (totalWedges >= wedgesNeeded) {
      break
    }
    totalWedges += wedgesPerLime(lime)
    ++index
  }
  return index
}

export function remainingOrders(timeLeft, orders) {
  let juicePrepTime = 0
  let index = 0
  while (juicePrepTime < timeLeft && index < orders.length) {
    juicePrepTime += timeToMixJuice(orders[index])
    ++index
  }
  return orders.slice(index)
}