const COLORS = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
] as const;

export type Color = typeof COLORS[number];

export const decodedValue = (colors:Color[]) => {
  if (colors.length < 2) {
    throw new Error("At least two colors are required");
  }
  const firstColorValue = COLORS.indexOf(colors[0]);
  const secondColorValue = COLORS.indexOf(colors[1]);

  return firstColorValue * 10 + secondColorValue;
};