export class Gigasecond {
  private static GSEC_IN_MS = 1e12
  
  public constructor(private now: Date) {}

  public date(): Date {
    return new Date(this.now.valueOf() + Gigasecond.GSEC_IN_MS)
  }
}