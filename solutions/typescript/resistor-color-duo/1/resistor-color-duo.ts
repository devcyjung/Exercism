const colorMap:string[] = [
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
];

export function decodedValue(prop : string[]) : number {
  const indexArray:number[] = [];
  for (let i=0; i<prop.length && i<2; ++i) {
    indexArray.push(colorMap.indexOf(prop[i]));
    if (indexArray[i] === -1 || Number.isNaN(indexArray[i]) || indexArray[i] === Infinity || indexArray[i] === -Infinity)
    {
      throw new Error(`Unsupported color at index: ${i}`)
    }
  }
  if (indexArray.length !== 2) {
    throw new Error("Need at least 2 colors");
  }
  return indexArray.reduce((acc:number, cur:number) => {
    return acc * 10 + cur;
  }, 0)
}
