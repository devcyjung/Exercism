export class ComplexNumber {
  #real
  #imag
  
  constructor(real = 0, imag = 0) {
    this.#real = real
    this.#imag = imag
  }

  get real() {
    return this.#real
  }

  get imag() {
    return this.#imag
  }

  add(other) {
    const real = this.real + other.real
    const imag = this.imag + other.imag
    return new ComplexNumber(real, imag)
  }

  sub(other) {
    const real = this.real - other.real
    const imag = this.imag - other.imag
    return new ComplexNumber(real, imag)
  }

  div(other) {
    const divisor = other.real * other.real + other.imag * other.imag
    if (divisor === 0) {
      throw new Error('Division by zero')
    }
    const dividend = this.mul(other.conj)
    const real = dividend.real / divisor
    const imag = dividend.imag / divisor
    return new ComplexNumber(real, imag)
  }

  mul(other) {
    const real = this.real * other.real - this.imag * other.imag
    const imag = this.real * other.imag + this.imag * other.real
    return new ComplexNumber(real, imag)
  }

  pow(other) {
    const { r, theta } = this.polar
    const logR = Math.log(r)
    const coefficient = Math.pow(Math.E, other.real * logR - other.imag * theta)
    const domain = other.real * theta + other.imag * logR
    const real = coefficient * Math.cos(domain)
    const imag = coefficient * Math.sin(domain)
    return new ComplexNumber(real, imag)
  }

  get polar() {
    return { r: this.abs, theta: Math.atan2(this.imag, this.real) }
  }

  get abs() {
    return Math.hypot(this.real, this.imag)
  }

  get conj() {
    return new ComplexNumber(this.real, this.imag === 0 ? this.imag : -this.imag)
  }

  get exp() {
    const coefficient = Math.exp(this.real)
    const real = coefficient * Math.cos(this.imag)
    const imag = coefficient * Math.sin(this.imag)
    return new ComplexNumber(real, imag)
  }
}
