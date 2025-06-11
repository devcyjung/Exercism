export class List<T> {
  private items: T[] = []
  
  public static create<T>(...values: T[]): List<T> {
    const list = new List<T>()
    for (let i = 0; i < values.length; ++i) {
      list.set(i, values[i])
    }
    return list
  }

  toJSON(): string {
    let json = '['
    for (let i = 0; i < this.length(); ++i) {
      if (i > 0) {
        json += ', '
      }
      json += JSON.stringify(this.get(i))
    }
    json += ']'
    return json
  }

  [Symbol.iterator](): ListIterator<T> {
    return new ListIterator(this)
  }

  forEach(fn: (value: T, index: number) => void): void {
    for (let i = 0; i < this.length(); ++i) {
      fn(this.get(i), i)
    }
  }

  append(other: List<T>): List<T> {
    const result = new List<T>()
    let index = 0
    for (let i = 0; i < this.length(); ++i) {
      result.set(index++, this.get(i))
    }
    for (let i = 0; i < other.length(); ++i) {
      result.set(index++, other.get(i))
    }
    return result
  }

  concat(others: List<List<T>>): List<T> {
    const result = new List<T>()
    let index = 0
    for (let i = 0; i < this.length(); ++i) {
      result.set(index++, this.get(i))
    }
    for (let i = 0; i < others.length(); ++i) {
      const other = others.get(i)
      for (let j = 0; j < other.length(); ++j) {
        result.set(index++, other.get(j))
      }
    }
    return result
  }

  filter(fn: (value: T, index: number) => boolean): List<T> {
    const result = new List<T>()
    let index = 0
    for (let i = 0; i < this.length(); ++i) {
      if (fn(this.get(i), i)) {
        result.set(index++, this.get(i))
      }
    }
    return result
  }

  length(): number {
    return this.items.length
  }

  map<R>(fn: (value: T, index: number) => R): List<R> {
    const result = new List<R>()
    for (let i = 0; i < this.length(); ++i) {
      result.set(i, fn(this.get(i), i))
    }
    return result
  }

  foldl<R>(fn: (acc: R, value: T) => R, initial: R): R {
    let acc = initial
    for (let i = 0; i < this.length(); ++i) {
      acc = fn(acc, this.get(i))
    }
    return acc
  }

  foldr<R>(fn: (acc: R, value: T) => R, initial: R): R {
    let acc = initial
    for (let i = this.length() - 1; i >= 0; --i) {
      acc = fn(acc, this.get(i))
    }
    return acc
  }

  reverse(): List<T> {
    const result = new List<T>()
    let index = 0
    for (let i = this.length() - 1; i >= 0; --i) {
      result.set(index++, this.get(i))
    }
    return result
  }

  get(index: number): T {
    return this.items[index]
  }

  set(index: number, value: T): void {
    this.items[index] = value
  }
}

class ListIterator<T> {
  private list: List<T>
  private index: number

  constructor(list: List<T>) {
    this.list = list
    this.index = 0
  }
  
  next(): IteratorResult<T> {
    if (this.index < this.list.length()) {
      const value = this.list.get(this.index++)
      return { done: false, value }
    }
    return { done: true, value: undefined }
  }

  [Symbol.iterator](): ListIterator<T> {
    return this
  }
}