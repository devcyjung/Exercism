function createInput(initial: number): [() => number, (newValue: number) => void] {
  const ID = CallbackDAG.length
  const inputCell = {
    value: initial,
  }
  return [
    () => inputCell.value,
    (newValue) => {
      if (newValue === inputCell.value) {
        return
      }
      ++entropy
      inputCell.value = newValue
      CallbackDAG.slice(ID).forEach(callback => callback())
    },
  ]
}

function createComputed(
  compute: () => number, initial?: number, equalOnly?: boolean
): () => number {
  const ID = CallbackDAG.length
  const computeCell = {
    value: (initial === undefined) ? compute() : initial,
    lastEntropy: entropy,
  }
  callbackMap.set(ID, [])
  const getter = () => {
    if (callbackToRegister !== undefined) {
      cancelerID[0] = ID
      cancelerID[1] = callbackMap.get(ID)?.length
      callbackMap.get(ID)?.push(callbackToRegister)
      callbackToRegister = undefined
    }
    if (entropy === computeCell.lastEntropy) {
      return computeCell.value
    }
    computeCell.lastEntropy = entropy
    const newValue = compute()
    if (equalOnly === true && Object.is(newValue, computeCell.value)) {
      return computeCell.value
    }
    computeCell.value = newValue
    callbackMap.get(ID)?.forEach((callback, callbackIdx) => {
      if (callback !== undefined) {
        callback()
      }
    })
    return computeCell.value
  }
  CallbackDAG.push(getter)
  return getter
}

function createCallback(callback: () => void): () => void {
  callbackToRegister = callback
  callback()
  if (cancelerID[0] === undefined || cancelerID[1] === undefined) {
    throw new Error("Registrar has not set the cancelerID")
  }
  const ID = [cancelerID[0], cancelerID[1]]
  return () => {
    const entry = callbackMap.get(ID[0])
    if (entry !== undefined) {
      entry[ID[1]] = undefined
    }
  }
}

const CallbackDAG: (() => void)[] = []
let entropy: number = 0
let callbackToRegister: (() => void) | undefined
const cancelerID: [number | undefined, number | undefined] = [undefined, undefined]
const callbackMap = new Map<number, ((() => void) | undefined)[]>()

export { createInput, createComputed, createCallback }