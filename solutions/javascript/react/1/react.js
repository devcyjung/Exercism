export class InputCell {
  #callbacks = new Map()
  #value
  
  constructor(value) {
    this.#value = value
  }

  setValue(value) {
    this.value = value
  }

  get value() {
    return this.compute()
  }

  set value(newValue) {
    if (Object.is(this.#value, newValue)) {
      return
    }
    this.#value = newValue
    this.#callbacks.forEach(fn => {
      fn()
    })
  }

  compute() {
    return this.#value
  }

  addCallback(cb, fn = () => cb.values.push(cb.fn(this))) {
    this.#callbacks.set(cb, fn)
  }

  removeCallback(cb) {
    this.#callbacks.delete(cb)
  }
}

export class ComputeCell extends InputCell {
  #compute
  
  constructor(inputCells, fn) {
    super(fn(inputCells))
    this.#compute = () => fn(inputCells)
    inputCells.forEach(cell => {
      cell.addCallback(this, () => this.value = fn(inputCells))
    })
  }

  compute() {
    return this.#compute()
  }
}

export class CallbackCell {
  constructor(fn) {
    this.values = []
    this.fn = fn
  }
}