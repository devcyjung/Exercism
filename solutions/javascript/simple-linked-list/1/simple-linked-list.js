export class Element {
  #value
  #next

  constructor(value) {
    this.#value = value
    this.#next = null
  }

  get value() {
    return this.#value
  }

  get next() {
    return this.#next
  }

  set next(nextElement) {
    this.#next = nextElement
  }
}

export class List {
  #len
  #head

  constructor(array = []) {
    this.#len = 0
    this.#head = null
    for (const value of array) {
      this.add(new Element(value))
    }
  }

  add(nextElement) {
    ++this.#len
    nextElement.next = this.#head
    this.#head = nextElement
  }

  get length() {
    return this.#len
  }

  get head() {
    return this.#head
  }

  toArray() {
    return Array.from(this)
  }

  reverse() {
    let cursor = this.#head
    let prev = null
    while (cursor !== null) {
      const next = cursor.next
      cursor.next = prev
      prev = cursor
      this.#head = cursor
      cursor = next
    }
    return this
  }

  [Symbol.iterator]() {
    let cursor = this.#head
    return {
      next() {
        const iter = {
          value: cursor?.value,
          done: cursor === null,
        }
        cursor = cursor?.next
        return iter
      }
    }
  }
}
