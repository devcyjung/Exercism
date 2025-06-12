export class Robot {
  private static initialNames = ((): string[] => {
    const names = Array.from({ length: 26 * 26 * 10 * 10 * 10 }, (_, n) => {
      let name = ""
      let tmp = n
      for (let i = 0; i < 3; ++i) {
        name += String(tmp % 10)
        tmp = Math.floor(tmp / 10)
      }
      for (let i = 0; i < 2; ++i) {
        name = String.fromCharCode(65 + (tmp % 26)) + name
        tmp = Math.floor(tmp / 26)
      }
      return name
    })
    for (let i = names.length - 1; i > 0; --i) {
      const j = Math.floor(Math.random() * (i + 1))
      const [a, b] = [names[i], names[j]]
      names[i] = b
      names[j] = a
    }
    return names
  })()
  
  private static names = [...Robot.initialNames]

  #name: string
  
  constructor() {
    if (Robot.names.length === 0) {
      throw new Error('No available name left')
    }
    this.#name = Robot.names.pop()!
  }

  public get name(): string {
    return this.#name
  }

  public resetName(): void {
    Robot.names.unshift(this.#name)
    this.#name = Robot.names.pop()!
  }
  
  public static releaseNames(): void {
    Robot.names = [...Robot.initialNames]
  }
}