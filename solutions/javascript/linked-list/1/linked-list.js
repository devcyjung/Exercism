export class LinkedList {
  #head
  #len

  constructor() {
    this.#len = 0
  }
  
  [Symbol.iterator]() {
    let cursor = this.#head
    return {
      next() {
        const iter = {
          value: cursor,
          done: cursor === undefined
        }
        cursor = cursor?.next
        return iter
      }
    }
  }

  get tail() {
    let last
    for (const cursor of this) {
      last = cursor
    }
    return last
  }
  
  push(newValue) {
    ++this.#len
    const newNode = {
      value: newValue
    }
    if (this.#head === undefined) {
      this.#head = newNode
      return
    }
    const tail = this.tail
    tail.next = newNode
  }

  pop() {
    if (this.#len === 0) return
    --this.#len
    if (this.#len === 0) {
      const value = this.#head.value
      this.#head = undefined
      return value
    }
    let counter = 0
    for (const cursor of this) {
      if (++counter === this.#len) {
        const value = cursor.next.value
        cursor.next = undefined
        return value
      }
    }
  }

  shift() {
    if (this.#len === 0) return
    --this.#len
    const value = this.#head.value
    if (this.#len === 0) {
      this.#head = undefined
    } else {
      this.#head = this.#head.next
    }
    return value
  }

  unshift(newValue) {
    ++this.#len
    const newNode = {
      value: newValue,
      next: this.#head
    }
    this.#head = newNode
  }

  delete(value) {
    let prev
    for (const cursor of this) {
      if (Object.is(cursor.value, value)) {
        --this.#len
        if (prev === undefined) {
          this.#head = cursor.next
          cursor.next = undefined
        } else {
          prev.next = cursor.next
          cursor.next = undefined
        }
        return
      }
      prev = cursor
    }
  }

  count() {
    return this.#len
  }
}
