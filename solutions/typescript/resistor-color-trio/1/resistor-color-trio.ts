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

function decodedValue(prop:string[]):number {
  const indexArray:number[] = [];
  for (let i=0; i<prop.length && i<2; ++i) {
    indexArray.push(colorMap.indexOf(prop[i]));
    if (indexArray[i] === -1 || Number.isNaN(indexArray[i]) || indexArray[i] === Infinity || indexArray[i] === -Infinity) {
      throw new Error(`Unsupported color at index: ${i}`);
    }
  }
  if (indexArray.length !== 2) {
    throw new Error("Need at least 2 colors");
  }
  return indexArray.reduce((acc:number, cur:number) => {
    return acc * 10 + cur;
  }, 0)
}


export function decodedResistorValue(prop:string[]):string {
  const zeroes:number = colorMap.indexOf(prop[2]!);
  if (zeroes === -1 || Number.isNaN(zeroes) || zeroes === Infinity || zeroes === -Infinity) {
    throw new Error(`Unsupported color at index: 3`);
  }
  const ohmString:string = String(decodedValue(prop.slice(0, 2))).concat('0'.repeat(zeroes));
  let unit:string = 'ohm';
  if (ohmString.length > 3) unit = 'kiloohm';
  if (ohmString.length > 6) unit = 'megaohm';
  if (ohmString.length > 9) unit = 'gigaohm';
  if (!(ohmString[0] === '1' && (ohmString.length === 1 || ohmString[1] === '0'))) {
    unit += 's';
  }
  let baseNum:string = "";
  if (ohmString.length > 1 && ohmString[1] !== '0' && (ohmString.length === 4 || ohmString.length === 7 || ohmString.length === 10)){
    baseNum = ohmString[0].concat('.').concat(ohmString[1]);
  } else {
    let baseNumLen:number = ohmString.length % 3;
    if (baseNumLen === 0) {
      baseNumLen = 3;
    }
    baseNum = ohmString[0] + ((ohmString[1] && ohmString[1] !== "0")?ohmString[1]:"");
    if (baseNumLen > baseNum.length) {
      baseNum += "0".repeat(baseNumLen - baseNum.length);
    }
  }
  return baseNum + " " + unit;
}
