export class DiffieHellman {
  constructor(private p: number, private g: number) {
    if (!DiffieHellman.isPrime(p) || !DiffieHellman.isPrime(g)) {
      throw new Error('Must provide 2 prime numbers')
    }
  }

  private static primeTable = new Set([2, 3])
  private static maxCalculated = 3
  
  private static isPrime(input: number): boolean {
    if (DiffieHellman.primeTable.has(input)) {
      return true
    }
    if (((input >> 1) & 1) === 0) {
      return false
    }
    if (DiffieHellman.maxCalculated < input) {
      DiffieHellman.extendPrimeTable(input)
    }
    return DiffieHellman.primeTable.has(input)
  }

  private static extendPrimeTable(limit: number): void {
    for (let i = DiffieHellman.maxCalculated + 2; i <= limit; i += 2) {
      let isPrime = true
      for (const p of DiffieHellman.primeTable) {
        if (i % p === 0) {
          isPrime = false
          break
        }
      }
      if (isPrime) {
        DiffieHellman.primeTable.add(i)
      }
    }
    DiffieHellman.maxCalculated = limit
  }

  private validPrivate(privateKey: number): boolean {
    return 1 < privateKey && privateKey < this.p
  }

  public getPublicKey(privateKey: number): number {
    if (!this.validPrivate(privateKey)) {
      throw new Error('Invalid private key')
    }
    let result = 1
    for (let i = 0; i < privateKey; ++i) {
      result *= this.g
      result %= this.p
    }
    return result
  }

  public getSecret(theirPublicKey: number, myPrivateKey: number): number {
    if (!this.validPrivate(myPrivateKey)) {
      throw new Error('Invalid private key')
    }
    let result = 1
    for (let i = 0; i < myPrivateKey; ++i) {
      result *= theirPublicKey
      result %= this.p
    }
    return result
  }
}