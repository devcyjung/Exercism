const segmenter = new Intl.Segmenter()

export const truncate = input => segmenter.segment(input)[Symbol.iterator]().take(5).map(s => s.segment).toArray().join("")

// WRONG : produces "👍🏽👍🏽👍"
console.log(Array.from("👍🏽👍🏽👍🏽👍🏽👍🏽👍🏽").slice(0, 5).join("")) // This is WRONG: "👍🏽👍🏽👍"

// CORRECT: produces "👍🏽👍🏽👍🏽👍🏽👍🏽"
console.log(segmenter.segment("👍🏽👍🏽👍🏽👍🏽👍🏽👍🏽")[Symbol.iterator]().take(5).map(s => s.segment).toArray().join(""))