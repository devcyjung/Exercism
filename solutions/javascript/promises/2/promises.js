/**
 * Converts a callback-based function into a Promise-based one.
 * @template T
 * @param {(arg: any, cb: (err: Error | null, result?: T) => void) => void} f - The callback-style function to promisify.
 * @returns {(arg: any) => Promise<T>} A function that returns a Promise.
 */
function promisify(f) {
  return arg =>
    new Promise((resolve, reject) => {
      f(arg, (err, result) => {
        if (err) {
          reject(err)
        } else {
          resolve(result)
        }
      })
    })
}

/**
 * Resolves when all promises resolve, or rejects on the first rejection.
 * @template T
 * @param {Array<Promise<T>>} [promises=[]] - Array of promises.
 * @returns {Promise<T[]>} A promise that resolves to an array of resolved values.
 */
function all(promises) {
  return new Promise((resolve, reject) => {
    if (!promises) {
      resolve()
    }
    if (promises.length === 0) {
      resolve([])
    }
    const results = Array.from({ length: promises.length })
    let resolveCount = 0
    promises.forEach((p, i) => {
      Promise.resolve(p)
        .then(result => {
          results[i] = result
          if (++resolveCount === promises.length) {
            resolve(results)
          }
        })
        .catch(error => reject(error))
    })
  })
}

/**
 * Resolves when all promises settle, returning an array of results or errors.
 * @template T
 * @param {Array<Promise<T>>} [promises=[]] - Array of promises.
 * @returns {Promise<Array<T | Error>>} A promise that resolves to all results or errors.
 */
function allSettled(promises) {
  return new Promise((resolve, reject) => {
    if (!promises) {
      resolve()
    }
    if (promises.length === 0) {
      resolve([])
    }
    const results = Array.from({ length: promises.length })
    let resolveCount = 0
    promises.forEach((p, i) => {
      Promise.resolve(p)
        .then(result => {
          results[i] = result
          if (++resolveCount === promises.length) {
            resolve(results)
          }
        })
        .catch(error => {
          results[i] = error
          if (++resolveCount === promises.length) {
            resolve(results)
          }
        })
    })
  })
}

/**
 * Resolves or rejects as soon as any promise resolves or rejects.
 * @template T
 * @param {Array<Promise<T>>} [promises=[]] - Array of promises.
 * @returns {Promise<T>} A promise that resolves/rejects with the first settled value.
 */
function race(promises) {
  return new Promise((resolve, reject) => {
    if (!promises) {
      resolve()
    }
    if (promises.length === 0) {
      resolve([])
    }
    promises.forEach(p => {
      Promise.resolve(p)
        .then(result => resolve(result))
        .catch(error => reject(error))
    })
  })
}

/**
 * Resolves as soon as any promise resolves, or rejects with all errors if all fail.
 * @template T
 * @param {Array<Promise<T>>} [promises=[]] - Array of promises.
 * @returns {Promise<T>} A promise that resolves to the first successful value, or rejects with an array of errors.
 */
function any(promises) {
  return new Promise((resolve, reject) => {
    if (!promises) {
      resolve()
    }
    if (promises.length === 0) {
      resolve([])
    }
    const rejects = Array.from({ length: promises.length })
    let rejectCount = 0
    promises.forEach((p, i) => {
      Promise.resolve(p)
        .then(result => resolve(result))
        .catch(error => {
          rejects[i] = error
          if (++rejectCount === promises.length) {
            reject(rejects)
          }
        })
    })
  })
}

export {
  promisify,
  all,
  allSettled,
  race,
  any,
}