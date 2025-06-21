// @ts-check

/**
 * Generates a random starship registry number.
 *
 * @returns {string} the generated registry number.
 */
export function randomShipRegistryNumber() {
  const buffer = new Uint8Array(4)
  const bias9 = Uint8Bias(9)
  const bias10 = Uint8Bias(10)
  do {
    crypto.getRandomValues(buffer.subarray(0, 1))
  } while (buffer[0] >= bias9)
  do {
    crypto.getRandomValues(buffer.subarray(1))
  } while (buffer.subarray(1).some(e => e >= bias10))
  return `NCC-${buffer.map((e, i) => i === 0 ? e % 9 + 1 : e % 10).join('')}`
}

/**
 * Generates a random stardate.
 *
 * @returns {number} a stardate between 41000 (inclusive) and 42000 (exclusive).
 */
export function randomStardate() {
  const random = crypto.getRandomValues(new Uint32Array(2))
  const mantissa = ((random[0] / 2 ** 6) * 2 ** 26) + (random[1] / 2 ** 6)
  return 41000 + 1000 * mantissa / (2 ** 52)
}

/**
 * Generates a random planet class.
 *
 * @returns {string} a one-letter planet class.
 */
export function randomPlanetClass() {
  const PLANETS = ['D', 'H', 'J', 'K', 'L', 'M', 'N', 'R', 'T', 'Y']
  const buffer = new Uint8Array(1)
  const bias = Uint8Bias(PLANETS.length)
  do {
    crypto.getRandomValues(buffer)
  } while (buffer[0] >= bias)
  return PLANETS[buffer[0] % PLANETS.length]
}

function Uint8Bias(range) {
  return 2 ** 8 - 2 ** 8 % range
}