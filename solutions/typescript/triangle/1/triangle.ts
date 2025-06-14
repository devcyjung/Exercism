declare global {
  interface Array<T> {
    toSorted(): Array<T>
    toSorted(compareFn: (a: T, b: T) => number): Array<T>
  }
}

export class Triangle {
  private triangleType: 'equ' | 'iso' | 'sca' | 'not' = 'not'
  
  constructor(...sides: number[]) {
    if (sides.length !== 3) {
      return
    }
    const [shortest, medium, longest] = sides.toSorted((a, b) => a - b)
    if (longest >= shortest + medium) {
      return
    }
    if (longest === shortest) {
      this.triangleType = 'equ'
      return
    }
    if (medium === shortest || medium === longest) {
      this.triangleType = 'iso'
      return
    }
    this.triangleType = 'sca'
  }

  get isEquilateral(): boolean {
    return this.triangleType === 'equ'
  }

  get isIsosceles(): boolean {
    return this.triangleType === 'iso' || this.triangleType === 'equ'
  }

  get isScalene(): boolean {
    return this.triangleType === 'sca'
  }
}
