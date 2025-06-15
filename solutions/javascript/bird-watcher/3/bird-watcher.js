export function totalBirdCount(birdsPerDay) {
  return Array.prototype.reduce.call(birdsPerDay, (a, b) => a + b, 0)
}

export function birdsInWeek(birdsPerDay, week) {
  return Array.prototype.slice.call(birdsPerDay, 7 * (week-1), 7 * week).reduce((a, b) => a + b, 0)
}

export function fixBirdCountLog(birdsPerDay) {
  Array.prototype.forEach.call(birdsPerDay, (v, i, a) => a[i] = i % 2 === 0 ? v + 1 : v)
  return birdsPerDay
}