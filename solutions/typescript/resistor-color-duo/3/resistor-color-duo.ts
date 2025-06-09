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

export const decodedValue = (colors:Color[]) => COLORS.indexOf(colors[0]!) * 10 + COLORS.indexOf(colors[1]!);
