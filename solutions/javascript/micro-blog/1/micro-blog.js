const segmenter = new Intl.Segmenter()

export const truncate = (input) => {
  return segmenter.segment(input)[Symbol.iterator]().take(5).map(s => s.segment).toArray().join("")
};