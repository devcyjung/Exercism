const PLANET_YEAR = {
  "mercury" : 0.2408467,
  "venus": 0.61519726,
  "earth": 1.0,
  "mars": 1.8808158,
  "jupiter": 11.862615,
  "saturn": 29.447498,
  "uranus": 84.016846,
  "neptune": 164.79132
} as const;

const EARTH_YEAR_IN_SECONDS = 31557600;

export type Planet = keyof typeof PLANET_YEAR;

function age(planet: Planet, seconds: number): number {
  const planet_year_in_seconds = EARTH_YEAR_IN_SECONDS * PLANET_YEAR[planet];
  return Math.round(seconds / planet_year_in_seconds * 100) / 100;
}

export {age, EARTH_YEAR_IN_SECONDS, PLANET_YEAR};