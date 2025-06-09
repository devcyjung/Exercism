type Callback = () => void
type State<T> = () => T
type SetState<T> = (value: T) => void
type InputPair<T> = [State<T>, SetState<T>]
type CallbackContext = {
  callbacks: Set<Callback>,
  subscribers: Set<Set<Callback>>
}
type ComputeContext = {
  computes: Set<Callback>,
}

let activeCallback: CallbackContext | undefined
let activeCompute: ComputeContext | undefined

function createInput<T>(
  initial: T,
): InputPair<T> {
  let value: T = initial
  const subscriberComputes: Set<Callback> = new Set()
  const state: State<T> = () => {
    if (activeCompute) {
      activeCompute.computes.forEach((compute) => {
        subscriberComputes.add(compute)
      })
    }
    return value
  }
  const setState: SetState<T> = (newValue: T) => {
    value = newValue
    subscriberComputes.forEach((compute) => {
      compute()
    })
  }
  return [state, setState]
}

function createComputed<T>(
  updateFn: State<T>,
  _initial?: T,
  _ignoreEqual?: boolean,
): State<T> {
  let value: T | undefined
  const subscriberCallbacks: Set<Callback> = new Set()
  return () => {
    if (!activeCallback) {
      value = updateFn()
      return value
    } else {
      activeCallback.callbacks.forEach((callback) => {
        subscriberCallbacks.add(callback)
      })
      activeCallback.subscribers.add(subscriberCallbacks)
      activeCallback = undefined
      activeCompute = {
        computes: new Set([(): T => {
          const prev = value
          value = updateFn()
          if (!Object.is(prev, value)) {
            subscriberCallbacks.forEach((callback) => {
              callback()
            })
          }
          return value
        }])
      }
      value = updateFn()
      return value
    }
  }
}

function createCallback(callback: Callback): Callback {
  const callbackRegistry: Set<Set<Callback>> = new Set()
  activeCallback = {
    callbacks: new Set([callback]),
    subscribers: callbackRegistry,
  }
  callback()
  return () => {
    callbackRegistry.forEach((registry) => {
      registry.delete(callback)
    })
  }
}

export { createInput, createComputed, createCallback }