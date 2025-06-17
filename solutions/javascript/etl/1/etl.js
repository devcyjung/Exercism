export const transform = old => Object.fromEntries(Object.entries(old).flatMap(
  ([key, values]) => values.map(value => [value.toLowerCase(), Number(key)])))