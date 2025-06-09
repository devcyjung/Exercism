export function totalBirdCount(birdsPerDay) {
  return birdsPerDay.reduce((a, b) => a + b, 0)
}

export function birdsInWeek(birdsPerDay, week) {
  return birdsPerDay.slice(7 * (week-1), 7 * week).reduce((a, b) => a + b, 0)
}

export function fixBirdCountLog(birdsPerDay) {
  birdsPerDay.forEach((v, i, a) => a[i] = i % 2 === 0 ? v + 1 : v)
  return birdsPerDay
}