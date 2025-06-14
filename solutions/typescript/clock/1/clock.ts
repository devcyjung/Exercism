export class Clock {
  constructor(private hour: number, private minute: number = 0) {
    let netMinutes = (hour * 60 + minute) % (24 * 60)
    if (netMinutes < 0) {
      netMinutes += 24 * 60
    }
    this.hour = Math.floor(netMinutes / 60)
    this.minute = netMinutes % 60
  }

  public toString(): string {
    return `${String(this.hour).padStart(2, "0")}:${String(this.minute).padStart(2, "0")}`
  }

  public plus(minutes: number): Clock {
    return new Clock(this.hour, this.minute + minutes)
  }

  public minus(minutes: number): Clock {
    return this.plus(-minutes)
  }

  public equals(other: Clock): boolean {
    return this.hour - other.hour === 0 && this.minute - other.minute === 0
  }
}