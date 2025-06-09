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

const METRICS = [
  '',
  'kilo',
  'mega',
  'giga',
  'tera'
] as const;

export type Color = typeof COLORS[number];
export type Metric = typeof METRICS[number];

export type PluralitySuffix = ''|'s';
export type ResistanceUnit = `${Metric}ohm${PluralitySuffix}`;
export type ResistanceValue = `${number} ${ResistanceUnit}`;

export function decodedResistorValue(colors:Color[]):ResistanceValue {
  const num0:number = COLORS.indexOf(colors[0]!);
  const num1:number = COLORS.indexOf(colors[1]!);
  const num2:number = COLORS.indexOf(colors[2]!);
  
  const numericOhm:number = (num0 * 10 + num1) * 10 ** num2;
  let baseNum:number = numericOhm;
  let expThousand:number = 0;
  while (baseNum >= 1000) {
    baseNum /= 1000;
    ++expThousand;
  }
  
  const metricUnit:Metric = METRICS[expThousand];
  const pluralitySuffix:PluralitySuffix = (baseNum === 1) ? '' : 's';
  const resistanceUnit:ResistanceUnit = `${metricUnit}ohm${pluralitySuffix}`;
  const resistanceValue:ResistanceValue = `${baseNum} ${resistanceUnit}`;
  return resistanceValue;
}
