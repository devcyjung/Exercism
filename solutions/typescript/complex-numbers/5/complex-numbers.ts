export class ComplexNumber {
  constructor(public readonly real: number = 0, public readonly imag: number = 0) {}

  public add(other: ComplexNumber): ComplexNumber {
    const real = this.real + other.real
    const imag = this.imag + other.imag
    return new ComplexNumber(real, imag)
  }

  public sub(other: ComplexNumber): ComplexNumber {
    const real = this.real - other.real
    const imag = this.imag - other.imag
    return new ComplexNumber(real, imag)
  }

  public div(other: ComplexNumber): ComplexNumber | never {
    const divisor = other.real * other.real + other.imag * other.imag
    if (divisor === 0) {
      throw new Error('Division by zero')
    }
    const dividend = this.mul(other.conj)
    const real = dividend.real / divisor
    const imag = dividend.imag / divisor
    return new ComplexNumber(real, imag)
  }

  public mul(other: ComplexNumber): ComplexNumber {
    const real = this.real * other.real - this.imag * other.imag
    const imag = this.real * other.imag + this.imag * other.real
    return new ComplexNumber(real, imag)
  }

  public pow(other: ComplexNumber): ComplexNumber {
    const { r, theta } = this.polar
    const logR = Math.log(r)
    const coefficient = Math.pow(Math.E, other.real * logR - other.imag * theta)
    const domain = other.real * theta + other.imag * logR
    const real = coefficient * Math.cos(domain)
    const imag = coefficient * Math.sin(domain)
    return new ComplexNumber(real, imag)
  }

  public get polar(): { r: number, theta: number } {
    return { r: this.abs, theta: Math.atan2(this.imag, this.real) }
  }

  public get abs(): number {
    return Math.hypot(this.real, this.imag)
  }

  public get conj(): ComplexNumber {
    return new ComplexNumber(this.real, this.imag === 0 ? this.imag : -this.imag)
  }

  public get exp(): ComplexNumber {
    const coefficient = Math.exp(this.real)
    const real = coefficient * Math.cos(this.imag)
    const imag = coefficient * Math.sin(this.imag)
    return new ComplexNumber(real, imag)
  }
}