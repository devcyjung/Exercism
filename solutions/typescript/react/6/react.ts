type Effect = () => void
type Getter<T> = () => T
type Setter<T> = (value: T) => void
type EffectBuffer = {
  register: boolean
  effect: Effect
}

const updateFns: (() => void)[] = []
const effectBuffer: EffectBuffer = {
  register: false,
  effect: () => {}
}

export function createInput<T>(initial: T): [Getter<T>, Setter<T>] {
  let value = initial
  const index = updateFns.length
  const getter = () => value
  const setter = (newValue: T) => {
    if (Object.is(value, newValue)) return
    value = newValue
    for (let i = index; i < updateFns.length; ++i) {
      updateFns[i]()
    }
  }
  return [getter, setter]
}

export function createComputed<T>(compute: Getter<T>, initial?: T, equalOnly?: boolean): Getter<T> {
  let value = initial === undefined ? compute() : initial
  const callbacks = new Set<Effect>()
  updateFns.push(() => {
    const newValue = compute()
    if (!!equalOnly && Object.is(newValue, value)) return
    value = newValue
    callbacks.forEach(cb => {
      cb()
    })
  })
  return () => {
    if (effectBuffer.register) {
      effectBuffer.register = false
      const callback = effectBuffer.effect
      callbacks.add(callback)
      effectBuffer.effect = () => {
        callbacks.delete(callback)
      }
    }
    return value
  }
}

export function createCallback(callback: Effect): Effect {
  effectBuffer.register = true
  effectBuffer.effect = callback
  callback()
  return effectBuffer.effect
}