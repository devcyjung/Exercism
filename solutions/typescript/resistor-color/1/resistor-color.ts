export const COLORS: string[] = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white'
];

export const colorCode: (color:string) => number|Error = (color: string) => {
  const i: number = COLORS.indexOf(color);
  return (i > -1) ? i : new Error("Not supported Color");
}