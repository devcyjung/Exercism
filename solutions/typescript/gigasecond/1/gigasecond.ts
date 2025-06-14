export class Gigasecond {
  private static GIGASECONDS_IN_MS = 1_000_000_000_000
  
  public constructor(private now: Date) {}

  public date(): Date {
    return new Date(this.now.valueOf() + Gigasecond.GIGASECONDS_IN_MS)
  }
}