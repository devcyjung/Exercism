type Effect = () => void
type State<T> = () => T
type SetState<T> = (value: T) => void
type EffectBuffer = {
  register: boolean
  callback: Effect
  remove: Effect
}

const updateFns: Effect[] = []
const effectBuffer: EffectBuffer = {
  register: false,
  callback: () => {},
  remove: () => {}
}

export function createInput<T>(initial: T): [State<T>, SetState<T>] {
  let value = initial
  const index = updateFns.length
  const state = () => value
  const setState = (newValue: T) => {
    if (Object.is(value, newValue)) return
    value = newValue
    for (let i = index; i < updateFns.length; ++i) {
      updateFns[i]()
    }
  }
  return [state, setState]
}

export function createComputed<T>(compute: State<T>, initial?: T, equalOnly?: boolean): State<T> {
  let value = initial === undefined ? compute() : initial
  const effects = new Set<Effect>()
  updateFns.push(() => {
    const newValue = compute()
    if (!!equalOnly && Object.is(newValue, value)) return
    value = newValue
    for (const effect of effects) {
      effect()
    }
  })
  return () => {
    if (effectBuffer.register) {
      const callback = effectBuffer.callback
      effects.add(callback)
      const remove = effectBuffer.remove
      effectBuffer.remove = () => {
        effects.delete(callback)
        remove()
      }
    }
    return value
  }
}

export function createCallback(callback: Effect): Effect {
  effectBuffer.register = true
  effectBuffer.callback = callback
  effectBuffer.remove = () => {}
  callback()
  effectBuffer.register = false
  return effectBuffer.remove
}