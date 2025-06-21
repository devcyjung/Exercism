// @ts-check

/**
 * Generates a random starship registry number.
 *
 * @returns {string} the generated registry number.
 */
export function randomShipRegistryNumber() {
  const randoms = crypto.getRandomValues(new Uint8Array(4))
  return `NCC-${randoms[0] % 9 + 1}${randoms[1] % 10}${randoms[2] % 10}${randoms[3] % 10}`
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
  return PLANETS[crypto.getRandomValues(new Uint8Array(1))[0] % PLANETS.length]
}
