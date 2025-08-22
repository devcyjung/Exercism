const euclidRem = (dividend, divisor) => (dividend % divisor + divisor) % divisor
const divMod = (dividend, divisor) => [Math.floor(dividend / divisor), dividend % divisor]

export class Clock {
  constructor(hour = 0, minute = 0) {
    [this.hour, this.minute] = divMod(euclidRem(hour * 60 + minute, 24 * 60), 60)
  }

  toString() {
    return `${String(this.hour).padStart(2, '0')}:${String(this.minute).padStart(2, '0')}`
  }

  plus(minute) {
    return new Clock(this.hour, this.minute + minute)
  }

  minus(minute) {
    return this.plus(-minute)
  }

  equals(other) {
    return other instanceof Clock && this.hour === other.hour && this.minute === other.minute
  }
}
