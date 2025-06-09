const segmenter = new Intl.Segmenter()

export const truncate = (input) => {
  console.log(Array.from("👍🏽👍🏽👍🏽👍🏽👍🏽👍🏽").slice(0, 5).join("")) // This is WRONG: "👍🏽👍🏽👍"
  console.log(segmenter.segment("👍🏽👍🏽👍🏽👍🏽👍🏽👍🏽")[Symbol.iterator]().take(5).map(s => s.segment).toArray().join("")) // This is CORRECT: "👍🏽👍🏽👍🏽👍🏽👍🏽"
  return segmenter.segment(input)[Symbol.iterator]().take(5).map(s => s.segment).toArray().join("")
};