declare global {
  interface Set<T> {
    isSubsetOf(other: Set<T>): boolean
    isDisjointFrom(other: Set<T>): boolean
    symmetricDifference(other: Set<T>): Set<T>
    union(other: Set<T>): Set<T>
    intersection(other: Set<T>): Set<T>
    difference(other: Set<T>): Set<T>
  }
}

export class CustomSet<T> {
  private set: Set<T>
  
  constructor(initial?: T[] | Set<T>) {
    if (initial instanceof Set) {
      this.set = initial
      return
    }
    this.set = new Set(initial)
  }

  empty(): boolean {
    return this.set.size === 0
  }

  contains(element: T): boolean {
    return this.set.has(element)
  }

  add(element: T): CustomSet<T> {
    return new CustomSet(this.set.add(element))
  }

  subset(other: CustomSet<T>): boolean {
    return this.set.isSubsetOf(other.set)
  }

  disjoint(other: CustomSet<T>): boolean {
    return this.set.isDisjointFrom(other.set)
  }

  eql(other: CustomSet<T>): boolean {
    return this.set.symmetricDifference(other.set).size === 0
  }

  union(other: CustomSet<T>): CustomSet<T> {
    return new CustomSet(this.set.union(other.set))
  }

  intersection(other: CustomSet<T>): CustomSet<T> {
    return new CustomSet(this.set.intersection(other.set))
  }

  difference(other: CustomSet<T>): CustomSet<T> {
    return new CustomSet(this.set.difference(other.set))
  }
}