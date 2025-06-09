export function needsLicense(kind) {
  return kind === 'car' || kind === 'truck'
}

export function chooseVehicle(option1, option2) {
  let chosen
  if (option1 < option2) {
    chosen = option1
  } else {
    chosen = option2
  }
  return `${chosen} is clearly the better choice.`
}

export function calculateResellPrice(originalPrice, age) {
  let modifier
  if (age < 3) {
    modifier = 0.8
  } else if (age > 10) {
    modifier = 0.5
  } else {
    modifier = 0.7
  }
  return modifier * originalPrice
}
