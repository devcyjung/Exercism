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

export type MetricUnitPrefix = ''|'kilo'|'mega'|'giga';
export type PluralitySuffix = ''|'s';
export type ResistanceUnit = `${MetricUnitPrefix}ohm${PluralitySuffix}`;
export type ResistanceValue = `${number} ${ResistanceUnit}`;

export function decodedResistorValue(colors:Color[]):ResistanceValue {
  const num0:number = COLORS.indexOf(colors[0]!);
  const num1:number = COLORS.indexOf(colors[1]!);
  const num2:number = COLORS.indexOf(colors[2]!);
  
  let lenNumber:number = 0;
  if (num0 == 0 && num1 == 0) lenNumber = 1;
  else if (num0 == 0) lenNumber = num2 + 1;
  else lenNumber = num2 + 2;
  
  let metricUnitPrefix:MetricUnitPrefix = '';
  switch (Math.floor((lenNumber - 1)/3)) {
    case 0: metricUnitPrefix = ''; break;
    case 1: metricUnitPrefix = 'kilo'; break;
    case 2: metricUnitPrefix = 'mega'; break;
    case 3: metricUnitPrefix = 'giga'; break;
    default: metricUnitPrefix = 'giga';
  }
  // number: 230,000,000 -> baseNumber: 230, number 2,300,000 -> baseNumber: 2.3
  let baseNumber:number = num0 * 10 + num1;
  let x:number;
  let y:number;
  while (
    (
      x = (lenNumber%3===0) ? 3 : lenNumber%3
    ) !== (
      y = Math.floor(baseNumber).toString().length
    )
      ){
    if (y > x) {
      baseNumber /= 10;
    } else {
      baseNumber *= 10;
    }
  }
  const pluralitySuffix:PluralitySuffix = (baseNumber === 1) ? '' : 's';
  // For some reason test case says "0 ohms", rather than "0 ohm" is correct. 
  // const pluralitySuffix:PluralitySuffix = (baseNumber === 1 || baseNumber === 0) ? '' : 's';
  const resistanceUnit:ResistanceUnit = `${metricUnitPrefix}ohm${pluralitySuffix}`;
  const resistanceValue:ResistanceValue = `${baseNumber} ${resistanceUnit}`;
  return resistanceValue;
}
