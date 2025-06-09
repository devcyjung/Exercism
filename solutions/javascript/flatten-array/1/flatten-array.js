/**
 * @internal
 * Recursively flattens a nested iterator, yielding all non-null, non-undefined elements.
 *
 * Strings are treated as atomic and not flattened. Iterable objects are recursively expanded.
 *
 * @param {Iterator<any>} iter - The input iterator to flatten.
 * @returns {Iterator<any>} - A new iterator yielding flattened values.
 */
const flatIterator = iter => iter.flatMap(elem => {
  if (elem === undefined || elem === null) {
    return Iterator.from([])
  }
  if (typeof elem !== 'object') {
    return Iterator.from([elem])
  }
  if (Symbol.iterator in elem) {
    return flatIterator(Iterator.from(elem))
  }
  return Iterator.from([elem])
})

/**
 * Recursively flattens any nested iterable or iterator into a flat array.
 *
 * Strings are treated as atomic values and not flattened.
 *
 * @param {*} input - The input to flatten. Can be anything.
 * @returns {Array<unknown>} - A fully flattened array containing all non-null, non-undefined values.
 */
export const flatten = input => {
  if (input === undefined || input === null) {
    return []
  }
  if (typeof input !== 'object') {
    return [input]
  }
  if (Symbol.iterator in input) {
    return flatIterator(Iterator.from(input)).toArray()
  }
  return [input]
}